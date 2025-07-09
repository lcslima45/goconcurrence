package main

import (
	"fmt"
	"runtime"
)

func main() {
	largeSlice := make([]int, 10000000)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	runtime.GC()
	var memStats runtime.MemStats

	runtime.ReadMemStats(&memStats)

	fmt.Printf("total allocated memory in bytes %d\n", memStats.TotalAlloc)
	fmt.Printf("total heap memory %d\n", memStats.HeapAlloc)
	fmt.Printf("number of garbage collections %d\n", memStats.NumGC)
}
