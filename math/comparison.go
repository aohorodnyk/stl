package math

import (
	"github.com/aohorodnyk/stl/constraints"
)

// Min function returns the smaller of two values for all types that support strict order comparisons.
// Float values are not excluded, but keep in mind they can contain INF and NAN values that require additional handling.
// This function is variadic to make it more flexible. At the same time variadic functin is still faster
// then a function from std lib, because it does not support types with INFs and NANs.
func Min[T constraints.Ordered](values ...T) T {
	if len(values) == 0 {
		// Panic is an antipattern, but it's unclear what to return when values is empty.
		panic("Min function requires minimum one parameter")
	}

	result := values[0]
	for index := 1; index < len(values); index++ {
		if result > values[index] {
			result = values[index]
		}
	}

	return result
}

// Max function returns the larger of two values for all types that support strict order comparisons.
// Float values are excluded, but keep in mind they can contain INF and NAN values that require additional handling.
// This function is variadic to make it more flexible. At the same time variadic functin is still faster
// then a function from std lib, because it does not support types with INFs and NANs.
func Max[T constraints.Ordered](values ...T) T {
	if len(values) == 0 {
		// Panic is an antipattern, but it's unclear what to return when values is empty.
		panic("Max function requires minimum one parameter")
	}

	result := values[0]
	for index := 1; index < len(values); index++ {
		if result < values[index] {
			result = values[index]
		}
	}

	return result
}
