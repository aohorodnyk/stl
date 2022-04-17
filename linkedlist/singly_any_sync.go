package linkedlist

import "sync"

func NewSinglyAnySync[T any]() *SinglyAnySync[T] {
	return &SinglyAnySync[T]{
		linkedlist: NewSinglyAny[T](),
	}
}

func NewSinglyAnySyncCmp[T any](cmp func(T, T) bool) *SinglyAnySync[T] {
	return &SinglyAnySync[T]{
		linkedlist: NewSinglyAnyCmp(cmp),
	}
}

type SinglyAnySync[T any] struct {
	linkedlist *SinglyAny[T]
	mutex      sync.RWMutex
}

func (s *SinglyAnySync[T]) NodeFirst() (Node[T], bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeFirst()
}

func (s *SinglyAnySync[T]) NodeLast() (Node[T], bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeLast()
}

func (s *SinglyAnySync[T]) NodeAt(index int) (Node[T], bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeAt(index)
}

func (s *SinglyAnySync[T]) ValueFirst() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueFirst()
}

func (s *SinglyAnySync[T]) ValueLast() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueLast()
}

func (s *SinglyAnySync[T]) ValueAt(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueAt(index)
}

func (s *SinglyAnySync[T]) IndexOf(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOf(value)
}

func (s *SinglyAnySync[T]) IndexOfLast(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOfLast(value)
}

func (s *SinglyAnySync[T]) Contains(value T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Contains(value)
}

func (s *SinglyAnySync[T]) Length() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Length()
}

func (s *SinglyAnySync[T]) Empty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Empty()
}

func (s *SinglyAnySync[T]) AddFirst(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddFirst(value)
}

func (s *SinglyAnySync[T]) AddLast(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddLast(value)
}

func (s *SinglyAnySync[T]) AddAt(index int, value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddAt(index, value)
}

func (s *SinglyAnySync[T]) PopFirst() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopFirst()
}

func (s *SinglyAnySync[T]) PopLast() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopLast()
}

func (s *SinglyAnySync[T]) PopAt(index int) (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopAt(index)
}

func (s *SinglyAnySync[T]) RemoveNode(node Node[T]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveNode(node)
}

func (s *SinglyAnySync[T]) RemoveFirstBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveFirstBy(value)
}

func (s *SinglyAnySync[T]) RemoveLastBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveLastBy(value)
}

func (s *SinglyAnySync[T]) RemoveAllBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveAllBy(value)
}

func (s *SinglyAnySync[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.linkedlist.Clear()
}
