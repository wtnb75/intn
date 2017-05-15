package intn

import (
	"math"
	"testing"
)

func getError(t *testing.T, a Array) {
	Push(a, 1)
	Push(a, 2)
	if val, err := GetWithError(a, 1); err != nil {
		t.Errorf("error 1 val=%v, err=%v", val, err)
	} else {
		t.Logf("ok 1 val=%v, err=%v", val, err)
	}
	if val, err := GetWithError(a, 10); err != nil {
		t.Logf("error 10 val=%v, err=%v", val, err)
	} else {
		t.Errorf("ok 10 val=%v, err=%v", val, err)
	}
}

func setError(t *testing.T, a Array) {
	Push(a, 1)
	Push(a, 2)
	if val, err := SetWithError(a, 1, 4); err != nil {
		t.Errorf("error 1 val=%v, err=%v", val, err)
	} else {
		t.Logf("ok 1 val=%v, err=%v", val, err)
	}
	if val, err := SetWithError(a, 10, 4); err != nil {
		t.Logf("error 10 val=%v, err=%v", val, err)
	} else {
		t.Errorf("ok 10 val=%v, err=%v", val, err)
	}
	if val, err := SetForce(a, 20, 4); err != nil {
		t.Errorf("error 20 val=%v, err=%v", val, err)
	} else {
		t.Logf("ok 20 val=%v, err=%v", val, err)
	}
}

func TestGetError(t *testing.T) {
	t.Log("Array8")
	getError(t, NewArray8())
	t.Log("ArrayBit")
	getError(t, NewArrayBit(3))
	t.Log("ArrayNum")
	getError(t, NewArrayNum(5))
}

func TestSetError(t *testing.T) {
	t.Log("Array8")
	setError(t, NewArray8())
	t.Log("ArrayBit")
	setError(t, NewArrayBit(3))
	t.Log("ArrayNum")
	setError(t, NewArrayNum(5))
}

func benchAppend(b *testing.B, a Array) {
	var maxval = int(a.MaxVal())
	for i := 0; i < b.N; i++ {
		Push(a, uint64(i%maxval))
	}
}

func BenchmarkAppendRaw8(b *testing.B) {
	var maxval = math.MaxUint8
	data := []uint8{}
	for i := 0; i < b.N; i++ {
		data = append(data, uint8(i%maxval))
	}
}

func BenchmarkAppend8(b *testing.B) {
	benchAppend(b, NewArray8())
}

func BenchmarkAppend8_noif(b *testing.B) {
	ar := NewArray8().(*Array8)
	var maxval = int(ar.MaxVal())
	for i := 0; i < b.N; i++ {
		Push(ar, uint64(i%maxval))
	}
}

func BenchmarkAppendBit8(b *testing.B) {
	benchAppend(b, NewArrayBit(8))
}

func BenchmarkAppendNum8(b *testing.B) {
	benchAppend(b, NewArrayNum(8))
}
