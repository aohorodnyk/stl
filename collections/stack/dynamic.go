package stack

import "github.com/aohorodnyk/stl/collections/linkedlist"

func NewDynamicLinkedList[T any](list linkedlist.LinkedList[T]) *Dynamic[T] {
	return &Dynamic[T]{
		list: list,
	}
}

// NewDynamic returns a new dynamic stack for any type T.
// The stack is not safe for concurrent use.
// The stack based on NewSinglyAnyDeep linked list. The factory function does not use reflection.
// Stack does not use comparison functions.
func NewDynamic[T any]() *Dynamic[T] {
	return &Dynamic[T]{
		list: linkedlist.NewSinglyAnyDeep[T](),
	}
}

// Dynamic is a stack implementation based on linked list.
// This stack is not safe for concurrent use.
// All operations are performed in O(1) time.
//nolint:varnamelen // It's good name for the generic type.
type Dynamic[T any] struct {
	list linkedlist.LinkedList[T]
}

// Push adds a new element to the top of the stack.
// The latest pushed element is the first one to be popped.
// Returns true if the element was pushed successfully.
// Push is a O(1) operation.
func (d *Dynamic[Value]) Push(v Value) bool {
	return d.list.AddFirst(v)
}

// Pop removes the top element from the stack and returns it.
// If the stack is empty, Pop returns the default value for the type T and false as a second return value.
// Pop is a O(1) operation.
func (d *Dynamic[Value]) Pop() (Value, bool) {
	return d.list.PopFirst()
}

// Peek returns the top element from the stack without removing it.
// If the stack is empty, Peek returns the default value for the type T and false as a second return value.
// Peek is a O(1) operation.
func (d *Dynamic[Value]) Peek() (Value, bool) {
	return d.list.ValueFirst()
}

// Empty returns true if the stack is empty.
// Empty is a O(1) operation.
func (d *Dynamic[Value]) Empty() bool {
	return d.list.Empty()
}

// Length returns the number of elements in the stack.
// Length is a O(1) operation.
func (d *Dynamic[Value]) Length() int {
	return d.list.Length()
}

// Clear removes all elements from the stack.
// Length is a O(1) operation.
func (d *Dynamic[Value]) Clear() {
	d.list.Clear()
}
