package linkedlist

import "sync"

// NewSinglyAnySync returns a new Singly linked list for any type T.
// Current instance of the LinkedList will be concurrent safe.
func NewSinglyAnySync[T any]() *AnySync[T] {
	return &AnySync[T]{
		linkedlist: NewSinglyAny[T](),
	}
}

// NewSinglyAnySyncDeep returns a new Singly linked list for any type T with reflect.DeepEqual comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewSinglyAnySyncDeep[T any]() *AnySync[T] {
	return &AnySync[T]{
		linkedlist: NewSinglyAnyDeep[T](),
	}
}

// NewSinglyAnySyncCmp returns a new Singly linked list for any type T with the given comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewSinglyAnySyncCmp[T any](cmp func(T, T) bool) *AnySync[T] {
	return &AnySync[T]{
		linkedlist: NewSinglyAnyCmp(cmp),
	}
}

// NewDoublyAnySync returns a new Doubly linked list for any type T.
// Current instance of the LinkedList will be concurrent safe.
func NewDoublyAnySync[T any]() *AnySync[T] {
	return &AnySync[T]{
		linkedlist: NewDoublyAny[T](),
	}
}

// NewDoublyAnySyncDeep returns a new Doubly linked list for any type T with reflect.DeepEqual comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewDoublyAnySyncDeep[T any]() *AnySync[T] {
	return &AnySync[T]{
		linkedlist: NewDoublyAnyDeep[T](),
	}
}

// NewDoublyAnySyncCmp returns a new Doubly linked list for any type T with the given comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewDoublyAnySyncCmp[T any](cmp func(T, T) bool) *AnySync[T] {
	return &AnySync[T]{
		linkedlist: NewDoublyAnyCmp(cmp),
	}
}

// AnySync is a concurrent safe linked list for any type T.
// It implements the LinkedList interface with the dependent of any LinkedList implementation.
// This implementation is based on sync.RWMutex to make any implementation thread safe.
type AnySync[T any] struct {
	linkedlist LinkedList[T]
	mutex      sync.RWMutex
}

// NodeFirst returns the first node in the list.
func (s *AnySync[T]) NodeFirst() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeFirst()
}

// NodeLast returns the last node in the list.
func (s *AnySync[T]) NodeLast() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeLast()
}

// NodeAt returns the node at the given index.
func (s *AnySync[T]) NodeAt(index int) Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeAt(index)
}

func (s *AnySync[T]) ValueFirst() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueFirst()
}

func (s *AnySync[T]) ValueLast() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueLast()
}

func (s *AnySync[T]) ValueAt(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueAt(index)
}

func (s *AnySync[T]) IndexOf(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOf(value)
}

func (s *AnySync[T]) IndexOfLast(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOfLast(value)
}

func (s *AnySync[T]) Contains(value T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Contains(value)
}

func (s *AnySync[T]) Length() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Length()
}

func (s *AnySync[T]) Empty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Empty()
}

func (s *AnySync[T]) AddFirst(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddFirst(value)
}

func (s *AnySync[T]) AddLast(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddLast(value)
}

func (s *AnySync[T]) AddAt(index int, value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddAt(index, value)
}

func (s *AnySync[T]) PopFirst() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopFirst()
}

func (s *AnySync[T]) PopLast() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopLast()
}

func (s *AnySync[T]) PopAt(index int) (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopAt(index)
}

func (s *AnySync[T]) RemoveNode(node Node[T]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveNode(node)
}

func (s *AnySync[T]) RemoveFirstBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveFirstBy(value)
}

func (s *AnySync[T]) RemoveLastBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveLastBy(value)
}

func (s *AnySync[T]) RemoveAllBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveAllBy(value)
}

func (s *AnySync[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.linkedlist.Clear()
}
