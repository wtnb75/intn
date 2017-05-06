package intn

import (
	"fmt"
)

// Array is n-bit integer array
type Array struct {
	Nbits   uint
	perData uint64
	curSize uint64
	Data    []uint64
}

// NewArray returns new array struct
func NewArray(nBits uint) *Array {
	ret := new(Array)
	ret.Nbits = nBits
	ret.perData = (uint64)(64 / nBits)
	ret.Data = []uint64{}
	ret.curSize = 0
	return ret
}

// NewArraySized returns new Array with initial size
func NewArraySized(nBits uint, size uint64) *Array {
	ret := new(Array)
	ret.Nbits = nBits
	ret.perData = uint64(64 / nBits)
	ret.Data = make([]uint64, (size+ret.perData-1)/ret.perData)
	ret.curSize = size
	return ret
}

// Append appends value
func (na *Array) Append(val ...uint) (err error) {
	for _, v := range val {
		if (na.curSize)%(uint64)(na.perData) == 0 {
			na.Data = append(na.Data, 0)
		}
		if err = na.rawSet(na.curSize, v); err != nil {
			return
		}
		na.curSize++
	}
	return
}

// Size returns size of array
func (na *Array) Size() uint64 {
	return na.curSize
}

// Capacity returns capacity of array
func (na *Array) Capacity() uint64 {
	return na.perData * (uint64)(len(na.Data))
}

// Get returns value
func (na *Array) Get(n uint64) (uint, error) {
	if n > na.Capacity() {
		return 0, fmt.Errorf("array index out of range: %d/%d", n, na.Capacity())
	}
	v := na.Data[n/na.perData]
	v >>= (n % na.perData) * (uint64)(na.Nbits)
	return (uint)(v) & ((1 << na.Nbits) - 1), nil
}

func (na *Array) rawSet(n uint64, val uint) (err error) {
	if val >= (1 << na.Nbits) {
		return fmt.Errorf("overflow: %d > %d", val, 1<<na.Nbits)
	}
	// mask
	mask := ((uint64)(1) << na.Nbits) - 1
	mask <<= (n % na.perData) * (uint64)(na.Nbits)
	v := (uint64)(val)
	v <<= (n % na.perData) * (uint64)(na.Nbits)
	// clear
	na.Data[n/na.perData] &^= mask
	// and set
	na.Data[n/na.perData] |= v
	return nil
}

// Set sets value
func (na *Array) Set(n uint64, val uint) (prev uint, err error) {
	prev, err = na.Get(n)
	if err != nil {
		return
	}
	err = na.rawSet(n, val)
	return
}

// Add adds value
func (na *Array) Add(n uint64, val uint) (prev uint, err error) {
	prev, err = na.Get(n)
	if err != nil {
		return
	}
	nextval := prev + val
	err = na.rawSet(n, nextval)
	return
}

// Sub subtracts value
func (na *Array) Sub(n uint64, val uint) (prev uint, err error) {
	prev, err = na.Get(n)
	if err != nil {
		return
	}
	if prev < val {
		return prev, fmt.Errorf("underflow: %d < %d", prev, val)
	}
	nextval := prev - val
	err = na.rawSet(n, nextval)
	return
}

// Uint2Int returns Two's complement
func (na *Array) Uint2Int(val uint) int {
	if val < (1 << (na.Nbits - 1)) {
		return (int)(val)
	}
	// 0<->1
	val ^= (1 << na.Nbits) - 1
	// +1
	val++
	return -(int)(val)
}

// Int2Uint returns Two's complement
func (na *Array) Int2Uint(val int) uint {
	return (uint)(val & ((1 << na.Nbits) - 1))
}

func (na *Array) String() string {
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
func (na *Array) Each() <-chan uint {
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
func (na *Array) EachCb(cb func(uint64, uint)) {
	for i := (uint64)(0); i < na.curSize; i++ {
		if v, err := na.Get(i); err == nil {
			cb(i, v)
		}
	}
}
