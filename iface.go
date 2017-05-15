package intn

import (
	"fmt"
	"math"
)

// Array is interface of intn's integer array
type Array interface {
	fmt.Stringer
	MaxVal() uint64
	Size() uint64
	Capacity() uint64
	Resize(s uint64)
	Get(idx uint64) uint64
	Set(idx uint64, val uint64) uint64
}

type ArrayType int

const (
	ARRAYUINT ArrayType = iota
	ARRAYBIT
	ARRAYNUM
)

func NewArray(t ArrayType, maxval uint64, size uint64) (ret Array) {
	switch t {
	case ARRAYUINT:
		if maxval < math.MaxUint8 {
			ret = NewArray8()
		} else if maxval < math.MaxUint16 {
			// ret = NewArray16()
		} else if maxval < math.MaxUint32 {
			// ret = NewArray32()
		} else {
			// ret = NewArray64()
		}
		break
	case ARRAYBIT:
		var bits uint
		for bits = 0; (1 << bits) < maxval; bits++ {
			// pass
		}
		ret = NewArrayBit(bits)
		break
	case ARRAYNUM:
		ret = NewArrayNum(maxval)
		break
	}
	ret.Resize(size)
	return ret
}

func String(ar Array) string {
	ret := "["
	EachCb(ar, func(idx uint64, val uint64) {
		if idx != 0 {
			ret += " "
		}
		ret += fmt.Sprintf("%d", val)
	})
	ret += "]"
	return ret
}

// Pop get last value and shrink
func Pop(ar Array) uint64 {
	ret := ar.Get(ar.Size() - 1)
	ar.Resize(ar.Size() - 1)
	return ret
}

// Push appends value
func Push(ar Array, v uint64) {
	sz := ar.Size()
	ar.Resize(sz + 1)
	ar.Set(sz, v)
}

func Extend(ar1 Array, ar2 Array) {
	for v := range Each(ar2) {
		Push(ar1, v)
	}
}

func GetWithError(ar Array, idx uint64) (val uint64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	val = ar.Get(idx)
	return
}

func SetWithError(ar Array, idx uint64, val uint64) (prev uint64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	prev = ar.Set(idx, val)
	return
}

func SetForce(ar Array, idx uint64, val uint64) (prev uint64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	if idx > ar.Size() {
		ar.Resize(idx + 1)
	}
	prev = ar.Set(idx, val)
	return
}

// Each returns channel to get each value
func Each(ar Array) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		for i := (uint64)(0); i < ar.Size(); i++ {
			ch <- ar.Get(i)
		}
		close(ch)
	}()
	return ch
}

// EachCb calls callback with each index and value as arguments
func EachCb(ar Array, cb func(uint64, uint64)) {
	for i := (uint64)(0); i < ar.Size(); i++ {
		cb(i, ar.Get(i))
	}
}

// Add adds value
func Add(ar Array, n uint64, val uint64) (prev uint64) {
	prev = ar.Get(n)
	nextval := prev + val
	ar.Set(n, nextval)
	return
}

// Sub subtract value
func Sub(ar Array, n uint64, val uint64) (prev uint64) {
	prev = ar.Get(n)
	nextval := prev - val
	ar.Set(n, nextval)
	return
}
