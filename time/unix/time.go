/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	_ "unsafe" // for go:linkname
)

// Provided by package runtime.
//
//go:linkname Now time.now
func Now() (sec int64, nsec int32, mono int64)
