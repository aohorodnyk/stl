package stack

import "sync"

// NewSync returns a new container over given stack for any type T.
// This container is concurrent safe.
// nolint:varnamelen // It's good name for the generic type.
func NewSync[T any](stackObj Stack[T]) *Sync[T] {
	return &Sync[T]{
		stack: stackObj,
	}
}

// NewDynamicSync returns a new dynamic-size stack for any type T.
// This container is concurrent safe.
// nolint:varnamelen // It's good name for the generic type.
func NewDynamicSync[T any]() *Sync[T] {
	return &Sync[T]{
		stack: NewDynamic[T](),
	}
}

// NewFixedSync returns a new fixed-size stack for any type T.
// The stack is initialized with the given capacity.
// The capacity must be greater than zero.
// This container is concurrent safe.
// nolint:varnamelen // It's good name for the generic type.
func NewFixedSync[T any](capacity int) *Sync[T] {
	return &Sync[T]{
		stack: NewFixed[T](capacity),
	}
}

// Sync is a container for any Stack[T any] implementation to be used in concurrent environment.
//nolint:varnamelen // It's good name for the generic type.
type Sync[T any] struct {
	stack Stack[T]
	mutex sync.Mutex
}

// Push adds a new element to the top of the stack.
// The latest pushed element is the first one to be popped.
// Returns true if the element was pushed successfully.
// Push is a O(1) operation.
func (d *Sync[T]) Push(v T) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.stack.Push(v)
}

// Pop removes the top element from the stack and returns it.
// If the stack is empty, Pop returns the default value for the type T and false as a second return value.
// Pop is a O(1) operation.
func (d *Sync[T]) Pop() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.stack.Pop()
}

// Peek returns the top element from the stack without removing it.
// If the stack is empty, Peek returns the default value for the type T and false as a second return value.
// Peek is a O(1) operation.
func (d *Sync[T]) Peek() (T, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.stack.Peek()
}

// Empty returns true if the stack is empty.
// Empty is a O(1) operation.
func (d *Sync[T]) Empty() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.stack.Empty()
}

// Length returns the number of elements in the stack.
// Length is a O(1) operation.
func (d *Sync[T]) Length() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.stack.Length()
}

// Clear removes all elements from the stack.
// Clear is a O(1) operation.
func (d *Sync[T]) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.stack.Clear()
}
