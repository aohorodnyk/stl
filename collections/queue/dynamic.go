package queue

import "github.com/aohorodnyk/stl/collections/linkedlist"

func NewDynamicLinkedList[T any](list linkedlist.LinkedList[T]) *Dynamic[T] {
	return &Dynamic[T]{
		list: list,
	}
}

// NewDynamic returns a new dynamic queue for any type T.
// The queue is not safe for concurrent use.
// The queue based on NewDoublyAnyDeep linked list. The factory function does not use reflection.
// Queue does not use comparison functions.
func NewDynamic[T any]() *Dynamic[T] {
	return &Dynamic[T]{
		list: linkedlist.NewDoublyAnyDeep[T](),
	}
}

// Dynamic is a queue implementation based on linked list.
// This queue is not safe for concurrent use.
// All operations are performed in O(1) time.
//nolint:varnamelen // It's good name for the generic type.
type Dynamic[T any] struct {
	list linkedlist.LinkedList[T]
}

// Push adds a new element to the bottom of the queue.
// The latest pushed element is the first one to be popped.
// Returns true if the element was pushed successfully.
// Push is a O(1) operation.
func (d *Dynamic[Value]) Push(v Value) bool {
	return d.list.AddFirst(v)
}

// Pop removes the top element from the queue and returns it.
// If the queue is empty, Pop returns the default value for the type T and false as a second return value.
// Pop is a O(1) operation.
func (d *Dynamic[Value]) Pop() (Value, bool) {
	return d.list.PopLast()
}

// Peek returns the top element from the queue without removing it.
// If the queue is empty, Peek returns the default value for the type T and false as a second return value.
// Peek is a O(1) operation.
func (d *Dynamic[Value]) Peek() (Value, bool) {
	return d.list.ValueLast()
}

// Empty returns true if the queue is empty.
// Empty is a O(1) operation.
func (d *Dynamic[Value]) Empty() bool {
	return d.list.Empty()
}

// Length returns the number of elements in the queue.
// Length is a O(1) operation.
func (d *Dynamic[Value]) Length() int {
	return d.list.Length()
}

// Clear removes all elements from the queue.
// Length is a O(1) operation.
func (d *Dynamic[Value]) Clear() {
	d.list.Clear()
}
