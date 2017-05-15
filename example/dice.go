package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"

	"github.com/wtnb75/intn"
)

func main() {
	var n = flag.Int("num", 10000, "try count")
	flag.Parse()
	data := intn.NewArrayBitSized(3, (uint64)(*n))
	for i := (uint64)(0); i < (uint64)(*n); i++ {
		v := rand.Intn(6) + 1
		data.Set(i, uint64(v))
	}
	var count [6]uint
	var relation [6][6]uint
	var prev uint = 99
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	for i := range intn.Each(data) {
		if prev != 99 {
			relation[prev][i-1]++
		}
		count[i-1]++
		prev = uint(i - 1)
	}
	fmt.Println("data size: ", data.Size(), "capacity:", data.Capacity())
	fmt.Println("count: ", count)
	fmt.Println("relation: ", relation)
	fmt.Println("memory alloc", mem.Alloc, "total", mem.TotalAlloc, "heap", mem.HeapAlloc, "sys", mem.HeapSys)
}
