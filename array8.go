package intn

import (
	"math"
)

// Array8 is array of uint8
type Array8 []uint8

func NewArray8() Array {
	ret := new(Array8)
	return &(*ret)
}

func (na *Array8) MaxVal() uint64 {
	return uint64(math.MaxUint8)
}

func (na *Array8) Size() uint64 {
	return uint64(len(*na))
}

func (na *Array8) Capacity() uint64 {
	return uint64(cap(*na))
}

func (na *Array8) Resize(sz uint64) {
	if sz < na.Size() {
		(*na) = (*na)[:sz]
	} else if sz > na.Size() {
		(*na) = append((*na), make([]uint8, sz-na.Size())...)
	}
}

func (na *Array8) Get(idx uint64) uint64 {
	return uint64((*na)[idx])
}

func (na *Array8) Set(idx uint64, val uint64) uint64 {
	ret := (*na)[idx]
	(*na)[idx] = uint8(val)
	return uint64(ret)
}

func (na *Array8) String() string {
	return String(na)
}
