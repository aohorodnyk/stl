package slices

import (
	"sort"

	"github.com/aohorodnyk/stl/types"
)

// Sort is a helper function that simplifies the API for standard library.
func Sort[T types.Ordered](slice []T) {
	sort.Sort(SliceOrdered[T](slice))
}

// SortReverse is a helper function that simplifies the API for standard library.
func SortReverse[T types.Ordered](slice []T) {
	sort.Sort(sort.Reverse(SliceOrdered[T](slice)))
}

// Sorted is a helper function that simplifies the API for standard library.
func Sorted[T types.Ordered](slice []T) bool {
	return sort.IsSorted(SliceOrdered[T](slice))
}

// SortedReverse is a helper function that simplifies the API for standard library.
func SortedReverse[T types.Ordered](slice []T) bool {
	return sort.IsSorted(sort.Reverse(SliceOrdered[T](slice)))
}

// Stable is a helper function that simplifies the API for standard library.
func Stable[T types.Ordered](slice []T) {
	sort.Stable(SliceOrdered[T](slice))
}

// Stable is a helper function that simplifies the API for standard library.
func StableReverse[T types.Ordered](slice []T) {
	sort.Stable(sort.Reverse(SliceOrdered[T](slice)))
}
