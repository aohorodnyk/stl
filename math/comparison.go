package math

import (
	"github.com/aohorodnyk/stl/constraints"
)

// Min function returns the smaller of two values for all types that support strict order comparisons.
// Float values are not excluded, but keep in mind they can contain INF and NAN values that require additional handling.
// This function is variadic to make it more flexible. At the same time variadic functin is still faster
// then a function from std lib, because it does not support types with INFs and NANs.
// If array with zero size provided, default value for the type will be returned.
func Min[Type constraints.Ordered](values ...Type) Type {
	if len(values) == 0 {
		var result Type

		return result
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
// If array with zero size provided, default value for the type will be returned.
func Max[Type constraints.Ordered](values ...Type) Type {
	if len(values) == 0 {
		var result Type

		return result
	}

	result := values[0]
	for index := 1; index < len(values); index++ {
		if result < values[index] {
			result = values[index]
		}
	}

	return result
}
