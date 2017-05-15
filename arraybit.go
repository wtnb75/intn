package intn

import (
	"fmt"
)

// ArrayBit is n-bit integer array
type ArrayBit struct {
	Nbits   uint
	perData uint
	curSize uint64
	Data    []uint64
}

// NewArray returns new array struct
func NewArrayBit(nBits uint) Array {
	if nBits < 1 {
		panic(fmt.Sprintf("nBits too small: %d < 1", nBits))
	}
	if nBits > 64 {
		panic(fmt.Sprintf("nBits too large: %d > 64", nBits))
	}
	ret := new(ArrayBit)
	ret.Nbits = nBits
	ret.perData = 64 / nBits
	ret.Data = []uint64{}
	ret.curSize = 0
	return &(*ret)
}

// NewArrayBitSized returns new Array with initial size
func NewArrayBitSized(nBits uint, size uint64) Array {
	ret := NewArrayBit(nBits).(*ArrayBit)
	ret.Data = make([]uint64, (size+uint64(ret.perData)-1)/uint64(ret.perData))
	ret.curSize = size
	return &(*ret)
}

func (na *ArrayBit) MaxVal() uint64 {
	return (uint64(1) << na.Nbits) - 1
}

func (na *ArrayBit) Resize(n uint64) {
	if na.curSize > n {
		// shrink
		na.curSize = n
		na.Data = na.Data[:(n+uint64(na.perData)-1)/uint64(na.perData)]
		// clear
		for i := na.Size(); i < na.Capacity(); i++ {
			na.Set(i, 0)
		}
	} else if na.curSize < n {
		// extend
		plus := (n+uint64(na.perData)-1)/uint64(na.perData) - uint64(len(na.Data))
		na.Data = append(na.Data, make([]uint64, plus)...)
		na.curSize = n
	}
}

// Size returns size of array
func (na *ArrayBit) Size() uint64 {
	return na.curSize
}

// Capacity returns capacity of array
func (na *ArrayBit) Capacity() uint64 {
	return uint64(na.perData) * (uint64)(len(na.Data))
}

// Get returns value
func (na *ArrayBit) Get(n uint64) uint64 {
	if n > na.curSize {
		panic(fmt.Sprintf("array index out of range: %d/%d", n, na.curSize))
	}
	v := na.Data[n/uint64(na.perData)]
	v >>= (n % uint64(na.perData)) * (uint64)(na.Nbits)
	return v & ((1 << na.Nbits) - 1)
}

func (na *ArrayBit) rawSet(n uint64, val uint64) {
	if val >= (1 << na.Nbits) {
		panic(fmt.Sprintf("overflow: %d > %d", val, 1<<na.Nbits))
	}
	// mask
	mask := (uint64(1) << na.Nbits) - 1
	mask <<= (n % uint64(na.perData)) * (uint64)(na.Nbits)
	v := (uint64)(val)
	v <<= (n % uint64(na.perData)) * (uint64)(na.Nbits)
	// clear
	na.Data[n/uint64(na.perData)] &^= mask
	// and set
	na.Data[n/uint64(na.perData)] |= v
	return
}

// Set sets value
func (na *ArrayBit) Set(n uint64, val uint64) (prev uint64) {
	prev = na.Get(n)
	na.rawSet(n, val)
	return
}

// Uint2Int returns Two's complement
func (na *ArrayBit) Uint2Int(val uint64) int64 {
	if val < (1 << (na.Nbits - 1)) {
		return (int64)(val)
	}
	// 0<->1
	val ^= (1 << na.Nbits) - 1
	// +1
	val++
	return -(int64)(val)
}

// Int2Uint returns Two's complement
func (na *ArrayBit) Int2Uint(val int64) uint64 {
	return (uint64)(val & ((1 << na.Nbits) - 1))
}

func (na *ArrayBit) String() string {
	return String(na)
}
