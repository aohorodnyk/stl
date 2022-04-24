package sort

import (
	"sort"

	"github.com/aohorodnyk/stl/types"
)

// Ordered is a helper function that sorts a slice of elements of Ordered type in ascending order.
func Ordered[T types.Ordered](slice []T) {
	sort.Sort(SliceOrdered[T](slice))
}

// OrderedReverse is a helper function that sorts a slice of elements of Ordered type in reverse order.
func OrderedReverse[T types.Ordered](slice []T) {
	sort.Sort(sort.Reverse(SliceOrdered[T](slice)))
}

// LessFunc is a helper function that sorts a slice of elements of any type use custom less function.
// It is useful for sorting slices of custom types.
// It is better implementation than sort.Slice for custom types, because it does not use reflection.
func LessFunc[T any](slice []T, less func(int, int) bool) {
	sortSlice := &SliceLessFunc[T]{
		Data:     slice,
		LessFunc: less,
	}

	sort.Sort(sortSlice)
}
