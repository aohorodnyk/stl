package queue

import "sync"

// NewSync returns a new container over given queue for any type T.
// This container is concurrent safe.
// nolint:varnamelen // It's good name for the generic type.
func NewSync[T any](queueObj Queue[T]) *Sync[T] {
	return &Sync[T]{
		queue: queueObj,
	}
}

// NewDynamicSync returns a new dynamic-size queue for any type T.
// This container is concurrent safe.
// nolint:varnamelen // It's good name for the generic type.
func NewDynamicSync[T any]() *Sync[T] {
	return &Sync[T]{
		queue: NewDynamic[T](),
	}
}

// NewFixedSync returns a new fixed-size queue for any type T.
// The queue is initialized with the given capacity.
// The capacity must be greater than zero.
// This container is concurrent safe.
// nolint:varnamelen // It's good name for the generic type.
func NewFixedSync[T any](capacity int) *Sync[T] {
	return &Sync[T]{
		queue: NewFixed[T](capacity),
	}
}

// Sync is a container for any Queue[T any] implementation to be used in concurrent environment.
//nolint:varnamelen // It's good name for the generic type.
type Sync[T any] struct {
	queue Queue[T]
	mutex sync.Mutex
}

// Push adds a new element to the top of the queue.
// The latest pushed element is the first one to be popped.
// Returns true if the element was pushed successfully.
// Push is a O(1) operation.
func (d *Sync[T]) Push(v T) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.queue.Push(v)
}

// Pop removes the top element from the queue and returns it.
// If the queue is empty, Pop returns the default value for the type T and false as a second return value.
// Pop is a O(1) operation.
func (d *Sync[T]) Pop() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.queue.Pop()
}

// Peek returns the top element from the queue without removing it.
// If the queue is empty, Peek returns the default value for the type T and false as a second return value.
// Peek is a O(1) operation.
func (d *Sync[T]) Peek() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.queue.Peek()
}

// Empty returns true if the queue is empty.
// Empty is a O(1) operation.
func (d *Sync[T]) Empty() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.queue.Empty()
}

// Length returns the number of elements in the queue.
// Length is a O(1) operation.
func (d *Sync[T]) Length() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.queue.Length()
}

// Clear removes all elements from the queue.
// Clear is a O(1) operation.
func (d *Sync[T]) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.queue.Clear()
}
