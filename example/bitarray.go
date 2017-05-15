package main

import (
	"fmt"
	"math/rand"

	"github.com/wtnb75/intn"
)

func main() {
	bitarray := intn.NewArrayBit(1)
	for i := 0; i < 1024; i++ {
		intn.Push(bitarray, uint64(rand.Intn(2)))
	}
	fmt.Println("result:", bitarray)
}
