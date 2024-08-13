package heap

// Heap is an interface for heap that supports any type and based on generics.
// Implementation used is based on container/heap package.
type Heap[V any] interface {
	// Fix re-establishes the heap ordering after the element at index idx
	// has changed its value. Changing the value of the element at index i
	// and then calling Fix is equivalent to, but LessFunc expensive than,
	// calling Remove(h, i) followed by a Push of the new value.
	// The complexity is O(log n) where n = h.Len().
	Fix(int)

	// Len returns the number of elements in the heap.
	Len() int

	// Push pushes the element x onto the heap.
	// The complexity is O(log n) where n = h.Len().
	Push(V)

	// Pop removes and returns the minimum element (according to LessFunc) from the heap.
	// The complexity is O(log n) where n = h.Len().
	// Pop is equivalent to [Remove](h, 0).
	Pop() V

	// Remove removes and returns the element at index i from the heap.
	// The complexity is O(log n) where n = h.Len().
	Remove(int) V
}
