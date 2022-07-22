/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	_ "unsafe" // for go:linkname
)

// Provided by package runtime.
//go:linkname HardwareNow time.now
func HardwareNow() (sec int64, nsec int32, mono int64)

// RuntimeNano returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timing, ...
//go:linkname RuntimeNano runtime.nanotime
func RuntimeNano() int64
