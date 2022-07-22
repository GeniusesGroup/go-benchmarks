/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"sync"
)

type newTimer struct {
	when   int64
	period int64
	f      func(any, uintptr)
	arg    any
	seq    uintptr
	status uint32
	timing *newTiming
}

type newTiming struct {
	timersLock sync.Mutex
	timers     []timerBucket
}

type timerBucket struct {
	timer *newTimer
	// Two reason to have timer when here:
	// - hot cache to prevent dereference timer to get when field
	// - It can be difference with timer when filed in timerModifiedXX status.
	when int64
}

func (th *newTiming) AddTimer(t *newTimer) {
	th.timersLock.Lock()

	var timerWhen = t.when
	t.timing = th
	var i = len(th.timers)
	th.timers = append(th.timers, timerBucket{t, timerWhen})

	th.siftUpTimer(i)

	th.timersLock.Unlock()
}

func (th *newTiming) siftUpTimer(i int) int {
	var timers = th.timers
	var timerWhen = timers[i].when

	var tmp = timers[i]
	for i > 0 {
		var p = (i - 1) / 4 // parent
		if timerWhen >= timers[p].when {
			break
		}
		timers[i] = timers[p]
		i = p
	}
	if tmp != timers[i] {
		timers[i] = tmp
	}
	return i
}

func (th *newTiming) siftDownTimer(i int) {
	var timers = th.timers
	var timersLen = len(timers)
	var timerWhen = timers[i].when

	var tmp = timers[i]
	for {
		var c = i*4 + 1 // left child
		var c3 = c + 2  // mid child
		if c >= timersLen {
			break
		}
		var w = timers[c].when
		if c+1 < timersLen && timers[c+1].when < w {
			w = timers[c+1].when
			c++
		}
		if c3 < timersLen {
			var w3 = timers[c3].when
			if c3+1 < timersLen && timers[c3+1].when < w3 {
				w3 = timers[c3+1].when
				c3++
			}
			if w3 < w {
				w = w3
				c = c3
			}
		}
		if w >= timerWhen {
			break
		}
		timers[i] = timers[c]
		i = c
	}
	if tmp != timers[i] {
		timers[i] = tmp
	}
}
