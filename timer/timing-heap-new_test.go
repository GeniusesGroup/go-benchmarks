/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"math/rand"
	"testing"
)

func Benchmark_newTiming(b *testing.B) {
	var timing = newTiming{
		timers: make([]timerBucket, 0, b.N),
	}
	// fills memory to force timers allocate on non order locations that not fit on L1, L2 or L3 CPU cache
	var fillsMemory = make([][]byte, 0, b.N)

	for n := 0; n < b.N; n++ {
		var nt newTimer
		nt.when = rand.Int63()
		timing.AddTimer(&nt)

		nt.when = rand.Int63()

		fillsMemory = append(fillsMemory, make([]byte, 25600))
	}

	for n := 0; n < b.N; n++ {
		var tb timerBucket = timing.timers[n]
		var tw = tb.timer.when
		if tw > tb.when {
			timing.siftUpTimer(n)
		} else {
			timing.siftDownTimer(n)
		}
	}
}
