//
// (C) Copyright 2020-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package server

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/daos-stack/daos/src/control/common"
	"github.com/daos-stack/daos/src/control/logging"
	"github.com/daos-stack/daos/src/control/server/engine"
	"github.com/daos-stack/daos/src/control/server/storage/scm"
	"github.com/daos-stack/daos/src/control/system"
)

func TestServer_Instance_createSuperblock(t *testing.T) {
	log, buf := logging.NewTestLogger(t.Name())
	defer ShowBufferOnFailure(t, buf)

	testDir, cleanup := CreateTestDir(t)
	defer cleanup()

	h := NewEngineHarness(log)
	for idx, mnt := range []string{"one", "two"} {
		if err := os.MkdirAll(filepath.Join(testDir, mnt), 0777); err != nil {
			t.Fatal(err)
		}
		cfg := engine.NewConfig().
			WithRank(uint32(idx)).
			WithSystemName(t.Name()).
			WithScmClass("ram").
			WithScmRamdiskSize(1).
			WithScmMountPoint(mnt)
		r := engine.NewRunner(log, cfg)
		msc := &scm.MockSysConfig{
			IsMountedBool: true,
		}
		mp := scm.NewMockProvider(log, nil, msc)
		ei := NewEngineInstance(log, nil, mp, nil, r).
			WithHostFaultDomain(system.MustCreateFaultDomainFromString("/host1"))
		ei.fsRoot = testDir
		if err := h.AddInstance(ei); err != nil {
			t.Fatal(err)
		}
	}

	for _, e := range h.Instances() {
		if err := e.(*EngineInstance).createSuperblock(false); err != nil {
			t.Fatal(err)
		}
	}

	h.started.SetTrue()
	mi := h.instances[0].(*EngineInstance)
	if mi._superblock == nil {
		t.Fatal("instance superblock is nil after createSuperblock()")
	}
	if mi._superblock.System != t.Name() {
		t.Fatalf("expected superblock system name to be %q, got %q", t.Name(), mi._superblock.System)
	}

	for idx, e := range h.Instances() {
		i := e.(*EngineInstance)
		if i._superblock.Rank.Uint32() != uint32(idx) {
			t.Fatalf("instance %d has rank %s (not %d)", idx, i._superblock.Rank, idx)
		}

		AssertEqual(t, i.hostFaultDomain.String(), i._superblock.HostFaultDomain, fmt.Sprintf("instance %d", idx))

		if i == mi {
			continue
		}
		if i._superblock.UUID == mi._superblock.UUID {
			t.Fatal("second instance has same superblock as first")
		}
	}
}
