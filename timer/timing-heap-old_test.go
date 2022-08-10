/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"math/rand"
	"testing"
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz

With fillsMemory:
Benchmark_newTiming-8   	  463105	      2778 ns/op	   27368 B/op	       2 allocs/op
Benchmark_oldTiming-8   	  248739	     24886 ns/op	   27376 B/op	       2 allocs/op

Without fillsMemory:
Benchmark_newTiming-8   	 7114608	       178.2 ns/op	      80 B/op	       1 allocs/op
Benchmark_oldTiming-8   	 5867421	       195.9 ns/op	      88 B/op	       1 allocs/op
*/

func Benchmark_oldTiming(b *testing.B) {
	var timing = oldTiming{
		timers: make([]*oldTimer, 0, b.N),
	}
	// fills memory to force timers allocate on non order locations that not fit on L1, L2 or L3 CPU cache
	var fillsMemory = make([][]byte, 0, b.N)

	for n := 0; n < b.N; n++ {
		var ot oldTimer
		ot.when = rand.Int63()
		timing.AddTimer(&ot)

		ot.nextWhen = rand.Int63()

		fillsMemory = append(fillsMemory, make([]byte, 25600))
	}

	for n := 0; n < b.N; n++ {
		var ot *oldTimer = timing.timers[n]
		if ot.nextWhen > ot.when {
			timing.siftUpTimer(n)
		} else {
			timing.siftDownTimer(n)
		}
	}
}
