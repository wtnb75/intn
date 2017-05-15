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

func TestArrayBitErrorMin(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("error recover:", err)
		}
	}()
	ar := NewArrayBit(0)
	t.Error("0-bit array", ar)
}

func TestArrayBitOkMin(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("error recover:", err)
		}
	}()
	ar := NewArrayBit(1)
	t.Log("1-bit array", ar)
}

func TestArrayBitErrorMax(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("error recover:", err)
		}
	}()
	ar := NewArrayBit(65)
	t.Error("65-bit array", ar)
}

func TestArrayBitOkMax(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("error recover:", err)
		}
	}()
	ar := NewArrayBit(64)
	t.Log("64-bit array", ar)
}

func TestU2I(t *testing.T) {
	ar := NewArrayBit(10).(*ArrayBit)
	var intval = []int64{
		-500, 500,
	}
	var uintval = []uint64{
		524, 500,
	}
	for i, v := range intval {
		if ar.Int2Uint(v) != uintval[i] {
			t.Errorf("int2uint: %d vs. %d", uintval[i], ar.Int2Uint(v))
		}
		if ar.Uint2Int(uintval[i]) != v {
			t.Errorf("uint2int: %d vs. %d", v, ar.Uint2Int(uintval[i]))
		}
	}
	if ar.Int2Uint(-500) != 524 {
		t.Errorf("int2uint: %d vs. %d", 2044, ar.Int2Uint(-500))
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
