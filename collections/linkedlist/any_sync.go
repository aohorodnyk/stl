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

// AnySync is a linked list container for any types T.
// It adds a synchornization layer to the linked list.
// All methods are safe for concurrent use.
// It uses RWMutex for synchronization, so read methods could be used in parallel.
type AnySync[T any] struct {
	linkedlist LinkedList[T]
	mutex      sync.RWMutex
}

// NodeFirst returns the first node in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) NodeFirst() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeFirst()
}

// NodeLast returns the last node in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) NodeLast() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeLast()
}

// NodeAt returns the node at the given index.
// See the details in the specific implementation.
func (s *AnySync[T]) NodeAt(index int) Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeAt(index)
}

// ValueFirst returns the first value in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) ValueFirst() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueFirst()
}

// ValueLast returns the last value in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) ValueLast() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueLast()
}

// ValueAt returns the value at the given index.
// See the details in the specific implementation.
func (s *AnySync[T]) ValueAt(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueAt(index)
}

// IndexOf returns the first index of the given value.
// See the details in the specific implementation.
func (s *AnySync[T]) IndexOf(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOf(value)
}

// IndexOfLast returns the last index of the given value.
// See the details in the specific implementation.
func (s *AnySync[T]) IndexOfLast(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOfLast(value)
}

// Contains returns true if the given value is in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) Contains(value T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Contains(value)
}

// Length returns the number of elements in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) Length() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Length()
}

// Empty returns true if the list is empty.
// See the details in the specific implementation.
func (s *AnySync[T]) Empty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Empty()
}

// AddFirst adds the given value to the front of the list.
// See the details in the specific implementation.
func (s *AnySync[T]) AddFirst(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddFirst(value)
}

// AddLast adds the given value to the end of the list.
// See the details in the specific implementation.
func (s *AnySync[T]) AddLast(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddLast(value)
}

// AddAt adds the given value at the given index.
// See the details in the specific implementation.
func (s *AnySync[T]) AddAt(index int, value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddAt(index, value)
}

// PopFirst removes and returns the first value in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) PopFirst() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopFirst()
}

// PopLast removes and returns the last value in the list.
// See the details in the specific implementation.
func (s *AnySync[T]) PopLast() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopLast()
}

// PopAt removes and returns the value at the given index.
// See the details in the specific implementation.
func (s *AnySync[T]) PopAt(index int) (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopAt(index)
}

// RemoveNode removes the given node from the list.
// See the details in the specific implementation.
func (s *AnySync[T]) RemoveNode(node Node[T]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveNode(node)
}

// RemoveFirstBy removes the first value that equals the given value.
// See the details in the specific implementation.
func (s *AnySync[T]) RemoveFirstBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveFirstBy(value)
}

// RemoveLastBy removes the last value that equals the given value.
// See the details in the specific implementation.
func (s *AnySync[T]) RemoveLastBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveLastBy(value)
}

// RemoveAllBy removes all values that equals the given value.
// See the details in the specific implementation.
func (s *AnySync[T]) RemoveAllBy(value T) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveAllBy(value)
}

// Clear removes all nodes from the list.
// See the details in the specific implementation.
func (s *AnySync[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.linkedlist.Clear()
}
