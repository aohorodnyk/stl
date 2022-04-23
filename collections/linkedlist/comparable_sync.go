package linkedlist

import "sync"

// AnySync is a concurrent safe linked list for comparable types T.
// It can be used to cover thread safe container over any custom LinkedList implementation.
func NewComparableSync[T comparable](list LinkedList[T]) *ComparableSync[T] {
	return &ComparableSync[T]{linkedlist: list}
}

// NewSinglyComparableSync returns a new ComparableSync for a singly linked list for comparable types T.
// The ComparableSync is not safe for concurrent use without additional synchronization.
// The result will be a singly linked list for comparable types T safe for concurrent use.
func NewSinglyComparableSync[T comparable]() *ComparableSync[T] {
	return &ComparableSync[T]{
		linkedlist: NewSinglyComparable[T](),
	}
}

// NewDoublyComparableSync returns a new ComparableSync for a doubly linked list for comparable types T.
// The ComparableSync is not safe for concurrent use without additional synchronization.
// The result will be a doubly linked list for comparable types T safe for concurrent use.
func NewDoublyComparableSync[T comparable]() *ComparableSync[T] {
	return &ComparableSync[T]{
		linkedlist: NewDoublyComparable[T](),
	}
}

// ComparableSync is a linked list container for comparable types T.
// It adds a synchornization layer to the linked list.
// All methods are safe for concurrent use.
// It uses RWMutex for synchronization, so read methods could be used in parallel.
type ComparableSync[T comparable] struct {
	linkedlist LinkedList[T]
	mutex      sync.RWMutex
}

// Lock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (c *ComparableSync[T]) Lock() {
	c.mutex.Lock()
}

// Unlock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (c *ComparableSync[T]) Unlock() {
	c.mutex.Unlock()
}

// RLock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (c *ComparableSync[T]) RLock() {
	c.mutex.RLock()
}

// RUnlock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (c *ComparableSync[T]) RUnlock() {
	c.mutex.RUnlock()
}

// NodeFirst returns the first node in the list.
// Returned node is not safe for concurrent use.
// If you would like to use plain Node, use RLock/RUnlock from the structure.
// See the details in the specific implementation.
func (c *ComparableSync[T]) NodeFirst() Node[T] {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.NodeFirst()
}

// NodeLast returns the last node in the list.
// Returned node is not safe for concurrent use.
// If you would like to use plain Node, use RLock/RUnlock from the structure.
// See the details in the specific implementation.
func (c *ComparableSync[T]) NodeLast() Node[T] {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.NodeLast()
}

// NodeAt returns the node at the given index.
// Returned node is not safe for concurrent use.
// If you would like to use plain Node, use RLock/RUnlock from the structure.
// See the details in the specific implementation.
func (c *ComparableSync[T]) NodeAt(index int) Node[T] {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.NodeAt(index)
}

// ValueFirst returns the first value in the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) ValueFirst() (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.ValueFirst()
}

// ValueLast returns the last value in the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) ValueLast() (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.ValueLast()
}

// ValueAt returns the value at the given index.
// See the details in the specific implementation.
func (c *ComparableSync[T]) ValueAt(index int) (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.ValueAt(index)
}

// IndexOf returns the first index of the given value.
// See the details in the specific implementation.
func (c *ComparableSync[T]) IndexOf(value T) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.IndexOf(value)
}

// IndexOfLast returns the last index of the given value.
// See the details in the specific implementation.
func (c *ComparableSync[T]) IndexOfLast(value T) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.IndexOfLast(value)
}

// Contains returns true if the given value is in the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) Contains(value T) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.Contains(value)
}

// Length returns the number of elements in the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) Length() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.Length()
}

// Empty returns true if the list is empty.
// See the details in the specific implementation.
func (c *ComparableSync[T]) Empty() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.Empty()
}

// AddFirst adds the given value to the front of the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) AddFirst(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.AddFirst(value)
}

// AddLast adds the given value to the end of the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) AddLast(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.AddLast(value)
}

// AddAt adds the given value at the given index.
// See the details in the specific implementation.
func (c *ComparableSync[T]) AddAt(index int, value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.AddAt(index, value)
}

// PopFirst removes and returns the first value in the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) PopFirst() (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.PopFirst()
}

// PopLast removes and returns the last value in the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) PopLast() (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.PopLast()
}

// PopAt removes and returns the value at the given index.
// See the details in the specific implementation.
func (c *ComparableSync[T]) PopAt(index int) (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.PopAt(index)
}

// RemoveNode removes the given node from the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) RemoveNode(node Node[T]) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveNode(node)
}

// RemoveFirstBy removes the first value that equals the given value.
// See the details in the specific implementation.
func (c *ComparableSync[T]) RemoveFirstBy(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveFirstBy(value)
}

// RemoveLastBy removes the last value that equals the given value.
// See the details in the specific implementation.
func (c *ComparableSync[T]) RemoveLastBy(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveLastBy(value)
}

// RemoveAllBy removes all values that equals the given value.
// See the details in the specific implementation.
func (c *ComparableSync[T]) RemoveAllBy(value T) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveAllBy(value)
}

// Clear removes all values from the list.
// See the details in the specific implementation.
func (c *ComparableSync[T]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.linkedlist.Clear()
}
