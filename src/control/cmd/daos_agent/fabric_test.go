//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package main

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"

	"github.com/daos-stack/daos/src/control/common"
	"github.com/daos-stack/daos/src/control/lib/netdetect"
	"github.com/daos-stack/daos/src/control/logging"
)

func TestAgent_NewNUMAFabric(t *testing.T) {
	result := newNUMAFabric()

	if result == nil {
		t.Fatal("result was nil")
	}

	if result.numaMap == nil {
		t.Fatal("map was nil")
	}
}

func TestAgent_NUMAFabric_NumNUMANodes(t *testing.T) {
	for name, tc := range map[string]struct {
		nf        *NUMAFabric
		expResult int
	}{
		"nil": {},
		"empty struct": {
			nf: &NUMAFabric{},
		},
		"multiple nodes": {
			nf: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					0:  {},
					1:  {},
					3:  {},
					10: {},
				},
			},
			expResult: 4,
		},
	} {
		t.Run(name, func(t *testing.T) {
			common.AssertEqual(t, tc.expResult, tc.nf.NumNUMANodes(), "")
		})
	}
}

func TestAgent_NUMAFabric_NumDevices(t *testing.T) {
	for name, tc := range map[string]struct {
		nf        *NUMAFabric
		node      int
		expResult int
	}{
		"nil": {},
		"empty": {
			nf:   &NUMAFabric{},
			node: 5,
		},
		"multiple devices on node": {
			nf: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					3: {&FabricInterface{}, &FabricInterface{}},
				},
			},
			node:      3,
			expResult: 2,
		},
		"multiple nodes": {
			nf: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					0:  {&FabricInterface{}, &FabricInterface{}, &FabricInterface{}},
					1:  {&FabricInterface{}},
					3:  {&FabricInterface{}},
					10: {&FabricInterface{}},
				},
			},
			node:      0,
			expResult: 3,
		},
	} {
		t.Run(name, func(t *testing.T) {
			common.AssertEqual(t, tc.expResult, tc.nf.NumDevices(tc.node), "")
		})
	}
}

func TestAgent_NUMAFabric_Add(t *testing.T) {
	for name, tc := range map[string]struct {
		nf        *NUMAFabric
		input     *FabricInterface
		node      int
		expErr    error
		expResult *NUMAFabric
	}{
		"nil": {
			expErr: errors.New("nil NUMAFabric"),
		},
		"empty": {
			nf:    newNUMAFabric(),
			input: &FabricInterface{Name: "test1"},
			node:  2,
			expResult: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					2: {{Name: "test1"}},
				},
			},
		},
		"non-empty": {
			nf: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					2: {{Name: "t1"}},
					3: {{Name: "t2"}, {Name: "t3"}},
				},
			},
			input: &FabricInterface{Name: "test1"},
			node:  2,
			expResult: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					2: {{Name: "t1"}, {Name: "test1"}},
					3: {{Name: "t2"}, {Name: "t3"}},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			err := tc.nf.Add(tc.node, tc.input)

			common.CmpErr(t, tc.expErr, err)
			if diff := cmp.Diff(tc.expResult, tc.nf, cmp.AllowUnexported(NUMAFabric{})); diff != "" {
				t.Fatalf("-want, +got:\n%s", diff)
			}
		})
	}
}

func TestAgent_NUMAFabricFromScan(t *testing.T) {
	for name, tc := range map[string]struct {
		input       []*netdetect.FabricScan
		getDevAlias func(ctx context.Context, devName string) (string, error)
		expResult   *NUMAFabric
	}{
		"no devices in scan": {
			expResult: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{},
			},
		},
		"skip lo": {
			input: []*netdetect.FabricScan{
				{
					Provider:    "ofi+sockets",
					DeviceName:  "test0",
					NUMANode:    1,
					NetDevClass: netdetect.Ether,
				},
				{
					Provider:   "ofi+sockets",
					DeviceName: "lo",
					NUMANode:   1,
				},
			},
			expResult: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					1: {
						{
							Name:        "test0",
							NetDevClass: netdetect.Ether,
						},
					},
				},
			},
		},
		"multiple devices": {
			input: []*netdetect.FabricScan{
				{
					Provider:    "ofi+sockets",
					DeviceName:  "test0",
					NUMANode:    1,
					NetDevClass: netdetect.Ether,
				},
				{
					Provider:    "ofi+verbs",
					DeviceName:  "test1",
					NUMANode:    0,
					NetDevClass: netdetect.Infiniband,
				},
				{
					Provider:    "ofi+sockets",
					DeviceName:  "test2",
					NUMANode:    0,
					NetDevClass: netdetect.Ether,
				},
			},
			expResult: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					0: {
						{
							Name:        "test1",
							NetDevClass: netdetect.Infiniband,
						},
						{
							Name:        "test2",
							NetDevClass: netdetect.Ether,
						},
					},
					1: {
						{
							Name:        "test0",
							NetDevClass: netdetect.Ether,
						},
					},
				},
			},
		},
		"with device alias": {
			getDevAlias: func(_ context.Context, dev string) (string, error) {
				return dev + "_alias", nil
			},
			input: []*netdetect.FabricScan{
				{
					Provider:    "ofi+sockets",
					DeviceName:  "test0",
					NUMANode:    2,
					NetDevClass: netdetect.Ether,
				},
				{
					Provider:    "ofi+verbs",
					DeviceName:  "test1",
					NUMANode:    1,
					NetDevClass: netdetect.Infiniband,
				},
				{
					Provider:    "ofi+sockets",
					DeviceName:  "test2",
					NUMANode:    1,
					NetDevClass: netdetect.Ether,
				},
			},
			expResult: &NUMAFabric{
				numaMap: map[int][]*FabricInterface{
					1: {
						{
							Name:        "test1",
							NetDevClass: netdetect.Infiniband,
							Domain:      "test1_alias",
						},
						{
							Name:        "test2",
							NetDevClass: netdetect.Ether,
							Domain:      "test2_alias",
						},
					},
					2: {
						{
							Name:        "test0",
							NetDevClass: netdetect.Ether,
							Domain:      "test0_alias",
						},
					},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			log, buf := logging.NewTestLogger(t.Name())
			defer common.ShowBufferOnFailure(t, buf)

			result := NUMAFabricFromScan(context.TODO(), log, tc.input, tc.getDevAlias)

			if diff := cmp.Diff(tc.expResult, result, cmp.AllowUnexported(NUMAFabric{})); diff != "" {
				t.Fatalf("-want, +got:\n%s", diff)
			}
		})
	}
}
