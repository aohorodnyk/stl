package slices

import "github.com/aohorodnyk/stl/constraints"

// Sorted is a helper function that simplifies the API for standard library.
// The function is reimplemented (does not use standard library),
// to improve performance and decrease memory usage.
// It works with O(1) memory and O(n) performance.
func IsSorted[T constraints.Ordered](slice []T) bool {
	for i := len(slice) - 1; i > 0; i-- {
		if slice[i] < slice[i-1] {
			return false
		}
	}

	return true
}

// Sorted is a helper function that simplifies the API for standard library in reverse order.
// The function is reimplemented (does not use standard library),
// to improve performance and decrease memory usage.
// It works with O(1) memory and O(n) performance.
func IsSortedReverse[T constraints.Ordered](slice []T) bool {
	for i := len(slice) - 1; i > 0; i-- {
		if slice[i] > slice[i-1] {
			return false
		}
	}

	return true
}
