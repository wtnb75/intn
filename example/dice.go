package main

import (
	"flag"
	"log"
	"math/rand"
	"runtime"

	"github.com/wtnb75/intn"
)

func main() {
	var n = flag.Int("num", 10000, "try count")
	var atyp = flag.String("type", "ARRAYNUM", "array type")
	flag.Parse()
	atypx := intn.ARRAYBIT
	switch *atyp {
	case "num":
		atypx = intn.ARRAYNUM
		break
	case "bit":
		atypx = intn.ARRAYBIT
		break
	case "uint":
		atypx = intn.ARRAYUINT
		break
	default:
		log.Println("no such type: %s ... use default", atyp)
	}
	log.Println("atyp:", atypx)
	data := intn.NewArray(atypx, 7, (uint64)(*n))
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
	log.Println("sizeof: ", data.Sizeof())
	log.Println("data size: ", data.Size(), "capacity:", data.Capacity())
	log.Println("count: ", count)
	log.Println("relation: ", relation)
	log.Println("memory alloc", mem.Alloc, "total", mem.TotalAlloc, "heap", mem.HeapAlloc, "sys", mem.HeapSys)
}
