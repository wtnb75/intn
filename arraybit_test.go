package intn

import (
	"math/rand"
	"testing"
)

func TestArrayBitSum(t *testing.T) {
	ar := NewArrayBit(8)
	s := uint64(0)
	for i := 0; i < 100000; i++ {
		val := uint64(rand.Intn(int(ar.MaxVal())))
		Push(ar, val)
		s += val
	}
	t.Logf("Sum(ar)=%d", Sum(ar))
	if Sum(ar) != s {
		t.Errorf("sum mismatch: %d vs. %d", s, Sum(ar))
	}
}

func benchSumA(b *testing.B, ar Array) {
	for i := 0; i < b.N; i++ {
		Sum(ar)
	}
}

func benchSumN(b *testing.B, ar Array) {
	for i := 0; i < b.N; i++ {
		var res uint64
		EachCb(ar, func(_, v uint64) {
			res += v
		})
	}
}

func BenchmarkAB_NormalSum(b *testing.B) {
	ar := NewArrayBit(5)
	for i := 0; i < 100000; i++ {
		Push(ar, uint64(rand.Intn(int(ar.MaxVal()))))
	}
	b.ResetTimer()
	benchSumN(b, ar)
}

func BenchmarkAB_OptSum(b *testing.B) {
	ar := NewArrayBit(5)
	for i := 0; i < 100000; i++ {
		Push(ar, uint64(rand.Intn(int(ar.MaxVal()))))
	}
	b.ResetTimer()
	benchSumA(b, ar)
}
