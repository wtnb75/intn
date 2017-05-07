package main

import (
	"fmt"
	"math/rand"

	"github.com/wtnb75/intn"
)

func main() {
	bitarray := intn.NewArray(1)
	for i := 0; i < 1024; i++ {
		bitarray.Append((uint)(rand.Intn(2)))
	}
	fmt.Println("result:", bitarray)
}
