package linkedlist

import (
	"github.com/aohorodnyk/stl/math"
)

func NewSinglyComparable[T comparable]() *SinglyComparable[T] {
	return &SinglyComparable[T]{}
}

type SinglyComparable[T comparable] struct {
	head   *SinglyNodeComparable[T]
	length int
}

func (s *SinglyComparable[T]) NodeFirst() (Node[T], bool) {
	return s.NodeAt(0)
}

func (s *SinglyComparable[T]) NodeLast() (Node[T], bool) {
	return s.NodeAt(math.Max(s.length-1, 0))
}

func (s *SinglyComparable[T]) NodeAt(index int) (Node[T], bool) {
	node := s.nodeAt(index)

	return node, node != nil
}

func (s *SinglyComparable[T]) ValueFirst() (T, bool) {
	return s.ValueAt(0)
}

func (s *SinglyComparable[T]) ValueLast() (T, bool) {
	return s.ValueAt(math.Max(s.length-1, 0))
}

func (s *SinglyComparable[T]) ValueAt(index int) (T, bool) {
	if node := s.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

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

func (s *SinglyComparable[T]) Contains(value T) bool {
	return s.IndexOf(value) != -1
}

func (s *SinglyComparable[T]) Length() int {
	return s.length
}

func (s *SinglyComparable[T]) Empty() bool {
	return s.length == 0
}

func (s *SinglyComparable[T]) AddFirst(value T) bool {
	return s.AddAt(0, value)
}

func (s *SinglyComparable[T]) AddLast(value T) bool {
	return s.AddAt(s.length, value)
}

func (s *SinglyComparable[T]) AddAt(index int, value T) bool {
	if index > s.length {
		return false
	}

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

func (s *SinglyComparable[T]) PopFirst() (T, bool) {
	return s.PopAt(0)
}

func (s *SinglyComparable[T]) PopLast() (T, bool) {
	return s.PopAt(math.Max(s.length-1, 0))
}

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

func (s *SinglyComparable[T]) RemoveLastBy(value T) bool {
	if s.head == nil {
		return false
	}

	if s.head.value == value {
		s.head = s.head.next
		s.length--

		return true
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

	return false
}

func (s *SinglyComparable[T]) RemoveAllBy(value T) bool {
	var removed bool

	if s.head == nil {
		return removed
	}

	pointer := s.head
	for pointer.next != nil {
		if pointer.next.value == value {
			pointer.next = pointer.next.next
			s.length--

			removed = true
		}

		pointer = pointer.next
	}

	if s.head.value == value {
		s.head = s.head.next
		s.length--

		removed = true
	}

	return removed
}

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

	return node
}

type SinglyNodeComparable[T comparable] struct {
	value T
	next  *SinglyNodeComparable[T]
}

func (s *SinglyNodeComparable[T]) Value() T {
	return s.value
}

func (s *SinglyNodeComparable[T]) Next() Node[T] {
	return s.next
}

func (s *SinglyNodeComparable[T]) Prev() Node[T] {
	panic("Singly linked list does not support prev, use doubly linked list instead")
}
