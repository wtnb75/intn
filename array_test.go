package intn

import (
	"bytes"
	"encoding/gob"
	"math"
	"math/rand"
	"reflect"
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

func resizeTest(t *testing.T, a Array) {
	a.Resize(987654)
	if a.Size() != 987654 {
		t.Errorf("resize failed? %d vs. %d", 987654, a.Size())
	}
	if a.Capacity() < 987654 {
		t.Errorf("capacity too small: %d vs. %d", 987654, a.Capacity())
	}
	a.Resize(10)
	if a.Size() != 10 {
		t.Errorf("resize failed? %d vs. %d", 10, a.Size())
	}
	if a.Capacity() < 10 {
		t.Errorf("capacity too small: %d vs. %d", 10, a.Capacity())
	}
}

func eachTest(t *testing.T, fn func(t *testing.T, ar Array)) {
	t.Log("Array8")
	fn(t, NewArray8())
	t.Log("Array16")
	fn(t, NewArray16())
	t.Log("ArrayBit3")
	fn(t, NewArrayBit(3))
	t.Log("ArrayBit14")
	fn(t, NewArrayBit(14))
	t.Log("ArrayBit30")
	fn(t, NewArrayBit(30))
	t.Log("ArrayBit49")
	fn(t, NewArrayBit(49))
	t.Log("ArrayNum5")
	fn(t, NewArrayNum(5))
	t.Log("ArrayNum10")
	fn(t, NewArrayNum(10))
	t.Log("ArrayNum1000")
	fn(t, NewArrayNum(1000))
	t.Log("ArrayNum654321")
	fn(t, NewArrayNum(6543421))
}

func TestGetError(t *testing.T) {
	eachTest(t, getError)
}

func TestSetError(t *testing.T) {
	eachTest(t, setError)
}

func TestResize(t *testing.T) {
	eachTest(t, resizeTest)
}

func checkMax(t *testing.T, ar Array, expected uint64) {
	if ar.MaxVal() != expected {
		t.Errorf("max value of %v is not %d (%d)", reflect.TypeOf(ar), expected, ar.MaxVal())
	}
}

func TestMaxVal(t *testing.T) {
	// uintX
	checkMax(t, NewArray(ARRAYUINT, 8765, 0), 65535)
	checkMax(t, NewArray(ARRAYUINT, 123, 0), 255)
	checkMax(t, NewArray(ARRAYUINT, 255, 0), 255)
	checkMax(t, NewArray(ARRAYUINT, 1, 0), 255)
	// bits
	checkMax(t, NewArray(ARRAYBIT, 122, 0), 127)
	checkMax(t, NewArray(ARRAYBIT, 1, 0), 1)
	checkMax(t, NewArray(ARRAYBIT, 2, 0), 3)
	checkMax(t, NewArray(ARRAYBIT, 8000, 0), 8191)
	checkMax(t, NewArray(ARRAYBIT, 56, 0), 63)
	checkMax(t, NewArray(ARRAYBIT, 32, 0), 63)
	checkMax(t, NewArray(ARRAYBIT, 31, 0), 31)
	checkMax(t, NewArray(ARRAYBIT, 31, 0), 31)
	// num
	checkMax(t, NewArray(ARRAYNUM, 31, 0), 31)
	checkMax(t, NewArray(ARRAYNUM, 3, 0), 3)
	checkMax(t, NewArray(ARRAYNUM, 5, 0), 5)
	checkMax(t, NewArray(ARRAYNUM, 8191, 0), 8191)
	checkMax(t, NewArray(ARRAYNUM, 654321, 0), 654321)
}

func checkSizeof(t *testing.T, ar Array, expected uint64) {
	if ar.Sizeof() != expected {
		t.Errorf("size mismatch: %d != %d", ar.Sizeof(), expected)
	} else {
		t.Logf("size ok: %d = %d, %.2f", ar.Sizeof(), expected, float64(ar.Sizeof())/float64(ar.Size()))
	}
}

func TestSizeof(t *testing.T) {
	checkSizeof(t, NewArray(ARRAYUINT, 10, 20), 20)
	checkSizeof(t, NewArray(ARRAYUINT, 250, 20), 20)
	checkSizeof(t, NewArray(ARRAYUINT, 1024, 20), 40)
	checkSizeof(t, NewArray(ARRAYUINT, 65535, 20), 40)

	checkSizeof(t, NewArray(ARRAYBIT, 3, 20), 56)
	checkSizeof(t, NewArray(ARRAYBIT, 3, 2000), 552)
	checkSizeof(t, NewArray(ARRAYBIT, 18, 5000), 3384)
	checkSizeof(t, NewArray(ARRAYBIT, 4, 10000), 3864)
	checkSizeof(t, NewArray(ARRAYBIT, 1, 20000), 2552)
	checkSizeof(t, NewArray(ARRAYBIT, 7, 20000), 7672)

	checkSizeof(t, NewArray(ARRAYNUM, 3, 20), 56)
	checkSizeof(t, NewArray(ARRAYNUM, 3, 2000), 448)
	checkSizeof(t, NewArray(ARRAYNUM, 18, 5000), 2720)
	checkSizeof(t, NewArray(ARRAYNUM, 4, 10000), 2632)
	checkSizeof(t, NewArray(ARRAYNUM, 2, 20000), 2592)
	checkSizeof(t, NewArray(ARRAYNUM, 7, 20000), 7328)
}

func testString(t *testing.T, ar Array) {
	ar.Resize(0)
	Push(ar, 1)
	Push(ar, 2)
	Push(ar, 3)
	if ar.String() != "[1 2 3]" {
		t.Error("invalid string:", ar)
	}
}

func TestStringVal(t *testing.T) {
	eachTest(t, testString)
}

func testGob(t *testing.T, ar Array, ty ArrayType) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(ar); err != nil {
		t.Error("encode error", err)
	} else {
		t.Log("encode ok")
	}

	dec := gob.NewDecoder(&buf)
	ar2 := NewArray(ty, ar.MaxVal(), ar.Size())
	if err := dec.Decode(ar2); err != nil {
		t.Error("decode error", err)
	} else {
		t.Log("decode ok")
	}

	for i := uint64(0); i < ar.Size(); i++ {
		if ar.Get(i) != ar2.Get(i) {
			t.Error("mismatch", i, ar.Get(i), ar2.Get(i))
		}
	}
}

func TestGob1(t *testing.T) {
	ar := NewArray(ARRAYBIT, 10, 0)
	for i := 0; i < 100000; i++ {
		Push(ar, uint64(rand.Int())%ar.MaxVal())
	}
	testGob(t, ar, ARRAYBIT)
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

func BenchmarkAppend16(b *testing.B) {
	benchAppend(b, NewArray16())
}

func BenchmarkAppendBit8(b *testing.B) {
	benchAppend(b, NewArrayBit(8))
}

func BenchmarkAppendNum8(b *testing.B) {
	benchAppend(b, NewArrayNum(8))
}
