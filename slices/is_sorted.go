package slices

import "github.com/aohorodnyk/stl/constraints"

// Sorted is a helper function that simplifies the API for standard library.
// It works with O(1) memory and O(n) performance.
// The function is reimplemented (does not use standard library),
// to improve performance and decrease memory usage.
func IsSorted[Type constraints.Ordered](slice []Type) bool {
	for i := len(slice) - 1; i > 0; i-- {
		if slice[i] < slice[i-1] {
			return false
		}
	}

	return true
}
