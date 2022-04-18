package linkedlist

import "sync"

func NewSinglyComparableSync[T comparable]() *ComparableSync[T] {
	return &ComparableSync[T]{
		linkedlist: NewSinglyComparable[T](),
	}
}

func NewDoublyComparableSync[T comparable]() *ComparableSync[T] {
	return &ComparableSync[T]{
		linkedlist: NewDoublyComparable[T](),
	}
}

type ComparableSync[T comparable] struct {
	linkedlist LinkedList[T]
	mutex      sync.RWMutex
}

func (s *ComparableSync[T]) NodeFirst() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeFirst()
}

func (s *ComparableSync[T]) NodeLast() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeLast()
}

func (s *ComparableSync[T]) NodeAt(index int) Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeAt(index)
}

func (s *ComparableSync[T]) ValueFirst() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueFirst()
}

func (s *ComparableSync[T]) ValueLast() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueLast()
}

func (s *ComparableSync[T]) ValueAt(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueAt(index)
}

func (s *ComparableSync[T]) IndexOf(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOf(value)
}

func (s *ComparableSync[T]) IndexOfLast(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOfLast(value)
}

func (s *ComparableSync[T]) Contains(value T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Contains(value)
}

func (s *ComparableSync[T]) Length() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Length()
}

func (s *ComparableSync[T]) Empty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Empty()
}

func (s *ComparableSync[T]) AddFirst(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddFirst(value)
}

func (s *ComparableSync[T]) AddLast(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddLast(value)
}

func (s *ComparableSync[T]) AddAt(index int, value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddAt(index, value)
}

func (s *ComparableSync[T]) PopFirst() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopFirst()
}

func (s *ComparableSync[T]) PopLast() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopLast()
}

func (s *ComparableSync[T]) PopAt(index int) (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopAt(index)
}

func (s *ComparableSync[T]) RemoveNode(node Node[T]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveNode(node)
}

func (s *ComparableSync[T]) RemoveFirstBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveFirstBy(value)
}

func (s *ComparableSync[T]) RemoveLastBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveLastBy(value)
}

func (s *ComparableSync[T]) RemoveAllBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveAllBy(value)
}

func (s *ComparableSync[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.linkedlist.Clear()
}
