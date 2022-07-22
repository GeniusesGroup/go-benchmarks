/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"sync"
)

// timer in go 1.18
type oldTimer struct {
	when     int64
	period   int64
	f        func(any, uintptr)
	arg      any
	seq      uintptr
	nextWhen int64
	status   uint32
	timing   *oldTiming
}

// timing in go 1.18
type oldTiming struct {
	timersLock sync.Mutex
	timers     []*oldTimer
}

func (th *oldTiming) AddTimer(t *oldTimer) {
	th.timersLock.Lock()

	t.timing = th
	var i = len(th.timers)
	th.timers = append(th.timers, t)

	th.siftUpTimer(i)

	th.timersLock.Unlock()
}

func (th *oldTiming) siftUpTimer(i int) int {
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

func (th *oldTiming) siftDownTimer(i int) {
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
