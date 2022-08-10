/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	_ "unsafe" // for go:linkname
)

// Now returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timing, ...
//
//go:linkname Now runtime.nanotime
func Now() int64
