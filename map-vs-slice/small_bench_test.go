/* For license and copyright information please see the LEGAL file in the code repository */

package ms

import (
	"fmt"
	"runtime"
	"testing"
)

/*
Number of CPU used: 8
goos: windows
goarch: amd64
pkg: github.com/GeniusesGroup/go-benchmarks/map-vs-slice
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz
Benchmark_smallMap_Insert-8          	  995197	      1253 ns/op	    1208 B/op	       2 allocs/op
Benchmark_smallSlice_Insert-8        	 3574993	       334.6 ns/op	     512 B/op	       1 allocs/op
Benchmark_smallMap_GetMiddle-8       	  898304	      1267 ns/op	    1208 B/op	       2 allocs/op
Benchmark_smallSlice_GetMiddle-8     	 2907439	       430.9 ns/op	     512 B/op	       1 allocs/op
Benchmark_smallMap_GetNotExist-8     	  908842	      1359 ns/op	    1208 B/op	       2 allocs/op
Benchmark_smallSlice_GetNotExist-8   	 3342152	       354.9 ns/op	     512 B/op	       1 allocs/op
*/

func init() {
	fmt.Println("Number of CPU used:", runtime.NumCPU())
}

/*
	Benchmarks
*/

func Benchmark_smallMap_Insert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var sm smallMap
		sm.Init()
		sm.Fill()
	}
}
func Benchmark_smallSlice_Insert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var ss smallSlice
		ss.Init()
		ss.Fill()
	}
}

func Benchmark_smallMap_GetMiddle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var sm smallMap
		sm.Init()
		sm.Fill()
		sm.GetMiddle()
	}
}
func Benchmark_smallSlice_GetMiddle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var ss smallSlice
		ss.Init()
		ss.Fill()
		ss.GetMiddle()
	}
}

func Benchmark_smallMap_GetNotExist(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var sm smallMap
		sm.Init()
		sm.Fill()
		sm.GetNotExist()
	}
}
func Benchmark_smallSlice_GetNotExist(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var ss smallSlice
		ss.Init()
		ss.Fill()
		ss.GetNotExist()
	}
}
