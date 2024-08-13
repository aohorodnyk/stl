package heap

import "github.com/aohorodnyk/stl/types"

// NewStdInterface creates new StdInterface with given data.
func NewStdInterface[V any](data []V) StdInterface[V] {
	return StdInterface[V]{
		data: data,
	}
}

// StdInterface is a heap implementation for generics with [V any]
// but for standard interface of `container/heap` package.
// This implementation is completely public and not thread-safe.
// It's exposed to allow users to manipulate with stardard interface and data.
// Implements heap.Interface and sort.Interface.
type StdInterface[V any] struct {
	// data is a slice of elements.
	// data can be modified directly and grow or shrink automatically.
	// In sort functions it's used as a reference to the data slice. Use StdInterface.Value(index) to get the value.
	data []V
	// LessFunc is a function that defines the sort order.
	// It is used to compare elements by their values based on their indexes in the data slice.
	// Make sure do not store a reference to the data slice, because it can be modified. Use StdInterface.data instead.
	// For example:
	// 	h := StdInterface{
	//    data: []int{1, 2, 3, 4, 5},
	//    MaxSize: math.MaxInt, // or any other value
	//	}
	// 	h.LessFunc = func(i, j int) bool {
	// 		return cmp.Less(h.data[i], h.data[j])
	// 	}
	// This function is required.
	LessFunc func(int, int) bool
}

func (h *StdInterface[V]) Value(index int) V {
	return h.data[index]
}

// Len returns the number of elements in the heap.
func (h *StdInterface[V]) Len() int {
	return len(h.data)
}

// Less checks if the element with index i is less than the element with index j.
func (h *StdInterface[V]) Less(i, j int) bool {
	return h.LessFunc(i, j)
}

// Swap swaps the elements with indexes i and j.
func (h *StdInterface[V]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push pushes the element x onto the heap.
func (h *StdInterface[V]) Push(value any) {
	converted, _ := types.Cast[V](value)
	h.data = append(h.data, converted)
}

// Pop pops the last element from the heap.
func (h *StdInterface[V]) Pop() any {
	if h.Len() == 0 {
		return nil
	}

	val := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]

	return val
}
