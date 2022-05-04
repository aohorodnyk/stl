package linkedlist

import (
	"github.com/aohorodnyk/stl/math"
	"github.com/aohorodnyk/stl/reflect"
)

// NewSingly returns a new Singly linked list for comparable types T.
func NewSingly[T comparable]() *Singly[T] {
	cmp := func(a, b T) bool {
		return a == b
	}

	return NewSinglyFunc(cmp)
}

// NewSinglyFunc returns a new Singly linked list for any types T.
// cmp is a function that will be used for comparisons.
// If the cmp function is nil, this function will throw panic.
func NewSinglyFunc[T any](cmp func(T, T) bool) *Singly[T] {
	if cmp == nil {
		panic("Comparable function is required")
	}

	return &Singly[T]{
		cmp: cmp,
	}
}

// NewSinglyFuncDeep returns a new Singly linked list for any types T.
// Comparisons will be done by `reflect.DeepEqual` function independently from the provided type.
func NewSinglyFuncDeep[T any]() *Singly[T] {
	cmp := func(a, b T) bool {
		return reflect.DeepEqual(a, b)
	}

	return NewSinglyFunc(cmp)
}

// Singly is a singly linked list for any types T.
// All search methods by value are compare values by a custom cmp function.
// If the type T is comparable, Singly[T comparable] implementation is suggest for use.
type Singly[T any] struct {
	head   *SinglyNode[T]
	length int
	cmp    func(T, T) bool
}

// NodeFirst returns the first node in the list.
// If the list is empty, it returns nil.
// This method has O(1) performance complexity.
func (s *Singly[T]) NodeFirst() Node[T] {
	return s.NodeAt(0)
}

// NodeLast returns the last node in the list.
// If the list is empty, it returns nil.
// This method has O(n) performance complexity.
func (s *Singly[T]) NodeLast() Node[T] {
	return s.NodeAt(math.Max(s.length-1, 0))
}

// NodeAt returns the node at the given index.
// If index is out of range, it returns nil.
// This method has O(n) performance complexity.
func (s *Singly[T]) NodeAt(index int) Node[T] {
	return s.nodeAt(index)
}

// ValueFirst returns the first value in the list.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (s *Singly[T]) ValueFirst() (T, bool) {
	return s.ValueAt(0)
}

// ValueLast returns the last value in the list.
// If the list is empty, it returns false.
// This method has O(n) performance complexity.
func (s *Singly[T]) ValueLast() (T, bool) {
	return s.ValueAt(math.Max(s.length-1, 0))
}

// ValueAt returns the value at the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (s *Singly[T]) ValueAt(index int) (T, bool) {
	if node := s.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

// IndexOf returns the index of the given value.
// If the value is not found, it returns -1.
// This method has O(n) performance complexity.
func (s *Singly[T]) IndexOf(value T) int {
	pointer := s.head
	for index := 0; pointer != nil; index++ {
		if s.cmp(pointer.value, value) {
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
func (s *Singly[T]) IndexOfLast(value T) int {
	result := -1
	pointer := s.head

	for index := 0; pointer != nil; index++ {
		if s.cmp(pointer.value, value) {
			result = index
		}

		pointer = pointer.next
	}

	return result
}

// Contains returns true if the list contains the given value.
// This method has O(n) performance complexity.
func (s *Singly[T]) Contains(value T) bool {
	return s.IndexOf(value) != -1
}

// Length returns the length of the list.
// This method has O(1) performance complexity.
func (s *Singly[T]) Length() int {
	return s.length
}

// Empty returns true if the list is empty.
// This method has O(1) performance complexity.
func (s *Singly[T]) Empty() bool {
	return s.length == 0
}

// AddFirst adds the given value to the first position of the list.
// This method has O(1) performance complexity.
func (s *Singly[T]) AddFirst(value T) bool {
	return s.AddAt(0, value)
}

// AddLast adds the given value to the last position of the list.
// This method has O(n) performance complexity.
func (s *Singly[T]) AddLast(value T) bool {
	return s.AddAt(s.length, value)
}

// AddAt adds the given value to the given index.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (s *Singly[T]) AddAt(index int, value T) bool {
	if index == 0 {
		s.length++
		s.head = &SinglyNode[T]{
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

	node.next = &SinglyNode[T]{
		value: value,
		next:  node.next,
	}

	return true
}

// PopFirst removes the first value from the list and return it.
// If the list is empty, it returns false.
// This method has O(1) performance complexity.
func (s *Singly[T]) PopFirst() (T, bool) {
	return s.PopAt(0)
}

// PopLast removes the last value from the list and return it.
// If the list is empty, it returns false.
// This method has O(n) performance complexity.
func (s *Singly[T]) PopLast() (T, bool) {
	return s.PopAt(math.Max(s.length-1, 0))
}

// PopAt removes the value at the given index and return it.
// If index is out of range, it returns false.
// This method has O(n) performance complexity.
func (s *Singly[T]) PopAt(index int) (T, bool) {
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
func (s *Singly[T]) RemoveNode(node Node[T]) bool {
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
func (s *Singly[T]) RemoveFirstBy(value T) bool {
	if s.head == nil {
		return false
	}

	if s.cmp(s.head.value, value) {
		s.head = s.head.next
		s.length--

		return true
	}

	pointer := s.head
	for pointer.next != nil {
		if s.cmp(pointer.next.value, value) {
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
func (s *Singly[T]) RemoveLastBy(value T) bool {
	if s.head == nil {
		return false
	}

	var pointerLast *SinglyNode[T]

	pointer := s.head
	for pointer.next != nil {
		if s.cmp(pointer.next.value, value) {
			pointerLast = pointer
		}

		pointer = pointer.next
	}

	if pointerLast != nil {
		pointerLast.next = pointerLast.next.next
		s.length--

		return true
	}

	if s.cmp(s.head.value, value) {
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
func (s *Singly[T]) RemoveAllBy(value T) int {
	var removed int

	if s.head == nil {
		return removed
	}

	pointer := s.head
	for pointer.next != nil {
		if s.cmp(pointer.next.value, value) {
			pointer.next = pointer.next.next
			s.length--

			removed++
		}

		pointer = pointer.next
	}

	if s.cmp(s.head.value, value) {
		s.head = s.head.next
		s.length--

		removed++
	}

	return removed
}

// Clear removes all nodes from the list.
// This method has O(1) performance complexity.
func (s *Singly[T]) Clear() {
	s.head = nil
	s.length = 0
}

func (s *Singly[T]) nodeAt(index int) *SinglyNode[T] {
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

// SinglyNode[T] is a node of a singly-linked list.
// It has a value and a pointer to the next node.
// It implements the Node interface.
type SinglyNode[T any] struct {
	value T
	next  *SinglyNode[T]
}

func (s *SinglyNode[T]) Value() T {
	return s.value
}

// Next returns the next node.
// If the node is the last node, it returns nil.
func (s *SinglyNode[T]) Next() Node[T] {
	if s.next == nil {
		return nil
	}

	return s.next
}

// Prev throws a panic, because this method cannot be implemented for the Singly linked list.
func (s *SinglyNode[T]) Prev() Node[T] {
	panic("Singly linked list does not support prev, use doubly linked list instead")
}
