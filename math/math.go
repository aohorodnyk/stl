package math

import (
	"github.com/aohorodnyk/stl/constraints"
)

// Abs is a modular function that always returns positive numbers.
// This function only supports Signed Integer types.
// Unsigned types have no sense, because they cannot be negative.
func Abs[T constraints.Signed](number T) T {
	if number < 0 {
		return -number
	}

	return number
}
