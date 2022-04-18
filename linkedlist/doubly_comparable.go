package linkedlist

import "github.com/aohorodnyk/stl/math"

func NewDoublyComparable[T comparable]() *DoublyComparable[T] {
	return &DoublyComparable[T]{}
}

type DoublyComparable[T comparable] struct {
	head   *DoublyNodeComparable[T]
	tail   *DoublyNodeComparable[T]
	length int
}

func (s *DoublyComparable[T]) NodeFirst() Node[T] {
	return s.NodeAt(0)
}

func (s *DoublyComparable[T]) NodeLast() Node[T] {
	return s.NodeAt(math.Max(s.length-1, 0))
}

func (s *DoublyComparable[T]) NodeAt(index int) Node[T] {
	return s.nodeAt(index)
}

func (s *DoublyComparable[T]) ValueFirst() (T, bool) {
	return s.ValueAt(0)
}

func (s *DoublyComparable[T]) ValueLast() (T, bool) {
	return s.ValueAt(math.Max(s.length-1, 0))
}

func (s *DoublyComparable[T]) ValueAt(index int) (T, bool) {
	if node := s.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

func (s *DoublyComparable[T]) IndexOf(value T) int {
	pointer := s.head
	for index := 0; pointer != nil; index++ {
		if pointer.value == value {
			return index
		}

		pointer = pointer.next
	}

	return -1
}

func (s *DoublyComparable[T]) IndexOfLast(value T) int {
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

func (s *DoublyComparable[T]) Contains(value T) bool {
	return s.IndexOf(value) != -1
}

func (s *DoublyComparable[T]) Length() int {
	return s.length
}

func (s *DoublyComparable[T]) Empty() bool {
	return s.length == 0
}

func (s *DoublyComparable[T]) AddFirst(value T) bool {
	return s.AddAt(0, value)
}

func (s *DoublyComparable[T]) AddLast(value T) bool {
	return s.AddAt(s.length, value)
}

func (s *DoublyComparable[T]) AddAt(index int, value T) bool {
	if index == 0 {
		s.length++
		s.head = &DoublyNodeComparable[T]{
			value: value,
			next:  s.head,
		}

		if s.head.next == nil {
			s.tail = s.head
		} else {
			s.head.next.prev = s.head
		}

		return true
	}

	node := s.nodeAt(index - 1)
	if node == nil {
		return false
	}

	s.length++

	node.next = &DoublyNodeComparable[T]{
		value: value,
		next:  node.next,
		prev:  node,
	}

	if node.next.next == nil {
		s.tail = node.next
	}

	return true
}

func (s *DoublyComparable[T]) PopFirst() (T, bool) {
	return s.PopAt(0)
}

func (s *DoublyComparable[T]) PopLast() (T, bool) {
	return s.PopAt(math.Max(s.length-1, 0))
}

func (s *DoublyComparable[T]) PopAt(index int) (T, bool) {
	var result T

	node := s.nodeAt(math.Max(index, 0))
	if node != nil {
		result = node.value
		s.removeNode(node)

		return result, true
	}

	return result, false
}

func (s *DoublyComparable[T]) RemoveNode(node Node[T]) bool {
	if node == nil {
		return false
	}

	pointer := s.head
	for pointer != nil {
		if pointer == node {
			s.removeNode(pointer)

			return true
		}

		pointer = pointer.next
	}

	return false
}

func (s *DoublyComparable[T]) RemoveFirstBy(value T) bool {
	if s.head == nil {
		return false
	}

	pointer := s.head
	for pointer != nil {
		if pointer.value == value {
			s.removeNode(pointer)

			return true
		}

		pointer = pointer.next
	}

	return false
}

func (s *DoublyComparable[T]) RemoveLastBy(value T) bool {
	if s.tail == nil {
		return false
	}

	pointer := s.tail
	for pointer != nil {
		if pointer.value == value {
			s.removeNode(pointer)

			return true
		}

		pointer = pointer.prev
	}

	return false
}

func (s *DoublyComparable[T]) RemoveAllBy(value T) bool {
	var removed bool

	if s.head == nil {
		return removed
	}

	pointer := s.head
	for pointer != nil {
		if pointer.value == value {
			pointer = s.removeNode(pointer)

			removed = true
		} else {
			pointer = pointer.next
		}
	}

	return removed
}

func (s *DoublyComparable[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.length = 0
}

func (s *DoublyComparable[T]) removeNode(node *DoublyNodeComparable[T]) *DoublyNodeComparable[T] {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		s.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		s.tail = node.prev
	}

	s.length--

	s.syncEnds()

	if node.prev != nil {
		return node.prev
	}

	return node.next
}

func (s *DoublyComparable[T]) nodeAt(index int) *DoublyNodeComparable[T] {
	if index < 0 || index >= s.length {
		return nil
	}

	var node *DoublyNodeComparable[T]

	if middle := s.length / 2; index <= middle {
		node = s.head
		for i := 0; i < index && node != nil; i++ {
			node = node.next
		}
	} else {
		node = s.tail
		for i := s.length - 1; i > index && node != nil; i-- {
			node = node.prev
		}
	}

	if node == nil {
		return nil
	}

	return node
}

func (s *DoublyComparable[T]) syncEnds() {
	if s.head == nil || s.tail == nil {
		s.head = nil
		s.tail = nil
		s.length = 0
	}
}

type DoublyNodeComparable[T comparable] struct {
	value T
	next  *DoublyNodeComparable[T]
	prev  *DoublyNodeComparable[T]
}

func (s *DoublyNodeComparable[T]) Value() T {
	return s.value
}

func (s *DoublyNodeComparable[T]) Next() Node[T] {
	if s.next == nil {
		return nil
	}

	return s.next
}

func (s *DoublyNodeComparable[T]) Prev() Node[T] {
	if s.prev == nil {
		return nil
	}

	return s.prev
}
