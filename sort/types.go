package sort

import (
	"github.com/aohorodnyk/stl/types"
)

// SliceOrdered is a sorting type that implements the sort.Interface.
// It simplifies interface, because of works for type T with all ordered types,
// instead of using different sort functions.
// This type can be used with all standard library functions, because of it implements sort.Interface.
type SliceOrdered[T types.Ordered] []T

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

// SliceLessFunc is a helper type that implements the sort.Interface.
// It improves performance of sort.Sort, because of removed reflection internals.
// This type can be used with all standard library functions, because of it implements sort.Interface.
type SliceLessFunc[T any] struct {
	Data     []T
	LessFunc func(int, int) bool
}

// Len returns the length of the slice.
func (s *SliceLessFunc[T]) Len() int {
	return len(s.Data)
}

// Less returns true if the element with index i should sort before the element with index j.
func (s *SliceLessFunc[T]) Less(i, j int) bool {
	return s.LessFunc(i, j)
}

// Swap swaps the elements with indexes i and j.
func (s *SliceLessFunc[T]) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}
