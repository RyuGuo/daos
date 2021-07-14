//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package main

import (
	"context"

	"github.com/daos-stack/daos/src/control/lib/netdetect"
	"github.com/daos-stack/daos/src/control/logging"
	"github.com/pkg/errors"
)

// FabricInterface represents a generic fabric interface.
type FabricInterface struct {
	Name        string `yaml:"name"`
	Domain      string `yaml:"domain"`
	NetDevClass uint32 `yaml:"net_dev_class"`
}

// NUMAFabric represents a set of fabric interfaces organized by NUMA node.
type NUMAFabric struct {
	numaMap map[int][]*FabricInterface
}

// Add adds a fabric interface to a specific NUMA node.
func (n *NUMAFabric) Add(numaNode int, fi *FabricInterface) error {
	if n == nil {
		return errors.New("nil NUMAFabric")
	}
	n.numaMap[numaNode] = append(n.numaMap[numaNode], fi)
	return nil
}

// NumDevices gets the number of devices on a given NUMA node.
func (n *NUMAFabric) NumDevices(numaNode int) int {
	if n == nil {
		return 0
	}
	if devs, exist := n.numaMap[numaNode]; exist {
		return len(devs)
	}
	return 0
}

// NumNUMANodes gets the number of NUMA nodes.
func (n *NUMAFabric) NumNUMANodes() int {
	if n == nil {
		return 0
	}
	return len(n.numaMap)
}

func newNUMAFabric() *NUMAFabric {
	return &NUMAFabric{
		numaMap: make(map[int][]*FabricInterface),
	}
}

// NUMAFabricFromScan generates a NUMAFabric from a fabric scan result.
func NUMAFabricFromScan(ctx context.Context, log logging.Logger, scan []*netdetect.FabricScan,
	getDevAlias func(ctx context.Context, devName string) (string, error)) *NUMAFabric {
	fabric := newNUMAFabric()

	for _, fs := range scan {
		if fs.DeviceName == "lo" {
			continue
		}

		newIF := &FabricInterface{
			Name:        fs.DeviceName,
			NetDevClass: fs.NetDevClass,
		}

		if getDevAlias != nil {
			deviceAlias, err := getDevAlias(ctx, newIF.Name)
			if err != nil {
				log.Debugf("Non-fatal error: failed to get device alias for %q: %v", newIF.Name, err)
			} else {
				newIF.Domain = deviceAlias
			}
		}

		numa := int(fs.NUMANode)
		fabric.Add(numa, newIF)

		log.Debugf("Added device %q, domain %q for NUMA %d, device number %d",
			newIF.Name, newIF.Domain, numa, fabric.NumDevices(numa)-1)
	}

	return fabric
}
