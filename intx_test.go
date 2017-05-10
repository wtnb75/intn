package intn

import (
	"fmt"
	"testing"
)

func TestInitX3(t *testing.T) {
	a := NewArrayX(3)
	if a.MaxVal != 3 {
		t.Errorf("bits")
	}
	if a.perData != 40 {
		t.Errorf("perData is not 40: %d", a.perData)
	}
	if a.curSize != 0 {
		t.Errorf("curSize is not 0")
	}
	for i := (uint)(0); i < 3; i++ {
		if err := a.Append(i); err != nil {
			t.Errorf("append error: %s", err)
		}
	}
	t.Logf("Data[0]=%#x", a.Data)
	// error val
	if err := a.Append(3); err == nil {
		t.Errorf("should be overflow")
	}
	if err := a.Append(10); err == nil {
		t.Errorf("should be overflow")
	}
	if a.Size() != 3 {
		t.Errorf("curSize is not 3")
	}
	if a.Capacity() != 40 {
		t.Errorf("capacity is not 40")
	}
	for i := (uint)(0); i < 3; i++ {
		if v, err := a.Get((uint64)(i)); err != nil || v != i {
			t.Errorf("invalid value or error: %v, %v", err, v)
		}
	}
	if v, err := a.Get(3); err != nil || v != 0 {
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
	a.Append(0)
	a.Append(1)
	a.Append(2)

	if fmt.Sprintf("%s", a) != "[0 1 2 0 1 2]" {
		t.Errorf("invalid stringify: %s", a)
	}
	t.Log("string:", a)
	if old, err := a.Add(0, 1); err != nil {
		t.Errorf("add error: %v, %d", err, old)
	} else {
		newval, _ := a.Get(0)
		t.Logf("add: old=%v, new=%v", old, newval)
	}
	if old, err := a.Add(2, 1); err == nil {
		t.Errorf("should overflow: err=%v, old=%v", err, old)
	} else {
		t.Logf("oldVal=%v", old)
	}
	if v, err := a.Get(2); err != nil || v != 2 {
		t.Errorf("a[2] is not 2: err=%v, v=%v", err, v)
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
	if v, err := a.Set(5, 1); err != nil || v != 2 {
		t.Errorf("set error: err=%v, v=%v", err, v)
	}
	var sum uint
	for v := range a.Each() {
		sum += v
	}
	if sum != 5 {
		t.Errorf("sum is not 5: %d", sum)
	}
	var sumcb uint
	a.EachCb(func(idx uint64, v uint) {
		sumcb += v
	})
	if sumcb != 5 {
		t.Errorf("sumcb is not 5: %d", sumcb)
	}
}

func TestInitX5(t *testing.T) {
	a := NewArrayXSized(5, 8192)
	if a.Size() != 8192 {
		t.Errorf("size is not 8192: %d", a.Size())
	}
	if a.Capacity() != 8208 {
		t.Errorf("capacity is not 8208: %d", a.Capacity())
	}
}

func TestXReadme(t *testing.T) {
	a := NewArrayX(9)
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

func TestSetXForce(t *testing.T) {
	bitarray := NewArrayX(2)
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
