package slices

import (
	"sort"

	"github.com/aohorodnyk/stl/constraints"
)

// Sort is a helper function that simplifies the API for standard library.
func Sort[T constraints.Ordered](slice []T) {
	sort.Sort(SliceOrdered[T](slice))
}

// SortReverse is a helper function that simplifies the API for standard library.
func SortReverse[T constraints.Ordered](slice []T) {
	sort.Sort(sort.Reverse(SliceOrdered[T](slice)))
}

// Stable is a helper function that simplifies the API for standard library.
func Stable[T constraints.Ordered](slice []T) {
	sort.Stable(SliceOrdered[T](slice))
}

// Stable is a helper function that simplifies the API for standard library.
func StableReverse[T constraints.Ordered](slice []T) {
	sort.Stable(sort.Reverse(SliceOrdered[T](slice)))
}
