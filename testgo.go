package testgo

import (
	"fmt"
	"strings"
)

// Add to integers
func Add(a, b int) int {
	return a + b
}

// Represents a slice of integers
type ArrInt []int

// Add elements of one slice to another slice
func ArrAdd(a, b ArrInt) ArrInt {
	length := len(a)
	if length-len(b) > 0 {
		length = len(b)
	}
	c := make(ArrInt, length)
	for i := 0; i < length; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

// Convert a slice into string
func (a ArrInt) String() string {
	out := make([]string, len(a))

	for i, v := range a {
		out[i] = fmt.Sprintf(`<%d>`, v)
	}
	return fmt.Sprintf(`[%s]`, strings.Join(out, ` `))
}
