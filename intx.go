package intn

import (
	"fmt"
	"math"
)

// ArrayX is integer array with max value
type ArrayX struct {
	MaxVal  uint
	perData uint
	curSize uint64
	Data    []uint64
}

// NewArrayX returns new array struct
func NewArrayX(maxv uint) *ArrayX {
	ret := new(ArrayX)
	ret.MaxVal = maxv
	for v := (uint64)(math.MaxUint64); v != 0; v /= (uint64)(maxv) {
		ret.perData++
	}
	ret.perData--
	ret.Data = []uint64{}
	ret.curSize = 0
	return ret
}

// NewArrayXSized returns new Array with initial size
func NewArrayXSized(maxv uint, size uint64) *ArrayX {
	ret := NewArrayX(maxv)
	pd := uint64(ret.perData)
	ret.Data = make([]uint64, (size+pd-1)/pd)
	ret.curSize = size
	return ret
}

// Append appends value
func (na *ArrayX) Append(val ...uint) (err error) {
	for _, v := range val {
		if na.curSize == (uint64)(len(na.Data))*(uint64)(na.perData) {
			na.Data = append(na.Data, 0)
		}
		if err = na.rawSet(na.curSize, 0, v); err != nil {
			return
		}
		na.curSize++
	}
	return
}

func (na *ArrayX) Extend(ext *ArrayX) (err error) {
	if na.MaxVal < ext.MaxVal {
		return fmt.Errorf("maxval mismatch: %d < %d", na.MaxVal, ext.MaxVal)
	}
	for i := range ext.Each() {
		if err := na.Append(i); err != nil {
			return err
		}
	}
	return nil
}

func (na *ArrayX) Shrink(n uint64) (err error) {
	if na.curSize > n {
		na.curSize = n
		pd := uint64(na.perData)
		na.Data = na.Data[:(n+pd-1)/pd]
		return nil
	}
	return fmt.Errorf("size %d >= %d", n, na.curSize)
}

// Size returns size of array
func (na *ArrayX) Size() uint64 {
	return na.curSize
}

// Capacity returns capacity of array
func (na *ArrayX) Capacity() uint64 {
	return uint64(na.perData) * (uint64)(len(na.Data))
}

// Get returns value
func (na *ArrayX) Get(n uint64) (uint, error) {
	if n > na.curSize {
		return 0, fmt.Errorf("array index out of range: %d/%d", n, na.curSize)
	}
	v := na.Data[n/uint64(na.perData)]
	for i := uint64(0); i < n%uint64(na.perData); i++ {
		v /= uint64(na.MaxVal)
	}
	return (uint)(v % uint64(na.MaxVal)), nil
}

func (na *ArrayX) rawSet(n uint64, old, val uint) (err error) {
	if val >= na.MaxVal {
		return fmt.Errorf("overflow: %d > %d", val, na.MaxVal)
	}
	// mask
	v := na.Data[n/uint64(na.perData)]
	old64 := uint64(old)
	val64 := uint64(val)
	for i := uint64(0); i < n%uint64(na.perData); i++ {
		old64 *= uint64(na.MaxVal)
		val64 *= uint64(na.MaxVal)
	}
	vv := v - old64 + val64
	na.Data[n/uint64(na.perData)] = vv
	return nil
}

// Set sets value
func (na *ArrayX) Set(n uint64, val uint) (prev uint, err error) {
	prev, err = na.Get(n)
	if err != nil {
		return
	}
	err = na.rawSet(n, prev, val)
	return
}

// SetForce extends array when out of range, and sets value
func (na *ArrayX) SetForce(n uint64, val uint) (prev uint, err error) {
	if n >= na.curSize {
		pd := (uint64)(na.perData)
		if n >= (uint64)(len(na.Data))*pd {
			dif := (n - ((uint64)(len(na.Data)) * pd) + pd) / pd
			na.Data = append(na.Data, make([]uint64, dif)...)
		}
		na.curSize = n + 1
	}
	return na.Set(n, val)
}

// Add adds value
func (na *ArrayX) Add(n uint64, val uint) (prev uint, err error) {
	prev, err = na.Get(n)
	if err != nil {
		return
	}
	nextval := prev + val
	err = na.rawSet(n, prev, nextval)
	return
}

// Sub subtracts value
func (na *ArrayX) Sub(n uint64, val uint) (prev uint, err error) {
	prev, err = na.Get(n)
	if err != nil {
		return
	}
	if prev < val {
		return prev, fmt.Errorf("underflow: %d < %d", prev, val)
	}
	nextval := prev - val
	err = na.rawSet(n, prev, nextval)
	return
}

func (na *ArrayX) String() string {
	ret := "["
	na.EachCb(func(idx uint64, val uint) {
		if idx != 0 {
			ret += " "
		}
		ret += fmt.Sprintf("%d", val)
	})
	ret += "]"
	return ret
}

// Each returns channel to get each value
func (na *ArrayX) Each() <-chan uint {
	ch := make(chan uint)
	go func() {
		for i := (uint64)(0); i < na.curSize; i++ {
			if v, err := na.Get(i); err == nil {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}

// EachCb calls callback with each index and value as arguments
func (na *ArrayX) EachCb(cb func(uint64, uint)) {
	for i := (uint64)(0); i < na.curSize; i++ {
		if v, err := na.Get(i); err == nil {
			cb(i, v)
		}
	}
}
