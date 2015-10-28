// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vom

import "v.io/v23/vdl"

// hasChunkLen returns true iff the type t is encoded with a top-level
// chunk length.
func hasChunkLen(t *vdl.Type) bool {
	if t.IsBytes() {
		// TODO(bprosnitz) This should probably be chunked
		return false
	}
	switch t.Kind() {
	case vdl.Complex64, vdl.Complex128, vdl.Array, vdl.List, vdl.Set, vdl.Map, vdl.Struct, vdl.Any, vdl.Union, vdl.Optional:
		return true
	}
	return false
}

// isAllowedVersion returns true if the VOM version specified is supported.
func isAllowedVersion(version Version) bool {
	return version >= 0x80 && version <= 0x81
}