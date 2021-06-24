//
// (C) Copyright 2020-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package main

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	mgmtpb "github.com/daos-stack/daos/src/control/common/proto/mgmt"
	"github.com/daos-stack/daos/src/control/lib/atm"
	"github.com/daos-stack/daos/src/control/lib/netdetect"
	"github.com/daos-stack/daos/src/control/logging"
)

const (
	invalidIndex         = -1
	verbsProvider        = "ofi+verbs"
	defaultNetworkDevice = "lo"
	defaultDomain        = "lo"
)

func newAttachInfoCache(log logging.Logger, enabled bool) *attachInfoCache {
	return &attachInfoCache{
		log:     log,
		enabled: atm.NewBool(enabled),
	}
}

type attachInfoCache struct {
	mutex sync.Mutex

	log         logging.Logger
	enabled     atm.Bool
	initialized atm.Bool

	// cached response from remote server
	AttachInfo *mgmtpb.GetAttachInfoResp
}

// IsEnabled reports whether the cache is enabled.
func (c *attachInfoCache) IsEnabled() bool {
	if c == nil {
		return false
	}

	return c.enabled.IsTrue()
}

// IsCached reports whether there is data in the cache.
func (c *attachInfoCache) IsCached() bool {
	if c == nil {
		return false
	}
	return c.enabled.IsTrue() && c.initialized.IsTrue()
}

// Cache preserves the results of a GetAttachInfo remote call.
func (c *attachInfoCache) Cache(ctx context.Context, resp *mgmtpb.GetAttachInfoResp) error {
	if c == nil {
		return errors.New("nil attachInfoCache")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.IsEnabled() {
		return errors.New("cache is not enabled")
	}

	if resp == nil {
		return errors.New("nil input")
	}

	c.AttachInfo = resp

	c.initialized.SetTrue()
	return nil
}

func newLocalFabricCache(log logging.Logger) *localFabricCache {
	return &localFabricCache{
		log:               log,
		currentNumaDevIdx: make(map[int]int),
	}
}

type fabricInterface struct {
	Name        string
	Alias       string
	NetDevClass uint32
}

type localFabricCache struct {
	sync.Mutex // Caller should lock and unlock around cache operations

	log         logging.Logger
	initialized atm.Bool
	// cached fabric interfaces organized by NUMA affinity
	localNUMAFabric map[int][]*fabricInterface
	// maps NUMA affinity to a device index
	currentNumaDevIdx map[int]int
	defaultNumaNode   int // if no NUMA node specified by client

	getDevAlias func(ctx context.Context, devName string) (string, error)
}

// IsCached reports whether there is data in the cache.
func (c *localFabricCache) IsCached() bool {
	if c == nil {
		return false
	}

	return c.initialized.IsTrue()
}

// Cache caches the results of a fabric scan locally.
func (c *localFabricCache) Cache(ctx context.Context, scan []*netdetect.FabricScan) error {
	if c == nil {
		return errors.New("nil localFabricCache")
	}

	if c.getDevAlias == nil {
		c.getDevAlias = netdetect.GetDeviceAlias
	}

	c.localNUMAFabric = make(map[int][]*fabricInterface)

	for _, fs := range scan {
		if fs.DeviceName == "lo" {
			continue
		}

		haveDefaultNuma := false

		newIF := &fabricInterface{
			Name:        fs.DeviceName,
			NetDevClass: fs.NetDevClass,
		}

		deviceAlias, err := c.getDevAlias(ctx, newIF.Name)
		if err != nil {
			c.log.Debugf("Non-fatal error: failed to get device alias for %q: %v", newIF.Name, err)
		} else {
			newIF.Alias = deviceAlias
		}

		numa := int(fs.NUMANode)
		c.localNUMAFabric[numa] = append(c.localNUMAFabric[numa], newIF)

		// Any client bound to a NUMA node that has no network devices associated with it will
		// get a response from this defaultNumaNode.
		// This response offers a valid network device at degraded performance for those clients
		// that need it.
		if !haveDefaultNuma {
			c.defaultNumaNode = numa
			haveDefaultNuma = true
			c.log.Debugf("The default NUMA node is: %d", c.defaultNumaNode)
		}

		c.log.Debugf("Added device %q, alias %q for NUMA %d, device number %d",
			newIF.Name, newIF.Alias, numa, len(c.localNUMAFabric[numa])-1)
	}

	if _, exists := c.localNUMAFabric[c.defaultNumaNode]; !exists {
		c.log.Info("No network devices detected in fabric scan; default AttachInfo response may be incorrect\n")

		defaultIF := &fabricInterface{
			Name: defaultNetworkDevice,
		}
		c.localNUMAFabric[c.defaultNumaNode] = append(c.localNUMAFabric[c.defaultNumaNode], defaultIF)
	}

	c.initialized.SetTrue()
	return nil
}

// loadBalance is a simple round-robin load balancing scheme
// to assign network interface adapters to clients
// on the same NUMA node that have multiple adapters
// to choose from.  Returns the index of the device to use.
func (c *localFabricCache) loadBalance(numaNode int) int {
	deviceIndex := invalidIndex
	numDevs := len(c.localNUMAFabric[numaNode])
	if numDevs > 0 {
		deviceIndex = c.currentNumaDevIdx[numaNode]
		c.currentNumaDevIdx[numaNode] = (deviceIndex + 1) % numDevs
	}
	return deviceIndex
}

func (c *localFabricCache) GetDevice(numaNode int, netDevClass uint32) (*fabricInterface, error) {
	if c == nil {
		return nil, errors.New("nil localFabricCache")
	}

	if !c.IsCached() {
		return nil, errors.New("fabric data not cached")
	}

	for checked := 0; checked < len(c.localNUMAFabric[numaNode]); checked++ {
		fabricIF, err := c.getNextDevice(numaNode)
		if err != nil {
			return nil, err
		}

		if fabricIF.NetDevClass != netDevClass {
			c.log.Debugf("Excluding device: %s, network device class: %s from attachInfoCache. Does not match network device class: %s",
				fabricIF.Name, netdetect.DevClassName(fabricIF.NetDevClass), netdetect.DevClassName(netDevClass))
			continue
		}

		return fabricIF, nil
	}

	return nil, errors.New("no suitable fabric interface found")
}

// selectRemote is called when no local network adapters
// was detected on the local NUMA node and we need to pick
// one attached to a remote NUMA node. Some balancing is
// still required here to avoid overloading a specific interface
func (c *localFabricCache) selectRemote() (int, int) {
	deviceIndex := invalidIndex
	// restart from previous NUMA node
	numaNode := c.defaultNumaNode

	for i := 1; i <= len(c.localNUMAFabric); i++ {
		node := (numaNode + i) % len(c.localNUMAFabric)
		numDevs := len(c.localNUMAFabric[node])
		if numDevs > 0 {
			numaNode = node
			deviceIndex = c.currentNumaDevIdx[numaNode]
			c.currentNumaDevIdx[numaNode] = (deviceIndex + 1) % numDevs
			break
		}
	}

	// update the default with the one we finally picked
	c.defaultNumaNode = numaNode

	return numaNode, deviceIndex
}

func (c *localFabricCache) getNextDevice(numaNode int) (*fabricInterface, error) {
	deviceIndex := c.loadBalance(numaNode)

	// no fabric available for the client's actual NUMA node
	if deviceIndex == invalidIndex {
		var numaSel int

		numaSel, deviceIndex = c.selectRemote()
		if deviceIndex == invalidIndex {
			return nil, errors.New("No fallback network device found")
		}
		c.log.Infof("No network devices bound to client NUMA node %d.  Using response from NUMA %d", numaNode, numaSel)
		numaNode = numaSel
	}

	numaFabric := c.localNUMAFabric[numaNode]
	if len(numaFabric) <= deviceIndex {
		return nil, errors.Errorf("Device for numaNode %d device index %d did not exist", numaNode, deviceIndex)
	}
	fabricIF := numaFabric[deviceIndex]

	return fabricIF, nil
}
