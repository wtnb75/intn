package intn

import (
	"fmt"
	"math"
)

// ArrayNum is integer array with max value
type ArrayNum struct {
	maxVal  uint64
	perData uint
	curSize uint64
	Data    []uint64
}

// NewArrayNum returns new array struct
func NewArrayNum(maxv uint64) Array {
	if maxv < 2 {
		panic(fmt.Sprintf("maxval too small: %d < 2", maxv))
	}
	ret := new(ArrayNum)
	ret.maxVal = maxv
	for v := (uint64)(math.MaxUint64); v != 0; v /= (uint64)(maxv) {
		ret.perData++
	}
	ret.perData--
	ret.Data = []uint64{}
	ret.curSize = 0
	return &(*ret)
}

func (na *ArrayNum) MaxVal() uint64 {
	return na.maxVal
}

func (na *ArrayNum) Resize(n uint64) {
	if na.curSize > n {
		// shrink
		na.curSize = n
		na.Data = na.Data[:(n+uint64(na.perData)-1)/uint64(na.perData)]
		// clear
		na.curSize = na.Capacity()
		for i := n; i < na.Capacity(); i++ {
			na.Set(i, 0)
		}
		na.curSize = n
	} else if na.curSize < n {
		// extend
		plus := (n+uint64(na.perData)-1)/uint64(na.perData) - uint64(len(na.Data))
		na.Data = append(na.Data, make([]uint64, plus)...)
		na.curSize = n
	}
}

// Size returns size of array
func (na *ArrayNum) Size() uint64 {
	return na.curSize
}

// Capacity returns capacity of array
func (na *ArrayNum) Capacity() uint64 {
	return uint64(na.perData) * (uint64)(len(na.Data))
}

// Get returns value
func (na *ArrayNum) Get(n uint64) uint64 {
	if n > na.curSize {
		panic(fmt.Sprintf("array index out of range: %d/%d", n, na.curSize))
	}
	v := na.Data[n/uint64(na.perData)]
	for i := uint64(0); i < n%uint64(na.perData); i++ {
		v /= uint64(na.maxVal)
	}
	return (uint64)(v % uint64(na.maxVal))
}

func (na *ArrayNum) rawSet(n uint64, old, val uint64) {
	if val >= na.maxVal {
		panic(fmt.Sprintf("overflow: %d > %d", val, na.maxVal))
	}
	// mask
	v := na.Data[n/uint64(na.perData)]
	old64 := uint64(old)
	val64 := uint64(val)
	for i := uint64(0); i < n%uint64(na.perData); i++ {
		old64 *= uint64(na.maxVal)
		val64 *= uint64(na.maxVal)
	}
	vv := v - old64 + val64
	na.Data[n/uint64(na.perData)] = vv
	return
}

// Set sets value
func (na *ArrayNum) Set(n uint64, val uint64) (prev uint64) {
	prev = na.Get(n)
	na.rawSet(n, prev, val)
	return
}

func (na *ArrayNum) String() string {
	return String(na)
}
