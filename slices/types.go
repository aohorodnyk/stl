package slices

import (
	"github.com/aohorodnyk/stl/constraints"
)

// SliceOrdered is a sorting type that implements the sort.Interface.
// It simplifies interface, because of works for type T with all ordered types,
// instead of using different sort functions.
// This type can be used with all standard library functions, because of it implements sort.Interface.
type SliceOrdered[T constraints.Ordered] []T

// Len returns the length of the slice.
func (s SliceOrdered[T]) Len() int {
	return len(s)
}

// Less returns true if the element with index i should sort before the element with index j.
func (s SliceOrdered[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap swaps the elements with indexes i and j.
func (s SliceOrdered[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
