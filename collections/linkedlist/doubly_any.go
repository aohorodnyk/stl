package linkedlist

import (
	reflectstd "reflect"

	"github.com/aohorodnyk/stl/math"
	"github.com/aohorodnyk/stl/reflect"
)

// NewDoublyAny returns a new DoublyAny linked list for any types T.
// Every call of this method will call the `reflect.TypeOf(T).Comparable()` method to speed up future comparison.
// If the type T is comparable, the comparison will work ~10 times faster than with non-comparable types.
// It will use direct comparison `==` for `any` types.
// If the type T is comparable, we suggest to use `NewDoublyComparable` instead, to use the fastest solution.
// If the type T is not comparable, the comparison will be done by `reflect.DeepEqual` function.
func NewDoublyAny[T any]() *DoublyAny[T] {
	var tmp T

	return &DoublyAny[T]{
		comparable: reflectstd.TypeOf(tmp).Comparable(),
	}
}

// NewDoublyAnyDeep returns a new DoublyAny linked list for any types T.
// Comparisons will be done by `reflect.DeepEqual` function independently from the provided type.
// We suggest to use this method only if you have a strong reason to choose exact this configuration.
func NewDoublyAnyDeep[T any]() *DoublyAny[T] {
	return &DoublyAny[T]{
		comparable: false,
	}
}

// NewDoublyAnyCmp returns a new DoublyAny linked list for any types T.
// cmp is a function that will be used for comparisons.
// If the cmp function is nil, this function will throw panic.
func NewDoublyAnyCmp[T any](cmp func(T, T) bool) *DoublyAny[T] {
	if cmp == nil {
		panic("Comparable function is required")
	}

	return &DoublyAny[T]{
		cmp: cmp,
	}
}

// DoublyAny is a doubly linked list for any types T.
// All search methods by value are compare values by a custom cmp function or stl/reflect.DeepEqual function.
// If the type T is comparable, DoublyComparable[T comparable] implementation is suggest for use.
type DoublyAny[T any] struct {
	head       *DoublyNodeAny[T]
	tail       *DoublyNodeAny[T]
	length     int
	cmp        func(T, T) bool
	comparable bool
}

// NodeFirst returns the first node in the list.
// If the list is empty, it returns nil.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) NodeFirst() Node[T] {
	return d.NodeAt(0)
}

// NodeLast returns the last node in the list.
// If the list is empty, it returns nil.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) NodeLast() Node[T] {
	return d.NodeAt(math.Max(d.length-1, 0))
}

// NodeAt returns the node at the given index.
// If index is out of range, it returns nil.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) NodeAt(index int) Node[T] {
	return d.nodeAt(index)
}

// ValueFirst returns the first value in the list.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) ValueFirst() (T, bool) {
	return d.ValueAt(0)
}

// ValueLast returns the last value in the list.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) ValueLast() (T, bool) {
	return d.ValueAt(math.Max(d.length-1, 0))
}

// ValueAt returns the value at the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) ValueAt(index int) (T, bool) {
	if node := d.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

// IndexOf returns the first index of the given value.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) IndexOf(value T) int {
	pointer := d.head
	for index := 0; pointer != nil; index++ {
		if d.compare(pointer.value, value) {
			return index
		}

		pointer = pointer.next
	}

	return -1
}

// IndexOfLast returns the last index of the given value.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) IndexOfLast(value T) int {
	result := -1
	pointer := d.head

	for index := 0; pointer != nil; index++ {
		if d.compare(pointer.value, value) {
			result = index
		}

		pointer = pointer.next
	}

	return result
}

// Contains returns true if the list contains the given value.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) Contains(value T) bool {
	return d.IndexOf(value) != -1
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) Length() int {
	return d.length
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) Empty() bool {
	return d.length == 0
}

// AddFirst adds a new node with the given value to the beginning of the list.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) AddFirst(value T) bool {
	return d.AddAt(0, value)
}

// AddLast adds a new node with the given value to the end of the list.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) AddLast(value T) bool {
	return d.AddAt(d.length, value)
}

// AddAt adds a new node with the given value at the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) AddAt(index int, value T) bool {
	if index == 0 {
		d.length++
		d.head = &DoublyNodeAny[T]{
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

	node.next = &DoublyNodeAny[T]{
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
func (d *DoublyAny[T]) PopFirst() (T, bool) {
	return d.PopAt(0)
}

// PopLast removes the last node from the list and returns its value.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (d *DoublyAny[T]) PopLast() (T, bool) {
	return d.PopAt(math.Max(d.length-1, 0))
}

// PopAt removes the node at the given index and returns its value.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (d *DoublyAny[T]) PopAt(index int) (T, bool) {
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
func (d *DoublyAny[T]) RemoveNode(node Node[T]) bool {
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
func (d *DoublyAny[T]) RemoveFirstBy(value T) bool {
	if d.head == nil {
		return false
	}

	pointer := d.head
	for pointer != nil {
		if d.compare(pointer.value, value) {
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
func (d *DoublyAny[T]) RemoveLastBy(value T) bool {
	if d.tail == nil {
		return false
	}

	pointer := d.tail
	for pointer != nil {
		if d.compare(pointer.value, value) {
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
func (d *DoublyAny[T]) RemoveAllBy(value T) int {
	var removed int

	if d.head == nil {
		return removed
	}

	pointer := d.head
	for pointer != nil {
		if d.compare(pointer.value, value) {
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
func (d *DoublyAny[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.length = 0
}

func (d *DoublyAny[T]) removeNode(node *DoublyNodeAny[T]) *DoublyNodeAny[T] {
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

func (d *DoublyAny[T]) nodeAt(index int) *DoublyNodeAny[T] {
	if index < 0 || index >= d.length {
		return nil
	}

	var node *DoublyNodeAny[T]

	if middle := d.length / 2; index < middle {
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

func (d *DoublyAny[T]) syncEnds() {
	if d.head == nil || d.tail == nil {
		d.head = nil
		d.tail = nil
		d.length = 0
	}
}

func (d *DoublyAny[T]) compare(a, b T) bool {
	if d.cmp != nil {
		return d.cmp(a, b)
	}

	return reflect.DeepEqual(a, b)
}

// DoublyNodeAny[T] is a node of a doubly-linked list.
// It has a value and a pointer to the next and previous nodes.
// It implements the Node interface.
type DoublyNodeAny[T any] struct {
	value T
	next  *DoublyNodeAny[T]
	prev  *DoublyNodeAny[T]
}

// Value returns the value of the node.
func (d *DoublyNodeAny[T]) Value() T {
	return d.value
}

// Next returns the next node.
// If the node is the last node, it returns nil.
func (d *DoublyNodeAny[T]) Next() Node[T] {
	if d.next == nil {
		return nil
	}

	return d.next
}

// Prev returns the previous node.
// If the node is the first node, it returns nil.
func (d *DoublyNodeAny[T]) Prev() Node[T] {
	if d.prev == nil {
		return nil
	}

	return d.prev
}
