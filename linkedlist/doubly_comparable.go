package linkedlist

import "github.com/aohorodnyk/stl/math"

// NewDoublyComparable returns a new DoublyComparable linked list for comparable types T.
func NewDoublyComparable[T comparable]() *DoublyComparable[T] {
	return &DoublyComparable[T]{}
}

// DoublyComparable is a doubly linked list for comparable types T.
// All search methods by value are compare values by direct comparison `==`.
// This implementation is much faster than the DoublyAny implementation with the default cmp method.
type DoublyComparable[T comparable] struct {
	head   *DoublyNodeComparable[T]
	tail   *DoublyNodeComparable[T]
	length int
}

// NodeFirst returns the first node in the list.
// If the list is empty, it returns nil.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) NodeFirst() Node[T] {
	return d.NodeAt(0)
}

// NodeLast returns the last node in the list.
// If the list is empty, it returns nil.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) NodeLast() Node[T] {
	return d.NodeAt(math.Max(d.length-1, 0))
}

// NodeAt returns the node at the given index.
// If index is out of range, it returns nil.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) NodeAt(index int) Node[T] {
	return d.nodeAt(index)
}

// ValueFirst returns the first value in the list.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) ValueFirst() (T, bool) {
	return d.ValueAt(0)
}

// ValueLast returns the last value in the list.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) ValueLast() (T, bool) {
	return d.ValueAt(math.Max(d.length-1, 0))
}

// ValueAt returns the value at the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) ValueAt(index int) (T, bool) {
	if node := d.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

// IndexOf returns the first index of the given value.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) IndexOf(value T) int {
	pointer := d.head
	for index := 0; pointer != nil; index++ {
		if pointer.value == value {
			return index
		}

		pointer = pointer.next
	}

	return -1
}

// IndexOfLast returns the last index of the given value.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) IndexOfLast(value T) int {
	result := -1
	pointer := d.head

	for index := 0; pointer != nil; index++ {
		if pointer.value == value {
			result = index
		}

		pointer = pointer.next
	}

	return result
}

// Contains returns true if the list contains the given value.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) Contains(value T) bool {
	return d.IndexOf(value) != -1
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) Length() int {
	return d.length
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) Empty() bool {
	return d.length == 0
}

// AddFirst adds a new node with the given value to the beginning of the list.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) AddFirst(value T) bool {
	return d.AddAt(0, value)
}

// AddLast adds a new node with the given value to the end of the list.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) AddLast(value T) bool {
	return d.AddAt(d.length, value)
}

// AddAt adds a new node with the given value at the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) AddAt(index int, value T) bool {
	if index == 0 {
		d.length++
		d.head = &DoublyNodeComparable[T]{
			value: value,
			next:  d.head,
		}

		if d.head.next == nil {
			d.tail = d.head
		} else {
			d.head.next.prev = d.head
		}

		return true
	}

	node := d.nodeAt(index - 1)
	if node == nil {
		return false
	}

	d.length++

	node.next = &DoublyNodeComparable[T]{
		value: value,
		next:  node.next,
		prev:  node,
	}

	if node.next.next == nil {
		d.tail = node.next
	} else {
		node.next.next.prev = node.next
	}

	return true
}

// PopFirts removes the first node from the list and returns its value.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) PopFirst() (T, bool) {
	return d.PopAt(0)
}

// PopLast removes the last node from the list and returns its value.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) PopLast() (T, bool) {
	return d.PopAt(math.Max(d.length-1, 0))
}

// PopAt removes the node at the given index and returns its value.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) PopAt(index int) (T, bool) {
	var result T

	node := d.nodeAt(math.Max(index, 0))
	if node != nil {
		result = node.value
		d.removeNode(node)

		return result, true
	}

	return result, false
}

// RemoveNode removes the given node from the list.
// This method returns true if node has been removed from the list.
// Nodes compare by reference (pointer) `==`. It will compare exact references.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) RemoveNode(node Node[T]) bool {
	if node == nil {
		return false
	}

	pointer := d.head
	for pointer != nil {
		if pointer == node {
			d.removeNode(pointer)

			return true
		}

		pointer = pointer.next
	}

	return false
}

// RemoveFirstBy removes the first node with the given value from the list.
// This method returns true if the node has been removed from the list.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) RemoveFirstBy(value T) bool {
	if d.head == nil {
		return false
	}

	pointer := d.head
	for pointer != nil {
		if pointer.value == value {
			d.removeNode(pointer)

			return true
		}

		pointer = pointer.next
	}

	return false
}

// RemoveLastBy removes the last node with the given value from the list.
// This method returns true if the node has been removed from the list.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) RemoveLastBy(value T) bool {
	if d.tail == nil {
		return false
	}

	pointer := d.tail
	for pointer != nil {
		if pointer.value == value {
			d.removeNode(pointer)

			return true
		}

		pointer = pointer.prev
	}

	return false
}

// RemoveAllBy removes all nodes that have the given value.
// This method returns the number of nodes that have been removed.
// This method will go through the whole list for every call.
// This method has O(n) performance complexity.
func (d *DoublyComparable[T]) RemoveAllBy(value T) int {
	var removed int

	if d.head == nil {
		return removed
	}

	pointer := d.head
	for pointer != nil {
		if pointer.value == value {
			pointer = d.removeNode(pointer)

			removed++
		} else {
			pointer = pointer.next
		}
	}

	return removed
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (d *DoublyComparable[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.length = 0
}

func (d *DoublyComparable[T]) removeNode(node *DoublyNodeComparable[T]) *DoublyNodeComparable[T] {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		d.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		d.tail = node.prev
	}

	d.length--

	d.syncEnds()

	if node.prev != nil {
		return node.prev
	}

	return node.next
}

func (d *DoublyComparable[T]) nodeAt(index int) *DoublyNodeComparable[T] {
	if index < 0 || index >= d.length {
		return nil
	}

	var node *DoublyNodeComparable[T]

	if middle := d.length / 2; index <= middle {
		node = d.head
		for i := 0; i < index && node != nil; i++ {
			node = node.next
		}
	} else {
		node = d.tail
		for i := d.length - 1; i > index && node != nil; i-- {
			node = node.prev
		}
	}

	if node == nil {
		return nil
	}

	return node
}

func (d *DoublyComparable[T]) syncEnds() {
	if d.head == nil || d.tail == nil {
		d.head = nil
		d.tail = nil
		d.length = 0
	}
}

// DoublyNodeComparable[T] is a node of a doubly-linked list.
// It has a value and a pointer to the next and previous nodes.
// It implements the Node interface.
type DoublyNodeComparable[T comparable] struct {
	value T
	next  *DoublyNodeComparable[T]
	prev  *DoublyNodeComparable[T]
}

// Value returns the value of the node.
func (d *DoublyNodeComparable[T]) Value() T {
	return d.value
}

// Next returns the next node.
// If the node is the last node, it returns nil.
func (d *DoublyNodeComparable[T]) Next() Node[T] {
	if d.next == nil {
		return nil
	}

	return d.next
}

// Prev returns the previous node.
// If the node is the first node, it returns nil.
func (d *DoublyNodeComparable[T]) Prev() Node[T] {
	if d.prev == nil {
		return nil
	}

	return d.prev
}
