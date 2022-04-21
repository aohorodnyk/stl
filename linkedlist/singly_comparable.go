package linkedlist

import (
	"github.com/aohorodnyk/stl/math"
)

// NewSinglyComparable returns a new SinglyComparable linked list for comparable types T.
func NewSinglyComparable[T comparable]() *SinglyComparable[T] {
	return &SinglyComparable[T]{}
}

// SinglyComparable is a singly linked list for comparable types T.
// All search methods by value are compare values by direct comparison `==`.
// This implementation is much faster than the SinglyAny implementation with the default cmp method.
type SinglyComparable[T comparable] struct {
	head   *SinglyNodeComparable[T]
	length int
}

// NodeFirst returns the first node in the list.
// If the list is empty, it returns nil.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) NodeFirst() Node[T] {
	return s.NodeAt(0)
}

// NodeLast returns the last node in the list.
// If the list is empty, it returns nil.
// This method will go through the whole list for every call.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) NodeLast() Node[T] {
	return s.NodeAt(math.Max(s.length-1, 0))
}

// NodeAt returns the node at the given index.
// If index is out of range, it returns nil.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) NodeAt(index int) Node[T] {
	return s.nodeAt(index)
}

// ValueFirst returns the first value in the list.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) ValueFirst() (T, bool) {
	return s.ValueAt(0)
}

// ValueLast returns the last value in the list.
// This method will go through the whole list for every call.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) ValueLast() (T, bool) {
	return s.ValueAt(math.Max(s.length-1, 0))
}

// ValueAt returns the value at the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) ValueAt(index int) (T, bool) {
	if node := s.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

// IndexOf returns the first index of the given value.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) IndexOf(value T) int {
	pointer := s.head
	for index := 0; pointer != nil; index++ {
		if pointer.value == value {
			return index
		}

		pointer = pointer.next
	}

	return -1
}

// IndexOfLast returns the index of the last given value.
// This method will go through the whole list for every call.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) IndexOfLast(value T) int {
	result := -1
	pointer := s.head

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
func (s *SinglyComparable[T]) Contains(value T) bool {
	return s.IndexOf(value) != -1
}

// Length returns the length of the list.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) Length() int {
	return s.length
}

// Empty returns true if the list is empty.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) Empty() bool {
	return s.length == 0
}

// AddFirst adds the given value to the first position of the list.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) AddFirst(value T) bool {
	return s.AddAt(0, value)
}

// AddLast adds the given value to the last position of the list.
// This method will go through the whole list for every call.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) AddLast(value T) bool {
	return s.AddAt(s.length, value)
}

// AddAt adds the given value to the given index of the list.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) AddAt(index int, value T) bool {
	if index == 0 {
		s.length++
		s.head = &SinglyNodeComparable[T]{
			value: value,
			next:  s.head,
		}

		return true
	}

	node := s.nodeAt(index - 1)
	if node == nil {
		return false
	}

	s.length++

	node.next = &SinglyNodeComparable[T]{
		value: value,
		next:  node.next,
	}

	return true
}

// PopFirst removes the first node from the list and returns the value.
// Second return value is true if node has been removed from the list.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) PopFirst() (T, bool) {
	return s.PopAt(0)
}

// PopLast removes the last node from the list and returns the value.
// Second return value is true if node has been removed from the list.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) PopLast() (T, bool) {
	return s.PopAt(math.Max(s.length-1, 0))
}

// PopAt removes the node at the given index from the list and returns the value.
// Second return value is true if node has been removed from the list.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) PopAt(index int) (T, bool) {
	var result T

	if index == 0 {
		if s.head != nil {
			s.length--
			result = s.head.value
			s.head = s.head.next

			return result, true
		}

		return result, false
	}

	node := s.nodeAt(index - 1)
	if node == nil || node.next == nil {
		return result, false
	}

	s.length--

	result = node.next.value
	node.next = node.next.next

	return result, true
}

// RemoveNode removes the given node from the list.
// This method returns true if node has been removed from the list.
// Nodes compare by reference (pointer) `==`. It will compare exact references.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) RemoveNode(node Node[T]) bool {
	if node == nil {
		return false
	}

	if s.head == node {
		s.head = s.head.next
		s.length--

		return true
	}

	pointer := s.head
	for pointer != nil {
		if pointer.next == node {
			pointer.next = pointer.next.next
			s.length--

			return true
		}

		pointer = pointer.next
	}

	return false
}

// RemoveFirstBy removes the first node that has the given value.
// This method returns true if node has been removed from the list.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) RemoveFirstBy(value T) bool {
	if s.head == nil {
		return false
	}

	if s.head.value == value {
		s.head = s.head.next
		s.length--

		return true
	}

	pointer := s.head
	for pointer.next != nil {
		if pointer.next.value == value {
			pointer.next = pointer.next.next
			s.length--

			return true
		}

		pointer = pointer.next
	}

	return false
}

// RemoveLastBy removes the last node that has the given value.
// This method returns true if node has been removed from the list.
// This method will go through the whole list for every call.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) RemoveLastBy(value T) bool {
	if s.head == nil {
		return false
	}

	var pointerLast *SinglyNodeComparable[T]

	pointer := s.head
	for pointer.next != nil {
		if pointer.next.value == value {
			pointerLast = pointer
		}

		pointer = pointer.next
	}

	if pointerLast != nil {
		pointerLast.next = pointerLast.next.next
		s.length--

		return true
	}

	if s.head.value == value {
		s.head = s.head.next
		s.length--

		return true
	}

	return false
}

// RemoveAllBy removes all nodes that have the given value.
// This method returns the number of nodes that have been removed.
// This method will go through the whole list for every call.
// This method has O(n) performance complexity.
func (s *SinglyComparable[T]) RemoveAllBy(value T) int {
	var removed int

	if s.head == nil {
		return removed
	}

	pointer := s.head
	for pointer.next != nil {
		if pointer.next.value == value {
			pointer.next = pointer.next.next
			s.length--

			removed++
		}

		pointer = pointer.next
	}

	if s.head.value == value {
		s.head = s.head.next
		s.length--

		removed++
	}

	return removed
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (s *SinglyComparable[T]) Clear() {
	s.head = nil
	s.length = 0
}

func (s *SinglyComparable[T]) nodeAt(index int) *SinglyNodeComparable[T] {
	if index < 0 || index >= s.length {
		return nil
	}

	node := s.head
	for i := 0; i < index && node != nil; i++ {
		node = node.next
	}

	if node == nil {
		return nil
	}

	return node
}

// SinglyNodeComparable[T] is a node of a singly-linked list.
// It has a value and a pointer to the next node.
// It implements the Node interface.
type SinglyNodeComparable[T comparable] struct {
	value T
	next  *SinglyNodeComparable[T]
}

// Value returns the value of the node.
func (s *SinglyNodeComparable[T]) Value() T {
	return s.value
}

// Next returns the next node in the list.
func (s *SinglyNodeComparable[T]) Next() Node[T] {
	if s.next == nil {
		return nil
	}

	return s.next
}

// Prev throws a panic, because this method cannot be implemented for the Singly linked list.
func (s *SinglyNodeComparable[T]) Prev() Node[T] {
	panic("Singly linked list does not support prev, use doubly linked list instead")
}
