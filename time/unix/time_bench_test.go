/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	"testing"
	"time"
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz
Benchmark_goUnixMilli-8   	100000000	        10.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_Now-8			   	278594536	        4.172 ns/op	       0 B/op	       0 allocs/op
*/

func Benchmark_goUnixMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Now().Unix()
	}
}

// func Benchmark_NowMilli(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		unix.Now().SecElapsed()
// 	}
// }

func Benchmark_Now(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Now()
	}
}
