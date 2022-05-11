package math

import (
	"github.com/aohorodnyk/stl/types"
)

// Abs returns the absolute value of a given number.
func Abs[T types.Signed | types.Float](num T) T {
	if num < 0 {
		num = -num
	}

	return num
}
