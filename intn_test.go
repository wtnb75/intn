package intn

import (
	"fmt"
	"testing"
)

func TestInit3(t *testing.T) {
	a := NewArray(3)
	if a.Nbits != 3 {
		t.Errorf("bits")
	}
	if a.perData != 21 {
		t.Errorf("perData is not 21")
	}
	if a.curSize != 0 {
		t.Errorf("curSize is not 0")
	}
	uinttbl := []int{0, 1, 2, 3, -4, -3, -2, -1}
	for i, v := range uinttbl {
		if a.Uint2Int((uint)(i)) != v {
			t.Errorf("uint2int %d %d -> %d", i, v, a.Uint2Int((uint)(i)))
		}
		if a.Int2Uint(v) != (uint)(i) {
			t.Errorf("int2uint %d %d -> %d", i, v, a.Int2Uint(v))
		}
	}
	for i := (uint)(0); i <= 7; i++ {
		if err := a.Append(i); err != nil {
			t.Errorf("append error: %s", err)
		}
	}
	t.Logf("Data[0]=%#x", a.Data)
	// error val
	if err := a.Append(8); err == nil {
		t.Errorf("should be overflow")
	}
	if err := a.Append(10); err == nil {
		t.Errorf("should be overflow")
	}
	if a.Size() != 8 {
		t.Errorf("curSize is not 8")
	}
	if a.Capacity() != 21 {
		t.Errorf("capacity is not 21")
	}
	for i := (uint)(0); i <= 7; i++ {
		if v, err := a.Get((uint64)(i)); err != nil || v != i {
			t.Errorf("invalid value or error: %v, %v", err, v)
		}
	}
	if v, err := a.Get(8); err != nil || v != 0 {
		t.Errorf("invalid value for invalid append: %v, %v", err, v)
	}
	if v, err := a.Get(1024); err == nil {
		t.Errorf("should array index out of bound: got %v", v)
	} else {
		t.Logf("got err=%v, v=%v", err, v)
	}
	if v, err := a.Add(1024, 1); err == nil {
		t.Errorf("should array index out of bound: got %v", v)
	} else {
		t.Logf("got err=%v, v=%v", err, v)
	}
	if v, err := a.Sub(1024, 1); err == nil {
		t.Errorf("should array index out of bound: got %v", v)
	} else {
		t.Logf("got err=%v, v=%v", err, v)
	}
	if v, err := a.Set(1024, 1); err == nil {
		t.Errorf("should array index out of bound: got %v", v)
	} else {
		t.Logf("got err=%v, v=%v", err, v)
	}
	if fmt.Sprintf("%s", a) != "[0 1 2 3 4 5 6 7]" {
		t.Errorf("invalid stringify: %s", a)
	}
	t.Log("string:", a)
	if old, err := a.Add(0, 1); err != nil {
		t.Errorf("add error: %v, %d", err, old)
	} else {
		newval, _ := a.Get(0)
		t.Logf("add: old=%v, new=%v", old, newval)
	}
	if old, err := a.Add(7, 1); err == nil {
		t.Errorf("should overflow: err=%v, old=%v", err, old)
	} else {
		t.Logf("oldVal=%v", old)
	}
	if v, err := a.Get(7); err != nil || v != 7 {
		t.Errorf("a[7] is not 7: err=%v, v=%v", err, v)
	}
	if v, err := a.Get(0); err != nil || v != 1 {
		t.Errorf("not added: err=%v, v=%v", err, v)
	}
	if old, err := a.Sub(0, 1); err != nil {
		t.Errorf("sub error: %v, %d", err, old)
	}
	if v, err := a.Get(0); err != nil || v != 0 {
		t.Errorf("not subed: err=%v, v=%v", err, v)
	}
	if old, err := a.Sub(0, 1); err == nil {
		t.Errorf("should underflow: %v, %d", err, old)
	} else {
		t.Logf("oldVal=%v", old)
	}
	if v, err := a.Get(0); err != nil || v != 0 {
		t.Errorf("changed: err=%v, v=%v", err, v)
	}
	if v, err := a.Set(5, 4); err != nil || v != 5 {
		t.Errorf("set error: err=%v, v=%v", err, v)
	}
	var sum uint
	for v := range a.Each() {
		sum += v
	}
	if sum != 27 {
		t.Errorf("sum is not 27: %d", sum)
	}
	var sumcb uint
	a.EachCb(func(idx uint64, v uint) {
		sumcb += v
	})
	if sumcb != 27 {
		t.Errorf("sumcb is not 27: %d", sumcb)
	}
}

func TestInit5(t *testing.T) {
	a := NewArraySized(5, 8192)
	if a.Size() != 8192 {
		t.Errorf("size is not 8192: %d", a.Size())
	}
	if a.Capacity() != 8196 {
		t.Errorf("capacity is not 8196: %d", a.Capacity())
	}
}

func TestReadme(t *testing.T) {
	a := NewArray(5)
	if err := a.Append(5); err != nil {
		t.Errorf("append 5: %s", err)
	}
	if err := a.Append(4); err != nil {
		t.Errorf("append 4: %s", err)
	}
	if fmt.Sprintln(a) != "[5 4]\n" {
		t.Errorf("invalid output: %s", fmt.Sprintln(a))
	}
	if old, err := a.Add(0, 2); err != nil || old != 5 {
		t.Errorf("add: %s, old=%d", err, old)
	}
	var sum uint
	for v := range a.Each() {
		sum += v
	}
	if sum != 11 {
		t.Errorf("invalid sum")
	}
	if fmt.Sprintln(a, sum) != "[7 4] 11\n" {
		t.Errorf("invalid output: %s", fmt.Sprintln(a, sum))
	}
}

func TestSetForce(t *testing.T) {
	bitarray := NewArray(1)
	bitarray.SetForce(1024, 1)
	bitarray.EachCb(func(idx uint64, v uint) {
		if v != 0 && idx != 1024 {
			t.Errorf("invalid value0: idx=%v, v=%v", idx, v)
		}
		if v == 0 && idx == 1024 {
			t.Errorf("invalid value1: idx=%v, v=%v", idx, v)
		}
	})
}

func TestCopy(t *testing.T) {
	a := NewArray(1)
	a.SetForce(1024, 1)
	b := NewArray(2)
	b.Extend(a)
	if b.Size() != a.Size() {
		t.Errorf("size mismatch: a=%v, b=%v", a.Size(), b.Size())
	}
	if len(b.Data) == len(a.Data) {
		t.Errorf("datasize mismatch: a=%v, b=%v", len(a.Data), len(b.Data))
	}
	if a.Nbits == b.Nbits {
		t.Errorf("bitsize mismatch: a=%v, b=%v", a.Nbits, b.Nbits)
	}
	b.EachCb(func(idx uint64, v uint) {
		if v != 0 && idx != 1024 {
			t.Errorf("invalid value0: idx=%v, v=%v", idx, v)
		}
		if idx == 1024 && v != 1 {
			t.Errorf("invalid value1: idx=%v, v=%v", idx, v)
		}
	})
}

func TestCopyNg(t *testing.T) {
	a := NewArray(1)
	b := NewArray(2)
	b.SetForce(1024, 1)
	if err := a.Extend(b); err == nil {
		t.Errorf("copy large should be ng")
	}
}

func TestCopyFast(t *testing.T) {
	a := NewArray(1)
	a.SetForce(1023, 1)
	b := NewArray(1)
	b.Extend(a)
	b.Extend(a)
	if b.Size() != a.Size()*2 {
		t.Errorf("size mismatch: a=%v, b=%v", a.Size(), b.Size())
	}
	if len(b.Data) != len(a.Data)*2 {
		t.Errorf("datasize mismatch: a=%v, b=%v", len(a.Data), len(b.Data))
	}
	if a.Nbits != b.Nbits {
		t.Errorf("bitsize mismatch: a=%v, b=%v", a.Nbits, b.Nbits)
	}
	if v, err := b.Get(1023); err != nil || v != 1 {
		t.Errorf("data mismatch: err=%v, v=%v", err, v)
	}
	if v, err := b.Get(2047); err != nil || v != 1 {
		t.Errorf("data mismatch: err=%v, v=%v", err, v)
	}
	b.EachCb(func(idx uint64, v uint) {
		if v != 0 && (idx != 1023 && idx != 2047) {
			t.Errorf("invalid value0: idx=%v, v=%v", idx, v)
		}
	})
}

func TestShrink(t *testing.T) {
	a := NewArraySized(1, 1024)
	if err := a.Shrink(16); err != nil {
		t.Errorf("shrink failed: err=%v", err)
	}
	if a.Size() != 16 {
		t.Errorf("size mismatch: %d != 16", a.Size())
	}
	t.Logf("shrinked: %s", a)
	if err := a.Shrink(1024); err == nil {
		t.Errorf("shrink to smaller: should failed, err=nil")
	}
}
