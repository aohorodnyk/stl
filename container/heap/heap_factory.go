package heap

import (
	"cmp"
	heapStd "container/heap"
	"sync"

	"github.com/aohorodnyk/stl/constraints"
	syncStl "github.com/aohorodnyk/stl/sync"
)

// NewMaxHeapOrdered creates new max heap without limits.
// It supports only ordered types. Not thread-safe.
func NewMaxHeapOrdered[V constraints.Ordered](data []V) Heap[V] {
	std := StdInterface[V]{
		data: data,
	}

	std.LessFunc = func(i, j int) bool {
		return cmp.Less(std.data[j], std.data[i])
	}

	return NewHeap[V](&syncStl.LockerDummy{}, &std)
}

// NewMinHeapOrdered creates new min heap without limits.
// It supports only ordered types. Not thread-safe.
func NewMinHeapOrdered[V constraints.Ordered](data []V) Heap[V] {
	std := StdInterface[V]{
		data: data,
	}

	std.LessFunc = func(i, j int) bool {
		return cmp.Less(std.data[i], std.data[j])
	}

	return NewHeap[V](&syncStl.LockerDummy{}, &std)
}

// NewSyncMaxHeapOrdered creates new max heap without limits.
// It supports only ordered types. Not thread-safe.
func NewSyncMaxHeapOrdered[V constraints.Ordered](data []V) Heap[V] {
	std := StdInterface[V]{
		data: data,
	}

	std.LessFunc = func(i, j int) bool {
		return cmp.Less(std.data[j], std.data[i])
	}

	return NewHeap[V](&sync.Mutex{}, &std)
}

// NewSyncMinHeapOrdered creates new min heap without limits.
// It supports only ordered types. Thread-safe.
func NewSyncMinHeapOrdered[V constraints.Ordered](data []V) Heap[V] {
	std := StdInterface[V]{
		data: data,
	}

	std.LessFunc = func(i, j int) bool {
		return cmp.Less(std.data[i], std.data[j])
	}

	return NewHeap[V](&sync.Mutex{}, &std)
}

// NewHeapSync creates a new heap with provided StdInterface.
// Thread-safe.
func NewHeapSync[V any](stdInterface *StdInterface[V]) Heap[V] {
	return NewHeap(&sync.Mutex{}, stdInterface)
}

// NewHeapSimple creates a new heap with provided StdInterface.
// Not thread-safe.
func NewHeapSimple[V any](stdInterface *StdInterface[V]) Heap[V] {
	return NewHeap(&syncStl.LockerDummy{}, stdInterface)
}

// NewHeap creates a new heap with provided locker and StdInterface.
// Locker option can be used sync.Mutex or sync.LockerDummy or any other implementation of sync.Locker.
// Both parameters are required.
func NewHeap[V any](locker sync.Locker, stdInterface *StdInterface[V]) Heap[V] {
	if locker == nil {
		panic("Locker is required for NewHeapLocker")
	}

	if stdInterface == nil {
		panic("StdInterface is required for NewHeapLocker")
	}

	heapStd.Init(stdInterface)

	return &heap[V]{
		locker: locker,
		data:   stdInterface,
	}
}
