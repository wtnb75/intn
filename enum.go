//go:generate stringer -type ArrayType enum.go
package intn

type ArrayType int

const (
	ARRAYUINT ArrayType = iota
	ARRAYBIT
	ARRAYNUM
)
