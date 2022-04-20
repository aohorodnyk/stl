package linkedlist

import (
	reflectstd "reflect"

	"github.com/aohorodnyk/stl/math"
	"github.com/aohorodnyk/stl/reflect"
)

func NewSinglyAny[T any]() *SinglyAny[T] {
	var tmp T

	return &SinglyAny[T]{
		comparable: reflectstd.TypeOf(tmp).Comparable(),
	}
}

func NewSinglyAnyDeep[T any]() *SinglyAny[T] {
	return &SinglyAny[T]{
		comparable: false,
	}
}

func NewSinglyAnyCmp[T any](cmp func(T, T) bool) *SinglyAny[T] {
	if cmp == nil {
		panic("Comparable function is required")
	}

	return &SinglyAny[T]{
		cmp: cmp,
	}
}

type SinglyAny[T any] struct {
	head       *SinglyNodeAny[T]
	length     int
	cmp        func(T, T) bool
	comparable bool
}

func (s *SinglyAny[T]) NodeFirst() Node[T] {
	return s.NodeAt(0)
}

func (s *SinglyAny[T]) NodeLast() Node[T] {
	return s.NodeAt(math.Max(s.length-1, 0))
}

func (s *SinglyAny[T]) NodeAt(index int) Node[T] {
	return s.nodeAt(index)
}

func (s *SinglyAny[T]) ValueFirst() (T, bool) {
	return s.ValueAt(0)
}

func (s *SinglyAny[T]) ValueLast() (T, bool) {
	return s.ValueAt(math.Max(s.length-1, 0))
}

func (s *SinglyAny[T]) ValueAt(index int) (T, bool) {
	if node := s.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

func (s *SinglyAny[T]) IndexOf(value T) int {
	pointer := s.head
	for index := 0; pointer != nil; index++ {
		if s.compare(pointer.value, value) {
			return index
		}

		pointer = pointer.next
	}

	return -1
}

func (s *SinglyAny[T]) IndexOfLast(value T) int {
	result := -1
	pointer := s.head

	for index := 0; pointer != nil; index++ {
		if s.compare(pointer.value, value) {
			result = index
		}

		pointer = pointer.next
	}

	return result
}

func (s *SinglyAny[T]) Contains(value T) bool {
	return s.IndexOf(value) != -1
}

func (s *SinglyAny[T]) Length() int {
	return s.length
}

func (s *SinglyAny[T]) Empty() bool {
	return s.length == 0
}

func (s *SinglyAny[T]) AddFirst(value T) bool {
	return s.AddAt(0, value)
}

func (s *SinglyAny[T]) AddLast(value T) bool {
	return s.AddAt(s.length, value)
}

func (s *SinglyAny[T]) AddAt(index int, value T) bool {
	if index == 0 {
		s.length++
		s.head = &SinglyNodeAny[T]{
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

	node.next = &SinglyNodeAny[T]{
		value: value,
		next:  node.next,
	}

	return true
}

func (s *SinglyAny[T]) PopFirst() (T, bool) {
	return s.PopAt(0)
}

func (s *SinglyAny[T]) PopLast() (T, bool) {
	return s.PopAt(math.Max(s.length-1, 0))
}

func (s *SinglyAny[T]) PopAt(index int) (T, bool) {
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

func (s *SinglyAny[T]) RemoveNode(node Node[T]) bool {
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

func (s *SinglyAny[T]) RemoveFirstBy(value T) bool {
	if s.head == nil {
		return false
	}

	if s.compare(s.head.value, value) {
		s.head = s.head.next
		s.length--

		return true
	}

	pointer := s.head
	for pointer.next != nil {
		if s.compare(pointer.next.value, value) {
			pointer.next = pointer.next.next
			s.length--

			return true
		}

		pointer = pointer.next
	}

	return false
}

func (s *SinglyAny[T]) RemoveLastBy(value T) bool {
	if s.head == nil {
		return false
	}

	var pointerLast *SinglyNodeAny[T]

	pointer := s.head
	for pointer.next != nil {
		if s.compare(pointer.next.value, value) {
			pointerLast = pointer
		}

		pointer = pointer.next
	}

	if pointerLast != nil {
		pointerLast.next = pointerLast.next.next
		s.length--

		return true
	}

	if s.compare(s.head.value, value) {
		s.head = s.head.next
		s.length--

		return true
	}

	return false
}

func (s *SinglyAny[T]) RemoveAllBy(value T) bool {
	var removed bool

	if s.head == nil {
		return removed
	}

	pointer := s.head
	for pointer.next != nil {
		if s.compare(pointer.next.value, value) {
			pointer.next = pointer.next.next
			s.length--

			removed = true
		}

		pointer = pointer.next
	}

	if s.compare(s.head.value, value) {
		s.head = s.head.next
		s.length--

		removed = true
	}

	return removed
}

func (s *SinglyAny[T]) Clear() {
	s.head = nil
	s.length = 0
}

func (s *SinglyAny[T]) nodeAt(index int) *SinglyNodeAny[T] {
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

func (s *SinglyAny[T]) compare(a, b T) bool {
	if s.cmp != nil {
		return s.cmp(a, b)
	}

	return reflect.DeepEqual(a, b)
}

type SinglyNodeAny[T any] struct {
	value T
	next  *SinglyNodeAny[T]
}

func (s *SinglyNodeAny[T]) Value() T {
	return s.value
}

func (s *SinglyNodeAny[T]) Next() Node[T] {
	return s.next
}

func (s *SinglyNodeAny[T]) Prev() Node[T] {
	panic("Singly linked list does not support prev, use doubly linked list instead")
}
