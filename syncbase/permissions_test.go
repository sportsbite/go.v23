// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncbase_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/security"
	"v.io/v23/security/access"
	wire "v.io/v23/services/syncbase"
	"v.io/v23/syncbase"
	"v.io/v23/syncbase/util"
	"v.io/v23/verror"
	vsecurity "v.io/x/ref/lib/security"
	_ "v.io/x/ref/runtime/factories/roaming"
	tu "v.io/x/ref/services/syncbase/testutil"
	vtestutil "v.io/x/ref/test/testutil"
)

type serviceTest struct {
	f func(ctx *context.T, s syncbase.Service) error
}

type databaseTest struct {
	f func(ctx *context.T, d syncbase.Database) error
}

type batchDatabaseTest struct {
	f func(ctx *context.T, b syncbase.BatchDatabase) error
}

type blobTest struct {
	commit bool
	f      func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error
}

type syncgroupTest struct {
	f func(ctx *context.T, sg syncbase.Syncgroup, c syncbase.Collection) error
}

type collectionTest struct {
	noBatch bool
	f       func(ctx *context.T, d syncbase.Database, c syncbase.Collection) error
}

type rowTest struct {
	noBatch bool
	f       func(ctx *context.T, r syncbase.Row) error
}

func expectNotImplemented(err error) error {
	if err == nil {
		return verror.New(verror.ErrInternal, nil, "unimplemented method returned nil error")
	}
	if verror.ErrorID(err) == verror.ErrNotImplemented.ID {
		return nil
	}
	return err
}

// checkLayer specifies the layer where the primary authorization and
// existence check of the method is carried out.
type checkLayer string

const (
	clService    checkLayer = "S"
	clDatabase   checkLayer = "D"
	clCollection checkLayer = "C"
	clRow        checkLayer = "R"
	clSyncgroup  checkLayer = "G"
)

type securitySpecTest struct {
	layer      interface{}
	name       string
	checkLayer checkLayer
	pattern    string
}

type securitySpecTestGroup struct {
	layer      interface{}
	name       string
	checkLayer checkLayer
	patterns   []string
	mutating   bool
}

// TODO(rogulenko): Add more test groups.
var securitySpecTestGroups = []securitySpecTestGroup{
	// Service tests.
	{
		layer: serviceTest{f: func(ctx *context.T, s syncbase.Service) error {
			_, err := wire.ServiceClient(s.FullName()).DevModeGetTime(ctx)
			return err
		}},
		name:       "service.DevModeGetTime",
		checkLayer: clService,
		patterns:   []string{"A___"},
	},
	{
		layer: serviceTest{f: func(ctx *context.T, s syncbase.Service) error {
			return wire.ServiceClient(s.FullName()).DevModeUpdateVClock(ctx, wire.DevModeUpdateVClockOpts{})
		}},
		name:       "service.DevModeUpdateVClock",
		checkLayer: clService,
		patterns:   []string{"A___"},
		mutating:   true,
	},
	{
		layer: serviceTest{f: func(ctx *context.T, s syncbase.Service) error {
			_, _, err := s.GetPermissions(ctx)
			return err
		}},
		name:       "service.GetPermissions",
		checkLayer: clService,
		patterns:   []string{"A___"},
	},
	{
		layer: serviceTest{f: func(ctx *context.T, s syncbase.Service) error {
			return s.SetPermissions(ctx, tu.DefaultPerms(access.AllTypicalTags(), "root"), "")
		}},
		name:       "service.SetPermissions",
		checkLayer: clService,
		patterns:   []string{"A___"},
		mutating:   true,
	},
	{
		layer: serviceTest{f: func(ctx *context.T, s syncbase.Service) error {
			_, err := s.ListDatabases(ctx)
			return err
		}},
		name:       "service.ListDatabases",
		checkLayer: clService,
		patterns:   []string{"R___"},
	},

	// Database tests.
	{
		layer: serviceTest{f: func(ctx *context.T, s syncbase.Service) error {
			return s.DatabaseForId(wire.Id{"root", "dNew"}, nil).Create(ctx, tu.DefaultPerms(wire.AllDatabaseTags, "root"))
		}},
		name:       "database.Create",
		checkLayer: clService,
		patterns:   []string{"W___"},
		mutating:   true,
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			return d.Destroy(ctx)
		}},
		name:       "database.Destroy",
		checkLayer: clDatabase, // (also Service, depending on specific perms)
		patterns:   []string{"XA__", "A___"},
		mutating:   true,
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := d.Exists(ctx)
			return err
		}},
		name:       "database.Exists",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__", "R___", "W___"},
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, _, err := d.GetPermissions(ctx)
			return err
		}},
		name:       "database.GetPermissions",
		checkLayer: clDatabase,
		patterns:   []string{"XA__"},
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			return d.SetPermissions(ctx, tu.DefaultPerms(wire.AllDatabaseTags, "root"), "")
		}},
		name:       "database.SetPermissions",
		checkLayer: clDatabase,
		patterns:   []string{"XA__"},
		mutating:   true,
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := d.ListCollections(ctx)
			return err
		}},
		name:       "database.ListCollections - nonbatch",
		checkLayer: clDatabase,
		patterns:   []string{"XR__"},
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			ws := d.Watch(ctx, nil, []wire.CollectionRowPattern{util.RowPrefixPattern(wire.Id{"u", "c"}, "")})
			defer ws.Cancel()
			// Advance() will not block indefinitely because at least the root update
			// is always sent for nil resume marker.
			ws.Advance()
			return ws.Err()
		}},
		name:       "database.Watch",
		checkLayer: clDatabase,
		patterns:   []string{"XR__"},
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := d.BeginBatch(ctx, wire.BatchOptions{})
			return err
		}},
		name:       "database.BeginBatch",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := d.GetResumeMarker(ctx)
			return err
		}},
		name:       "database.GetResumeMarker",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
	},
	// TODO(ivanpi): Test Exec.
	// TODO(ivanpi): Test RPC-only methods such as Glob.

	// Batch database tests.
	{
		layer: batchDatabaseTest{f: func(ctx *context.T, b syncbase.BatchDatabase) error {
			return b.Commit(ctx)
		}},
		name:       "database.Commit",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
		mutating:   true,
	},
	{
		layer: batchDatabaseTest{f: func(ctx *context.T, b syncbase.BatchDatabase) error {
			return b.Abort(ctx)
		}},
		name:       "database.Abort",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
		mutating:   true,
	},
	{
		layer: batchDatabaseTest{f: func(ctx *context.T, b syncbase.BatchDatabase) error {
			_, err := b.ListCollections(ctx)
			return err
		}},
		name:       "database.ListCollections - batch",
		checkLayer: clDatabase,
		patterns:   []string{"XR__"},
	},
	// TODO(ivanpi): Test Exec.

	// Conflict resolver tests.
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := wire.DatabaseClient(d.FullName()).GetSchemaMetadata(ctx)
			if verror.ErrorID(err) == verror.ErrNoExist.ID {
				return nil
			}
			return err
		}},
		name:       "database.GetSchemaMetadata",
		checkLayer: clDatabase,
		patterns:   []string{"XR__"},
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			return wire.DatabaseClient(d.FullName()).SetSchemaMetadata(ctx, wire.SchemaMetadata{})
		}},
		name:       "database.SetSchemaMetadata",
		checkLayer: clDatabase,
		patterns:   []string{"XA__"},
		mutating:   true,
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			return wire.DatabaseClient(d.FullName()).PauseSync(ctx)
		}},
		name:       "database.PauseSync",
		checkLayer: clDatabase,
		patterns:   []string{"XA__"},
		mutating:   true,
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			return wire.DatabaseClient(d.FullName()).ResumeSync(ctx)
		}},
		name:       "database.ResumeSync",
		checkLayer: clDatabase,
		patterns:   []string{"XA__"},
		mutating:   true,
	},
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			s, err := wire.DatabaseClient(d.FullName()).StartConflictResolver(ctx)
			if err != nil {
				return err
			}
			cancelTimer := time.AfterFunc(5*time.Second, cancel) // slows down only the allowed case
			defer cancelTimer.Stop()
			s.RecvStream().Advance()
			err = s.RecvStream().Err()
			if verror.ErrorID(err) == verror.ErrCanceled.ID {
				return nil
			}
			return err
		}},
		name:       "database.StartConflictResolver",
		checkLayer: clDatabase,
		patterns:   []string{"XA__"},
		mutating:   true,
	},

	// Blob manager tests.
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := d.CreateBlob(ctx)
			return err
		}},
		name:       "database.CreateBlob",
		checkLayer: clDatabase,
		patterns:   []string{"XW__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: false, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			w, err := blob.Put(ctx)
			if err != nil {
				return err
			}
			if err := w.Send([]byte("foo")); err != nil {
				return err
			}
			return w.Close()
		}},
		name:       "blob.Put",
		checkLayer: clDatabase,
		patterns:   []string{"XW__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: false, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			return blob.Commit(ctx)
		}},
		name:       "blob.Commit",
		checkLayer: clDatabase,
		patterns:   []string{"XW__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			_, err := blob.Size(ctx)
			return err
		}},
		name:       "blob.Size",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			return expectNotImplemented(blob.Delete(ctx))
		}},
		name:       "blob.Delete",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			r, err := blob.Get(ctx, 0)
			if err != nil {
				return err
			}
			defer r.Cancel()
			r.Advance()
			return r.Err()
		}},
		name:       "blob.Get",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			s, err := blob.Fetch(ctx, 0)
			if err != nil {
				return err
			}
			defer s.Cancel()
			s.Advance()
			return s.Err()
		}},
		name:       "blob.Fetch",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			return expectNotImplemented(blob.Pin(ctx))
		}},
		name:       "blob.Pin",
		checkLayer: clDatabase,
		patterns:   []string{"XW__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			return expectNotImplemented(blob.Unpin(ctx))
		}},
		name:       "blob.Unpin",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
		mutating:   true,
	},
	{
		layer: blobTest{commit: true, f: func(ctx *context.T, d syncbase.Database, blob syncbase.Blob) error {
			return expectNotImplemented(blob.Keep(ctx, 0))
		}},
		name:       "blob.Keep",
		checkLayer: clDatabase,
		patterns:   []string{"XX__", "XR__", "XW__", "XA__"},
		mutating:   true,
	},
	// TODO(ivanpi): Test Syncbase-to-Syncbase blob RPCs.

	// Syncgroup manager tests.
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			_, err := d.ListSyncgroups(ctx)
			return err
		}},
		name:       "database.ListSyncgroups",
		checkLayer: clDatabase,
		patterns:   []string{"XR__"},
	},
	{
		layer: collectionTest{noBatch: true, f: func(ctx *context.T, d syncbase.Database, c syncbase.Collection) error {
			return d.SyncgroupForId(wire.Id{"root", "sgNew"}).Create(ctx, wire.SyncgroupSpec{
				Perms:       tu.DefaultPerms(wire.AllSyncgroupTags, "root"),
				Collections: []wire.Id{c.Id()},
			}, wire.SyncgroupMemberInfo{})
		}},
		name:       "syncgroup.Create",
		checkLayer: clDatabase, // (Collection check is not recursive, Database check is sufficient against leaks locally)
		patterns:   []string{"XWR_"},
		mutating:   true,
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, _ syncbase.Collection) error {
			_, err := sg.Join(ctx, "", []string{}, wire.SyncgroupMemberInfo{})
			return err
		}},
		name:       "syncgroup.Join (already joined)",
		checkLayer: clDatabase, // (Collection and Syncgroup check is not recursive, Database check is sufficient against leaks locally)
		patterns:   []string{"XWRR"},
		mutating:   true,
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, _ syncbase.Collection) error {
			return expectNotImplemented(sg.Leave(ctx))
		}},
		name:       "syncgroup.Leave",
		checkLayer: clDatabase,
		patterns:   []string{"XW__"},
		mutating:   true,
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, _ syncbase.Collection) error {
			return expectNotImplemented(sg.Destroy(ctx))
		}},
		name:       "syncgroup.Destroy",
		checkLayer: clDatabase, // (Syncgroup check is not recursive, Database check is sufficient against leaks locally)
		patterns:   []string{"XW_A"},
		mutating:   true,
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, _ syncbase.Collection) error {
			return expectNotImplemented(sg.Eject(ctx, "root:nobody"))
		}},
		name:       "syncgroup.Eject",
		checkLayer: clSyncgroup,
		patterns:   []string{"XX_A"},
		mutating:   true,
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, _ syncbase.Collection) error {
			_, _, err := sg.GetSpec(ctx)
			return err
		}},
		name:       "syncgroup.GetSpec",
		checkLayer: clSyncgroup,
		patterns:   []string{"XX_R"},
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, c syncbase.Collection) error {
			return sg.SetSpec(ctx, wire.SyncgroupSpec{
				Perms:       tu.DefaultPerms(wire.AllSyncgroupTags, "root"),
				Collections: []wire.Id{c.Id()},
			}, "")
		}},
		name:       "syncgroup.SetSpec",
		checkLayer: clSyncgroup,
		patterns:   []string{"XX_A"},
		mutating:   true,
	},
	{
		layer: syncgroupTest{f: func(ctx *context.T, sg syncbase.Syncgroup, _ syncbase.Collection) error {
			_, err := sg.GetMembers(ctx)
			return err
		}},
		name:       "syncgroup.GetMembers",
		checkLayer: clSyncgroup,
		patterns:   []string{"XX_R"},
	},
	// TODO(ivanpi): Test Syncbase-to-Syncbase sync RPCs.

	// Collection tests.
	{
		layer: databaseTest{f: func(ctx *context.T, d syncbase.Database) error {
			return d.CollectionForId(wire.Id{"root", "cNew"}).Create(ctx, tu.DefaultPerms(wire.AllCollectionTags, "root"))
		}},
		name:       "collection.Create",
		checkLayer: clDatabase,
		patterns:   []string{"XW__"},
		mutating:   true,
	},
	{
		layer: collectionTest{f: func(ctx *context.T, _ syncbase.Database, c syncbase.Collection) error {
			return c.Destroy(ctx)
		}},
		name:       "collection.Destroy",
		checkLayer: clCollection, // (also Database, depending on specific perms)
		patterns:   []string{"XXA_", "XA__"},
		mutating:   true,
	},
	{
		layer: collectionTest{f: func(ctx *context.T, _ syncbase.Database, c syncbase.Collection) error {
			_, err := c.Exists(ctx)
			return err
		}},
		name:       "collection.Exists",
		checkLayer: clCollection,
		patterns:   []string{"XXR_", "XXW_", "XXA_", "XR__", "XW__"},
	},
	{
		layer: collectionTest{f: func(ctx *context.T, _ syncbase.Database, c syncbase.Collection) error {
			_, err := c.GetPermissions(ctx)
			return err
		}},
		name:       "collection.GetPermissions",
		checkLayer: clCollection,
		patterns:   []string{"XXA_"},
	},
	{
		layer: collectionTest{f: func(ctx *context.T, _ syncbase.Database, c syncbase.Collection) error {
			return c.SetPermissions(ctx, tu.DefaultPerms(wire.AllCollectionTags, "root"))
		}},
		name:       "collection.SetPermissions",
		checkLayer: clCollection,
		patterns:   []string{"XXA_"},
		mutating:   true,
	},
	{
		layer: collectionTest{f: func(ctx *context.T, _ syncbase.Database, c syncbase.Collection) error {
			ss := c.Scan(ctx, syncbase.Prefix(""))
			defer ss.Cancel()
			ss.Advance()
			return ss.Err()
		}},
		name:       "collection.Scan",
		checkLayer: clCollection,
		patterns:   []string{"XXR_"},
	},
	{
		layer: collectionTest{f: func(ctx *context.T, _ syncbase.Database, c syncbase.Collection) error {
			return c.DeleteRange(ctx, syncbase.Prefix(""))
		}},
		name:       "collection.DeleteRange",
		checkLayer: clCollection,
		patterns:   []string{"XXW_"},
		mutating:   true,
	},
	// TODO(ivanpi): Test RPC-only methods such as Glob. (Note, Glob is non-batch only.)

	// Row tests.
	{
		layer: rowTest{f: func(ctx *context.T, r syncbase.Row) error {
			_, err := r.Exists(ctx)
			return err
		}},
		name:       "row.Exists",
		checkLayer: clRow,
		patterns:   []string{"XXR_", "XXW_"},
	},
	{
		layer: rowTest{f: func(ctx *context.T, r syncbase.Row) error {
			var value string
			return r.Get(ctx, &value)
		}},
		name:       "row.Get",
		checkLayer: clCollection, // (Row check is only meaningful for Exists since Rows don't have separate ACLs)
		patterns:   []string{"XXR_"},
	},
	{
		layer: rowTest{f: func(ctx *context.T, r syncbase.Row) error {
			value := "NCC-1701-D"
			return r.Put(ctx, &value)
		}},
		name:       "row.Put",
		checkLayer: clCollection,
		patterns:   []string{"XXW_"},
		mutating:   true,
	},
	{
		layer: rowTest{f: func(ctx *context.T, r syncbase.Row) error {
			return r.Delete(ctx)
		}},
		name:       "row.Delete",
		checkLayer: clCollection,
		patterns:   []string{"XXW_"},
		mutating:   true,
	},
}

// TestSecuritySpec tests the Syncbase ACL specification. It runs a set of test
// groups.
//
// Each test group describes a Syncbase public method and a list of security
// patterns that should allow a client to call the method.
// A method might be service.GetPermissions and the pattern for it is "A___".
// A security pattern is a string of 4 bytes where each byte is an '_' or
// one of X (resolve), R (read), W (write), A (admin).
// The character index stands for:
// 0 - service ACL
// 1 - database ACL
// 2 - collection ACL
// 3 - syncgroup ACL
// A pattern defines a set of per-ACL permissions required for a client to call
// the method.
//
// For each test group we generate two sets of tests: allowed tests and denied
// tests.
// Allowed tests are generated by splitting the list of security patterns
// of the test group into one pattern per test.
// Denied tests are generated the following way: we take each allowed pattern
// and generate all patterns that differ in one character, preserving all '_'
// (excluding patterns that are a subset of an allowed pattern). This way we
// make sure that every character in a security pattern is important.
// Allowed tests are split in two categories: mutating tests and static tests.
// Mutating tests change the state of Syncbase, static don't.
// Denied tests are always static.
//
// Then we pack all tests into runs, where each run has only one type of tests
// (allowed or denied) and at most one mutating test. For each run we rebuild
// the following Syncbase structure:
// service s; database {a,d}; collection c; row prefix.
//
// For each test inside a run we generate 5 huge ACLs with a record for each
// of the test + one admin record.
func TestSecuritySpec(t *testing.T) {
	deniedTests, staticAllowedTests, mutatingAllowedTests := prepareTests()

	checkDenied := func(t *testing.T, test securitySpecTest, batch bool, err error) {
		if verror.ErrorID(err) != verror.ErrNoAccess.ID && verror.ErrorID(err) != verror.ErrNoExistOrNoAccess.ID {
			tu.Fatalf(t, "test %v (batch: %v) didn't fail with ErrNoAccess or ErrNoExistOrNoAccess error: %v", test, batch, err)
		}
	}
	runTests(t, csAll, checkDenied, deniedTests...)

	checkAllowed := func(t *testing.T, test securitySpecTest, batch bool, err error) {
		if err != nil {
			tu.Fatalf(t, "test %v (batch: %v) failed with non-nil error: %v", test, batch, err)
		}
	}
	runTests(t, csAll, checkAllowed, staticAllowedTests...)
	for _, test := range mutatingAllowedTests {
		runTests(t, csAll, checkAllowed, test)
	}
}

// TestErrorExistLeakSpec tests that the Syncbase RPCs do not leak information
// about database, collection, syncgroup [non]existence to unauthorized clients.
// It uses the test group infrastructure from TestSecuritySpec.
//
// For simplicity and speed, tests are run only with denied tests, avoiding
// mutable tests. Each test group specifies the layer where its primary
// authorization and existence check is carried out in checkLayer.
//
// All tests are run against different configurations of layer existence as
// specified by createSpec. Batch existence is only relevant for tests that
// specify batch usage.
//
// Since the main intent is to verify that ErrNoExistOrNoAccess is returned
// consistently when required, variance in other errors for different RPCs is
// more loosely checked.
func TestErrorExistLeakSpec(t *testing.T) {
	deniedTests, _, _ := prepareTests()

	createSpecs := []string{
		"",
		csDatabase,
		csDatabase + csBatch,
		csDatabase + csCollection,
		csDatabase + csBatch + csCollection,
		csDatabase + csCollection + csRow,
		csDatabase + csBatch + csCollection + csRow,
		csDatabase + csCollection + csSyncgroup,
		csDatabase + csBlob,
		csAll,
	}
	// Race tests run out of memory when all combinations are run. Run on a
	// representative subset instead.
	// TODO(ivanpi): Investigate further. Error encountered is:
	// "ThreadSanitizer: DenseSlabAllocator overflow. Dying."
	if vtestutil.RaceEnabled {
		createSpecs = []string{
			"",
			csDatabase + csBatch,
			csDatabase + csCollection + csSyncgroup,
		}
	}

	for _, createSpec := range createSpecs {
		runTests(t, createSpec, makeErrorChecker(createSpec), deniedTests...)
	}

	// TODO(ivanpi): Consider testing more complex scenarios, such as allowed
	// patterns and different ACLs inside vs outside batch.
}

func prepareTests() (deniedTests, staticAllowedTests, mutatingAllowedTests []securitySpecTest) {
	for _, testGroup := range securitySpecTestGroups {
		// Collect allowed patterns.
		allowedPatterns := map[string]bool{}
		for _, pattern := range testGroup.patterns {
			allowedPatterns[pattern] = true
		}
		// Fill in allowed test slices.
		for pattern := range allowedPatterns {
			test := securitySpecTest{
				layer:      testGroup.layer,
				name:       testGroup.name,
				checkLayer: testGroup.checkLayer,
				pattern:    pattern,
			}
			if testGroup.mutating {
				mutatingAllowedTests = append(mutatingAllowedTests, test)
			} else {
				staticAllowedTests = append(staticAllowedTests, test)
			}
		}

		// Collect denied patterns.
		deniedPatterns := map[string]bool{}
		for _, pattern := range testGroup.patterns {
			for i := 0; i < len(pattern); i++ {
				if pattern[i] != '_' {
					patternBytes := []byte(pattern)
					for _, c := range []byte{'_', 'X', 'R', 'W', 'A'} {
						patternBytes[i] = c
						if !isSubsetOfAny(string(patternBytes), allowedPatterns) {
							deniedPatterns[string(patternBytes)] = true
						}
					}
				}
			}
		}
		// Fill in denied test slice.
		for pattern := range deniedPatterns {
			test := securitySpecTest{
				layer:      testGroup.layer,
				name:       testGroup.name,
				checkLayer: testGroup.checkLayer,
				pattern:    pattern,
			}
			deniedTests = append(deniedTests, test)
		}
	}
	return
}

func isSubsetOf(subset, superset string) bool {
	for i := 0; i < len(superset); i++ {
		if superset[i] != '_' && superset[i] != subset[i] {
			return false
		}
	}
	// All perms in subset correspond to the same perm or '_' in superset.
	return true
}

func isSubsetOfAny(needle string, haystack map[string]bool) bool {
	for p := range haystack {
		if isSubsetOf(needle, p) {
			return true
		}
	}
	return false
}

const (
	// createSpec is a concatenation of zero or more of the constants below,
	// describing which layers should exist for a given test. Layer dependencies
	// are not automatically handled - if a layer requires another layer, both
	// layers must be specified.
	csDatabase   = "D"
	csBatch      = "T"
	csCollection = "C"
	csRow        = "R"
	csSyncgroup  = "G"
	csBlob       = "B"
	csAll        = csDatabase + csBatch + csCollection + csRow + csSyncgroup + csBlob
)

type checkerFunc func(t *testing.T, test securitySpecTest, batch bool, err error)

// runTest creates the Syncbase hierarchy according to the provided createSpec
// (see TestErrorExistLeakSpec comment for description) and runs the provided
// set of tests against it, running checkerFunc
func runTests(t *testing.T, createSpec string, check checkerFunc, tests ...securitySpecTest) {
	// Create permissions.
	servicePerms := makePerms("root:admin", 0, access.AllTypicalTags(), tests...)
	databasePerms := makePerms("root:admin", 1, wire.AllDatabaseTags, tests...)
	collectionPerms := makePerms("root:admin", 2, wire.AllCollectionTags, tests...)
	sgPerms := makePerms("root:admin", 3, wire.AllSyncgroupTags, tests...)

	// Create service/database/collection/row with permissions above.
	ctx, adminCtx, sName, rootp, cleanup := tu.SetupOrDieCustom("admin", "server", nil)
	defer cleanup()
	// Service
	s := syncbase.NewService(sName)
	if err := s.SetPermissions(adminCtx, servicePerms, ""); err != nil {
		tu.Fatalf(t, "s.SetPermissions failed: %v", err)
	}
	// Database
	d := s.DatabaseForId(wire.Id{"root", "d"}, nil)
	if strings.Contains(createSpec, csDatabase) {
		if err := d.Create(adminCtx, databasePerms); err != nil {
			tu.Fatalf(t, "d.Create failed: %v", err)
		}
	}
	// Collection
	c := d.CollectionForId(wire.Id{"root:admin", "c"})
	if strings.Contains(createSpec, csCollection) {
		if err := c.Create(adminCtx, collectionPerms); err != nil {
			tu.Fatalf(t, "c.Create failed: %v", err)
		}
	}
	// Row
	r := c.Row("prefix")
	if strings.Contains(createSpec, csRow) {
		if err := r.Put(adminCtx, "value"); err != nil {
			tu.Fatalf(t, "r.Put failed: %v", err)
		}
	}
	// Syncgroup
	sg := d.SyncgroupForId(wire.Id{"root:admin", "sg"})
	if strings.Contains(createSpec, csSyncgroup) {
		sgSpec := wire.SyncgroupSpec{
			Perms:       sgPerms,
			Collections: []wire.Id{c.Id()},
		}
		if err := sg.Create(adminCtx, sgSpec, wire.SyncgroupMemberInfo{}); err != nil {
			tu.Fatalf(t, "sg.Create failed: %v", err)
		}
	}
	// Blob
	blobCommitted := d.Blob("_fake_committed")
	blobUncommitted := d.Blob("_fake_uncommitted")
	if strings.Contains(createSpec, csBlob) {
		var err error
		if blobCommitted, err = d.CreateBlob(adminCtx); err != nil {
			tu.Fatalf(t, "d.CreateBlob failed: %v", err)
		}
		if err = blobCommitted.Commit(adminCtx); err != nil {
			tu.Fatalf(t, "blobCommitted.Commit failed: %v", err)
		}
		if blobUncommitted, err = d.CreateBlob(adminCtx); err != nil {
			tu.Fatalf(t, "d.CreateBlob failed: %v", err)
		}
	}
	// Batch
	dBatch := syncbase.NewNonexistentBatchForTest(d)
	if strings.Contains(createSpec, csBatch) {
		var err error
		if dBatch, err = d.BeginBatch(adminCtx, wire.BatchOptions{}); err != nil {
			tu.Fatalf(t, "d.BeginBatch failed: %v", err)
		}
		defer dBatch.Abort(adminCtx)
	}
	cBatch := dBatch.CollectionForId(c.Id())
	rBatch := cBatch.Row(r.Key())

	// Verify tests.
	for i, test := range tests {
		if filteredPattern := filterPermsPattern(createSpec, test.pattern); filteredPattern != test.pattern {
			// Skip tests whose ACL patterns are unsatisfiable because the
			// corresponding layers are missing in createSpec.
			continue
		}

		// Note, we run collection and row tests twice against the same hierarchy
		// (batch and non-batch) even if they are mutable. This should be safe
		// thanks to batch isolation.
		clientCtx, cancelClientCtx := context.WithCancel(tu.NewCtx(ctx, rootp, fmt.Sprintf("client%d", i)))
		switch layer := test.layer.(type) {
		case serviceTest:
			check(t, test, false, layer.f(clientCtx, s))
		case databaseTest:
			check(t, test, false, layer.f(clientCtx, d))
		case batchDatabaseTest:
			check(t, test, true, layer.f(clientCtx, dBatch))
		case blobTest:
			if layer.commit {
				check(t, test, false, layer.f(clientCtx, d, blobCommitted))
			} else {
				check(t, test, false, layer.f(clientCtx, d, blobUncommitted))
			}
		case syncgroupTest:
			check(t, test, false, layer.f(clientCtx, sg, c))
		case collectionTest:
			if !layer.noBatch {
				check(t, test, true, layer.f(clientCtx, d, cBatch))
			}
			check(t, test, false, layer.f(clientCtx, d, c))
		case rowTest:
			if !layer.noBatch {
				check(t, test, true, layer.f(clientCtx, rBatch))
			}
			check(t, test, false, layer.f(clientCtx, r))
		default:
			tu.Fatalf(t, "test %v: unknown test type", test)
		}
		cancelClientCtx()
	}
}

// filterPermsPattern filters out (sets to '_') permissions in the pattern that
// cannot exist because the corresponding layer doesn't exist in createSpec.
func filterPermsPattern(createSpec, pattern string) string {
	patternBytes := []byte(pattern)
	if !strings.Contains(createSpec, csDatabase) {
		patternBytes[1] = '_'
	}
	if !strings.Contains(createSpec, csCollection) {
		patternBytes[2] = '_'
	}
	if !strings.Contains(createSpec, csSyncgroup) {
		patternBytes[3] = '_'
	}
	return string(patternBytes)
}

// makePerms creates a perms object with permissions for each test.
// For each test we add a blessing pattern "root:clientXX" with a tag
// corresponding test.pattern[index].
func makePerms(admin string, index int, allowTags []access.Tag, tests ...securitySpecTest) access.Permissions {
	perms := tu.DefaultPerms(allowTags, admin)
	tagsMap := map[byte]string{
		'X': "Resolve",
		'R': "Read",
		'W': "Write",
		'A': "Admin",
	}
	for i, t := range tests {
		if tag, ok := tagsMap[t.pattern[index]]; ok {
			perms.Add(security.BlessingPattern(fmt.Sprintf("root:client%d", i)), tag)
		}
	}
	return util.FilterTags(perms, allowTags...)
}

// makeErrorChecker builds a checkerFunc for TestErrorExistLeakSpec for the
// given createSpec.
func makeErrorChecker(createSpec string) checkerFunc {
	return func(t *testing.T, test securitySpecTest, batch bool, gotErr error) {
		filteredPattern := filterPermsPattern(createSpec, test.pattern)
		if filteredPattern != test.pattern {
			tu.Fatalf(t, "test didn't filter out pattern %s invalid for createSpec %s", test.pattern, createSpec)
		}

		assertErrorOneOf := func(wantErrs ...verror.ID) {
			for _, wantErr := range wantErrs {
				if verror.ErrorID(gotErr) == wantErr {
					return
				}
			}
			tu.Fatalf(t, "setup %s, test %v (batch: %v, filtered perms: %s) didn't fail with one of %v, got error: %v", createSpec, test, batch, filteredPattern, wantErrs, gotErr)
		}

		checkPerms := func(patterns ...string) bool {
			for _, p := range patterns {
				if isSubsetOf(filteredPattern, p) {
					return true
				}
			}
			return false
		}

		if test.checkLayer == clService {
			// Service always exists, its methods only fail with ErrNoAccess.
			assertErrorOneOf(verror.ErrNoAccess.ID)
			return
		}

		if !checkPerms("XX__", "XR__", "XW__", "XA__", "R___", "W___") {
			// Caller is not allowed to know if Database exists, so methods on
			// Database and above fail with ErrNoExistOrNoAccess.
			assertErrorOneOf(verror.ErrNoExistOrNoAccess.ID)
			return
		}
		if !strings.Contains(createSpec, csDatabase) {
			// Database does not exist, so methods on Database and above fail with
			// ErrNoExist.
			// Exists and Destroy (idempotent case) return nil error.
			assertErrorOneOf(verror.ErrorID(nil), verror.ErrNoExist.ID)
			return
		}
		if batch {
			if !strings.Contains(createSpec, csBatch) {
				// Client is allowed to know that Database exists, but the batch handle
				// is invalid.
				assertErrorOneOf(wire.ErrUnknownBatch.ID)
				return
			}
		}
		if test.checkLayer == clDatabase {
			// Database exists and client is allowed to know it, so its methods fail
			// with ErrNoAccess.
			// Some of the Syncgroup Create and Join methods have patterns that do
			// pass the Database ACL, but fail on Collection existence check or
			// Syncgroup remote join check (after failing Syncgroup existence check).
			assertErrorOneOf(verror.ErrNoAccess.ID, verror.ErrNoExist.ID, wire.ErrSyncgroupJoinFailed.ID)
			return
		}
		if !checkPerms("X___") {
			// Client cannot resolve through Service, so methods above Database fail
			// with ErrNoAccess.
			assertErrorOneOf(verror.ErrNoAccess.ID)
			return
		}

		if test.checkLayer == clSyncgroup {
			// Local Syncgroup method branch:

			if !checkPerms("XX_R", "XX_A", "XR__", "XW__") {
				// Caller is not allowed to know if Syncgroup exists, so methods on
				// Syncgroup fail with ErrNoExistOrNoAccess.
				assertErrorOneOf(verror.ErrNoExistOrNoAccess.ID)
				return
			}
			if !strings.Contains(createSpec, csSyncgroup) {
				// Syncgroup does not exist, so methods on Syncgroup fail with
				// ErrNoExist.
				assertErrorOneOf(verror.ErrNoExist.ID)
				return
			}
			// Syncgroup exists and client is allowed to know it, so its methods fail
			// with ErrNoAccess.
			assertErrorOneOf(verror.ErrNoAccess.ID)
			return
		}

		// Collection and Row method branch:

		if !checkPerms("XXR_", "XXW_", "XXA_", "XR__", "XW__") {
			// Caller is not allowed to know if Collection exists, so methods on
			// Collection and above fail with ErrNoExistOrNoAccess.
			assertErrorOneOf(verror.ErrNoExistOrNoAccess.ID)
			return
		}
		if !strings.Contains(createSpec, csCollection) {
			// Collection does not exist, so methods on Collection and above fail
			// with ErrNoExist.
			assertErrorOneOf(verror.ErrorID(nil), verror.ErrNoExist.ID)
			return
		}
		if test.checkLayer == clCollection {
			// Collection exists and client is allowed to know it, so its methods fail
			// with ErrNoAccess.
			assertErrorOneOf(verror.ErrNoAccess.ID)
			return
		}

		if !checkPerms("XX__") {
			// Client cannot resolve through Database, so methods above Collection
			// fail with ErrNoAccess.
			assertErrorOneOf(verror.ErrNoAccess.ID)
			return
		}

		if !checkPerms("XXR_", "XXW_") {
			// Caller is not allowed to know if Row exists, so methods on Row fail
			// with ErrNoExistOrNoAccess.
			assertErrorOneOf(verror.ErrNoExistOrNoAccess.ID)
			return
		}
		// Note, since Exists is currently the only Row-checked method, the two
		// checks below are unnecessary.
		/*
			if !strings.Contains(createSpec, csRow)  {
				// Row does not exist, so methods on Row fail with ErrNoExist.
				assertErrorOneOf(verror.ErrNoExist.ID)
				return
			}
			if test.checkLayer == clRow {
				// Row exists and client is allowed to know it, so its methods fail
				// with ErrNoAccess.
				assertErrorOneOf(verror.ErrNoAccess.ID)
				return
			}
		*/

		// Blob layer has no permissions of its own, it is covered by Database
		// permissions.

		tu.Fatalf(t, "unknown permission check layer: %v", test.checkLayer)
	}
}

var (
	inferBlessingsOk = []struct {
		exts     []string
		wantApp  string
		wantUser string
	}{
		{
			[]string{"o:angrybirds:alice"},
			"root:o:angrybirds", "root:o:angrybirds:alice",
		},
		{
			[]string{"foo", "o:angrybirds:alice", "x:baz"},
			"root:o:angrybirds", "root:o:angrybirds:alice",
		},
		{
			[]string{"foo", "u:alice", "o:angrybirds:alice:device:phone", "x:baz"},
			"root:o:angrybirds", "root:o:angrybirds:alice",
		},
		{
			[]string{"foo", "u:bob", "o:todos:dave:friend:alice", "u:carol", "x:baz"},
			"root:o:todos", "root:o:todos:dave",
		},
		{
			[]string{"foo", "u:bob", "o:todos:dave:friend:alice", "u:carol", "o:todos:dave:device:phone", "x:baz"},
			"root:o:todos", "root:o:todos:dave",
		},
		{
			[]string{"u:bob"},
			"...", "root:u:bob",
		},
		{
			[]string{"foo", "u:bob", "x:baz"},
			"...", "root:u:bob",
		},
		{
			[]string{"foo", "u:bob:angrybirds", "u:bob:todos:phone", "x:baz"},
			"...", "root:u:bob",
		},
	}

	inferBlessingsFail = [][]string{
		{},
		{"foo", "x:baz"},
		{"foo", "u:bob", "o:todos:dave:friend:alice", "o:angrybirds:dave", "o:todos:dave:device:phone", "x:baz"},
		{"foo", "u:bob", "o:todos:dave:friend:alice", "o:todos:fred"},
		{"foo", "u:bob:angrybirds", "u:bob:todos:phone", "u:carol", "u:dave", "x:baz"},
	}

	dummyPerms = access.Permissions{}.Add("...", string(access.Read)).Add("root:u:nobody", string(access.Admin))
)

// TestIdBlessingInfer tests that Database/Collection/Syncgroup getter variants
// taking a context and name correctly infer the Id blessing from the context,
// failing when ambiguous.
func TestIdBlessingInfer(t *testing.T) {
	anchorPerms := access.Permissions{}.Add("root:u:admin", string(access.Admin)).Add("root", string(access.Resolve), string(access.Read), string(access.Write))
	rootCtx, adminCtx, sName, rootp, cleanup := tu.SetupOrDieCustom("u:admin", "server", anchorPerms)
	defer cleanup()
	s := syncbase.NewService(sName)
	rd := s.Database(adminCtx, "anchor", nil)
	if err := rd.Create(adminCtx, anchorPerms); err != nil {
		t.Fatalf("rd.Create() failed: %v", err)
	}
	rc := rd.Collection(adminCtx, "anchor")
	if err := rc.Create(adminCtx, util.FilterTags(anchorPerms, wire.AllCollectionTags...)); err != nil {
		t.Fatalf("rc.Create() failed: %v", err)
	}
	sgSpec := wire.SyncgroupSpec{
		Description: "dummy syncgroup",
		Collections: []wire.Id{rc.Id()},
		Perms:       dummyPerms,
	}
	sgInfo := wire.SyncgroupMemberInfo{
		SyncPriority: 42,
	}

	for i, test := range inferBlessingsOk {
		ctx := forkContextMultiPattern(rootCtx, rootp, test.exts)
		nameSuffix := fmt.Sprintf("ok_%02d", i)

		d := s.Database(ctx, "db_"+nameSuffix, nil)
		if got, want := d.Id(), (wire.Id{Blessing: test.wantApp, Name: "db_" + nameSuffix}); got != want {
			t.Errorf("blessings %v: expected inferred database id %v, got %v", test.exts, want, got)
		}
		if err := d.Create(ctx, dummyPerms); err != nil {
			t.Errorf("blessings %v: d.Create() failed: %v", test.exts, err)
		}

		c := rd.Collection(ctx, "cx_"+nameSuffix)
		if got, want := c.Id(), (wire.Id{Blessing: test.wantUser, Name: "cx_" + nameSuffix}); got != want {
			t.Errorf("blessings %v: expected inferred collection id %v, got %v", test.exts, want, got)
		}
		if err := c.Create(ctx, dummyPerms); err != nil {
			t.Errorf("blessings %v: c.Create() failed: %v", test.exts, err)
		}

		sg := rd.Syncgroup(ctx, "sg_"+nameSuffix)
		if got, want := sg.Id(), (wire.Id{Blessing: test.wantUser, Name: "sg_" + nameSuffix}); got != want {
			t.Errorf("blessings %v: expected inferred syncgroup id %v, got %v", test.exts, want, got)
		}
		if err := sg.Create(ctx, sgSpec, sgInfo); err != nil {
			t.Errorf("blessings %v: sg.Create() failed: %v", test.exts, err)
		}
	}

	for i, exts := range inferBlessingsFail {
		ctx := forkContextMultiPattern(rootCtx, rootp, exts)
		nameSuffix := fmt.Sprintf("notok_%02d", i)

		d := s.Database(ctx, "db_"+nameSuffix, nil)
		if got, want := d.Id(), (wire.Id{Blessing: "$", Name: "db_" + nameSuffix}); got != want {
			t.Errorf("blessings %v: expected inferred database id %v, got %v", exts, want, got)
		}
		if err := d.Create(ctx, dummyPerms); verror.ErrorID(err) != wire.ErrInvalidName.ID {
			t.Errorf("blessings %v: d.Create() should have failed with ErrInvalidName, got: %v", exts, err)
		}

		c := rd.Collection(ctx, "cx_"+nameSuffix)
		if got, want := c.Id(), (wire.Id{Blessing: "$", Name: "cx_" + nameSuffix}); got != want {
			t.Errorf("blessings %v: expected inferred collection id %v, got %v", exts, want, got)
		}
		if err := c.Create(ctx, dummyPerms); verror.ErrorID(err) != wire.ErrInvalidName.ID {
			t.Errorf("blessings %v: c.Create() should have failed with ErrInvalidName, got: %v", exts, err)
		}

		sg := rd.Syncgroup(ctx, "sg_"+nameSuffix)
		if got, want := sg.Id(), (wire.Id{Blessing: "$", Name: "sg_" + nameSuffix}); got != want {
			t.Errorf("blessings %v: expected inferred syncgroup id %v, got %v", exts, want, got)
		}
		if err := sg.Create(ctx, sgSpec, sgInfo); verror.ErrorID(err) != wire.ErrInvalidName.ID {
			t.Errorf("blessings %v: sg.Create() should have failed with ErrInvalidName, got: %v", exts, err)
		}
	}
}

// TestIdBlessingInfer tests that Database/Collection Create() calls correctly
// infer default perms from the context when nil perms are passed in, failing
// when ambiguous.
func TestCreatePermsInfer(t *testing.T) {
	anchorPerms := access.Permissions{}.Add("root:u:admin", string(access.Admin)).Add("root", string(access.Resolve), string(access.Read), string(access.Write))
	rootCtx, adminCtx, sName, rootp, cleanup := tu.SetupOrDieCustom("u:admin", "server", anchorPerms)
	defer cleanup()
	s := syncbase.NewService(sName)
	rd := s.Database(adminCtx, "anchor", nil)
	if err := rd.Create(adminCtx, anchorPerms); err != nil {
		t.Fatalf("rd.Create() failed: %v", err)
	}

	for i, test := range inferBlessingsOk {
		ctx := forkContextMultiPattern(rootCtx, rootp, test.exts)
		nameSuffix := fmt.Sprintf("ok_%02d", i)

		d := s.DatabaseForId(wire.Id{Blessing: "root", Name: "db_" + nameSuffix}, nil)
		if err := d.Create(ctx, nil); err != nil {
			t.Errorf("blessings %v: d.Create() failed: %v", test.exts, err)
		}
		if got, _, err := d.GetPermissions(ctx); err != nil {
			t.Fatalf("d.GetPermissions() failed: %v", err)
		} else if want := tu.DefaultPerms(wire.AllDatabaseTags, test.wantUser); !reflect.DeepEqual(got, want) {
			t.Errorf("blessings %v: expected inferred database perms %v, got %v", test.exts, want, got)
		}

		c := rd.CollectionForId(wire.Id{Blessing: "root", Name: "cx_" + nameSuffix})
		if err := c.Create(ctx, nil); err != nil {
			t.Errorf("blessings %v: c.Create() failed: %v", test.exts, err)
		}
		if got, err := c.GetPermissions(ctx); err != nil {
			t.Fatalf("c.GetPermissions() failed: %v", err)
		} else if want := tu.DefaultPerms(wire.AllCollectionTags, test.wantUser); !reflect.DeepEqual(got, want) {
			t.Errorf("blessings %v: expected inferred collection perms %v, got %v", test.exts, want, got)
		}
	}

	for i, exts := range inferBlessingsFail {
		if len(exts) == 0 {
			continue
		}
		ctx := forkContextMultiPattern(rootCtx, rootp, exts)
		nameSuffix := fmt.Sprintf("notok_%02d", i)

		d := s.DatabaseForId(wire.Id{Blessing: "root", Name: "db_" + nameSuffix}, nil)
		if err := d.Create(ctx, nil); verror.ErrorID(err) != wire.ErrInferDefaultPermsFailed.ID {
			t.Errorf("blessings %v: d.Create() should have failed with ErrInferDefaultPermsFailed, got: %v", exts, err)
		}
		if err := d.Create(ctx, dummyPerms); err != nil {
			t.Errorf("blessings %v: d.Create() with explicit perms failed: %v", exts, err)
		}

		c := rd.CollectionForId(wire.Id{Blessing: "root", Name: "cx_" + nameSuffix})
		if err := c.Create(ctx, nil); verror.ErrorID(err) != wire.ErrInferDefaultPermsFailed.ID {
			t.Errorf("blessings %v: c.Create() should have failed with ErrInferDefaultPermsFailed, got: %v", exts, err)
		}
		if err := c.Create(ctx, dummyPerms); err != nil {
			t.Errorf("blessings %v: c.Create() with explicit perms failed: %v", exts, err)
		}
	}
}

func forkContextMultiPattern(rootCtx *context.T, rootp security.Principal, extensions []string) *context.T {
	p, _ := vsecurity.NewPrincipal()
	db, _ := rootp.BlessingStore().Default()
	security.AddToRoots(p, db)
	bs := make([]security.Blessings, 0, len(extensions))
	for _, ext := range extensions {
		b, _ := rootp.Bless(p.PublicKey(), db, ext, security.UnconstrainedUse())
		bs = append(bs, b)
	}
	b, _ := security.UnionOfBlessings(bs...)
	p.BlessingStore().SetDefault(b)
	p.BlessingStore().Set(b, "...")
	ctx, _ := v23.WithPrincipal(rootCtx, p)
	return ctx
}

// TestCreateIdBlessingEnforce tests that only clients matching the blessing
// pattern in the Id are allowed to create a Database/Collection/Syncgroup.
// This requirement is waived for admin clients for Database, but not Collection
// or Syncgroup creation.
func TestCreateIdBlessingEnforce(t *testing.T) {
	anchorPerms := access.Permissions{}.Add("root:u:admin", string(access.Admin)).Add("root", string(access.Resolve), string(access.Read), string(access.Write))
	rootCtx, adminCtx, sName, rootp, cleanup := tu.SetupOrDieCustom("u:admin", "server", anchorPerms)
	defer cleanup()
	clientCtx := tu.NewCtx(rootCtx, rootp, "u:client")
	s := syncbase.NewService(sName)
	rd := s.Database(adminCtx, "anchor", nil)
	if err := rd.Create(adminCtx, anchorPerms); err != nil {
		t.Fatalf("rd.Create() failed: %v", err)
	}
	rc := rd.Collection(adminCtx, "anchor")
	if err := rc.Create(adminCtx, util.FilterTags(anchorPerms, wire.AllCollectionTags...)); err != nil {
		t.Fatalf("rc.Create() failed: %v", err)
	}
	sgSpec := wire.SyncgroupSpec{
		Description: "dummy syncgroup",
		Collections: []wire.Id{rc.Id()},
		Perms:       dummyPerms,
	}
	sgInfo := wire.SyncgroupMemberInfo{
		SyncPriority: 42,
	}

	patternsOk := []string{"...", "root", "root:u", "root:u:client", "root:u:client:$"}
	patternsFail := []string{"root:$", "root:u:$", "root:u:admin", "root:u:client:phone", "foobar"}

	for i, pattern := range patternsOk {
		nameSuffix := fmt.Sprintf("ok_%02d", i)

		d := s.DatabaseForId(wire.Id{Blessing: pattern, Name: "db_" + nameSuffix}, nil)
		if err := d.Create(clientCtx, dummyPerms); err != nil {
			t.Errorf("blessing %q: d.Create() failed: %v", pattern, err)
		}

		c := rd.CollectionForId(wire.Id{Blessing: pattern, Name: "cx_" + nameSuffix})
		if err := c.Create(clientCtx, dummyPerms); err != nil {
			t.Errorf("blessing %q: c.Create() failed: %v", pattern, err)
		}

		sg := rd.SyncgroupForId(wire.Id{Blessing: pattern, Name: "sg_" + nameSuffix})
		if err := sg.Create(clientCtx, sgSpec, sgInfo); err != nil {
			t.Errorf("blessing %q: sg.Create() failed: %v", pattern, err)
		}
	}

	for i, pattern := range patternsFail {
		nameSuffix := fmt.Sprintf("notok_%02d", i)

		d := s.DatabaseForId(wire.Id{Blessing: pattern, Name: "db_" + nameSuffix}, nil)
		if err := d.Create(clientCtx, dummyPerms); verror.ErrorID(err) != wire.ErrUnauthorizedCreateId.ID {
			t.Errorf("blessing %q: d.Create() should have failed with ErrUnauthorizedCreateId, got: %v", pattern, err)
		}

		c := rd.CollectionForId(wire.Id{Blessing: pattern, Name: "cx_" + nameSuffix})
		if err := c.Create(clientCtx, dummyPerms); verror.ErrorID(err) != wire.ErrUnauthorizedCreateId.ID {
			t.Errorf("blessing %q: c.Create() should have failed with ErrUnauthorizedCreateId, got: %v", pattern, err)
		}

		sg := rd.SyncgroupForId(wire.Id{Blessing: pattern, Name: "sg_" + nameSuffix})
		if err := sg.Create(clientCtx, sgSpec, sgInfo); verror.ErrorID(err) != wire.ErrUnauthorizedCreateId.ID {
			t.Errorf("blessing %q: sg.Create() should have failed with ErrUnauthorizedCreateId, got: %v", pattern, err)
		}
	}

	// Give full admin permissions to the client.
	clientAdminPerms := anchorPerms.Copy().Add("root:u:client", string(access.Admin))
	if err := s.SetPermissions(adminCtx, clientAdminPerms, ""); err != nil {
		t.Fatalf("s.SetPermissions() failed: %v", err)
	}
	if err := rd.SetPermissions(adminCtx, clientAdminPerms, ""); err != nil {
		t.Fatalf("rd.SetPermissions() failed: %v", err)
	}
	if err := rc.SetPermissions(adminCtx, util.FilterTags(clientAdminPerms, wire.AllCollectionTags...)); err != nil {
		t.Fatalf("rc.SetPermissions() failed: %v", err)
	}

	// Admin permissions override the Id blessing enforcement for databases, but
	// not for collections or syncgroups.
	for i, pattern := range patternsFail {
		nameSuffix := fmt.Sprintf("admin_%02d", i)

		d := s.DatabaseForId(wire.Id{Blessing: pattern, Name: "db_" + nameSuffix}, nil)
		if err := d.Create(clientCtx, dummyPerms); err != nil {
			t.Errorf("blessing %q: admin d.Create() failed: %v", pattern, err)
		}

		c := rd.CollectionForId(wire.Id{Blessing: pattern, Name: "cx_" + nameSuffix})
		if err := c.Create(clientCtx, dummyPerms); verror.ErrorID(err) != wire.ErrUnauthorizedCreateId.ID {
			t.Errorf("blessing %q: admin c.Create() should have failed with ErrUnauthorizedCreateId, got: %v", pattern, err)
		}

		sg := rd.SyncgroupForId(wire.Id{Blessing: pattern, Name: "sg_" + nameSuffix})
		if err := sg.Create(clientCtx, sgSpec, sgInfo); verror.ErrorID(err) != wire.ErrUnauthorizedCreateId.ID {
			t.Errorf("blessing %q: admin sg.Create() should have failed with ErrUnauthorizedCreateId, got: %v", pattern, err)
		}
	}
}

// TestPermsValidation tests that all operations that create or modify
// permissions properly validate the new permissions.
func TestPermsValidation(t *testing.T) {
	ctx, sName, cleanup := tu.SetupOrDie(nil)
	defer cleanup()
	s := syncbase.NewService(sName)
	rd := tu.CreateDatabase(t, ctx, s, "anchor", nil)
	rc := tu.CreateCollection(t, ctx, rd, "anchor")
	rsg := tu.CreateSyncgroup(t, ctx, rd, rc, "anchor", "anchor syncgroup")

	permsEmpty := access.Permissions{}
	permsNoAdmin := access.Permissions{}.Add("root:o:app:client", string(access.Read))
	permsXRWAD := access.Permissions{}.Add("root:o:app:client", access.TagStrings(access.AllTypicalTags()...)...)
	permsXRWA := access.Permissions{}.Add("root:o:app:client", access.TagStrings(wire.AllDatabaseTags...)...)
	permsRWA := access.Permissions{}.Add("root:o:app:client", access.TagStrings(wire.AllCollectionTags...)...)
	permsRA := access.Permissions{}.Add("root:o:app:client", access.TagStrings(wire.AllSyncgroupTags...)...)
	permsAdminOnly := access.Permissions{}.Add("root:o:app:client", string(access.Admin))
	permsComplex := permsRA.Copy().Add("root:u:client", access.TagStrings(access.Read, access.Admin)...).Blacklist("root:o:app:client:phone", string(access.Admin))

	testPermsValidationOp(t, "d.Create()",
		[]access.Permissions{nil /* infer */, permsAdminOnly, permsComplex, permsRA, permsRWA, permsXRWA},
		[]access.Permissions{permsEmpty, permsNoAdmin, permsXRWAD},
		func(nameSuffix string, perms access.Permissions) error {
			d := s.DatabaseForId(wire.Id{Blessing: "root", Name: "db_" + nameSuffix}, nil)
			return d.Create(ctx, perms)
		})

	testPermsValidationOp(t, "c.Create()",
		[]access.Permissions{nil /* infer */, permsAdminOnly, permsComplex, permsRA, permsRWA},
		[]access.Permissions{permsEmpty, permsNoAdmin, permsXRWA, permsXRWAD},
		func(nameSuffix string, perms access.Permissions) error {
			c := rd.CollectionForId(wire.Id{Blessing: "root", Name: "cx_" + nameSuffix})
			return c.Create(ctx, perms)
		})

	// Syncgroup ACL must allow Read access to client.
	testPermsValidationOp(t, "sg.Create()",
		[]access.Permissions{permsComplex, permsRA},
		[]access.Permissions{nil, permsEmpty, permsNoAdmin, permsAdminOnly, permsRWA, permsXRWA, permsXRWAD},
		func(nameSuffix string, perms access.Permissions) error {
			sg := rd.SyncgroupForId(wire.Id{Blessing: "root", Name: "sg_" + nameSuffix})
			return sg.Create(ctx, wire.SyncgroupSpec{
				Collections: []wire.Id{rc.Id()},
				Perms:       perms,
			}, wire.SyncgroupMemberInfo{SyncPriority: 42})
		})

	testPermsValidationOp(t, "s.SetPermissions()",
		[]access.Permissions{permsAdminOnly, permsComplex, permsRA, permsRWA, permsXRWA, permsXRWAD},
		[]access.Permissions{nil, permsEmpty, permsNoAdmin},
		func(_ string, perms access.Permissions) error {
			return s.SetPermissions(ctx, perms, "")
		})

	testPermsValidationOp(t, "d.SetPermissions()",
		[]access.Permissions{permsAdminOnly, permsComplex, permsRA, permsRWA, permsXRWA},
		[]access.Permissions{nil, permsEmpty, permsNoAdmin, permsXRWAD},
		func(_ string, perms access.Permissions) error {
			return rd.SetPermissions(ctx, perms, "")
		})

	testPermsValidationOp(t, "c.SetPermissions()",
		[]access.Permissions{permsAdminOnly, permsComplex, permsRA, permsRWA},
		[]access.Permissions{nil, permsEmpty, permsNoAdmin, permsXRWA, permsXRWAD},
		func(_ string, perms access.Permissions) error {
			return rc.SetPermissions(ctx, perms)
		})

	// New syncgroup ACL must allow Read access to client.
	testPermsValidationOp(t, "sg.SetSpec()",
		[]access.Permissions{permsComplex, permsRA},
		[]access.Permissions{nil, permsEmpty, permsNoAdmin, permsAdminOnly, permsRWA, permsXRWA, permsXRWAD},
		func(_ string, perms access.Permissions) error {
			return rsg.SetSpec(ctx, wire.SyncgroupSpec{
				Collections: []wire.Id{rc.Id()},
				Perms:       perms,
			}, "")
		})
}

func testPermsValidationOp(t *testing.T, desc string, validPerms, invalidPerms []access.Permissions, op func(nameSuffix string, perms access.Permissions) error) {
	for i, p := range validPerms {
		nameSuffix := fmt.Sprintf("ok_%02d", i)
		if err := op(nameSuffix, p); err != nil {
			t.Errorf("perms %v: %s failed: %v", p, desc, err)
		}
	}

	for i, p := range invalidPerms {
		nameSuffix := fmt.Sprintf("notok_%02d", i)
		if err := op(nameSuffix, p); err == nil {
			t.Errorf("perms %v: %s should have failed", p, desc)
		}
	}
}
