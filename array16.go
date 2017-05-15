package intn

import (
	"math"
)

// Array16 is array of uint16
type Array16 []uint16

func NewArray16() Array {
	ret := new(Array16)
	return &(*ret)
}

func (na *Array16) MaxVal() uint64 {
	return uint64(math.MaxUint16)
}

func (na *Array16) Size() uint64 {
	return uint64(len(*na))
}

func (na *Array16) Capacity() uint64 {
	return uint64(cap(*na))
}

func (na *Array16) Resize(sz uint64) {
	if sz < na.Size() {
		(*na) = (*na)[:sz]
	} else if sz > na.Size() {
		(*na) = append((*na), make([]uint16, sz-na.Size())...)
	}
}

func (na *Array16) Get(idx uint64) uint64 {
	return uint64((*na)[idx])
}

func (na *Array16) Set(idx uint64, val uint64) uint64 {
	ret := (*na)[idx]
	(*na)[idx] = uint16(val)
	return uint64(ret)
}

func (na *Array16) String() string {
	return String(na)
}
