package linkedlist

import "sync"

// AnySync is a concurrent safe linked list for any types T.
// It can be used to cover thread safe container over any custom LinkedList implementation.
func NewAnySync[T any](list LinkedList[T]) *SyncAny[T] {
	return &SyncAny[T]{linkedlist: list}
}

// NewSinglyFuncSync returns a new Singly linked list for any type T with the given comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewSinglyFuncSync[T any](cmp func(T, T) bool) *SyncAny[T] {
	return &SyncAny[T]{
		linkedlist: NewSinglyFunc(cmp),
	}
}

// NewSinglyFuncDeepSync returns a new Singly linked list for any type T with reflect.DeepEqual comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewSinglyFuncDeepSync[T any]() *SyncAny[T] {
	return &SyncAny[T]{
		linkedlist: NewSinglyFuncDeep[T](),
	}
}

// NewDoublyFuncSync returns a new Doubly linked list for any type T with the given comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewDoublyFuncSync[T any](cmp func(T, T) bool) *SyncAny[T] {
	return &SyncAny[T]{
		linkedlist: NewDoublyFunc(cmp),
	}
}

// NewDoublyFuncDeepSync returns a new Doubly linked list for any type T with reflect.DeepEqual comparting function.
// Current instance of the LinkedList will be concurrent safe.
func NewDoublyFuncDeepSync[T any]() *SyncAny[T] {
	return &SyncAny[T]{
		linkedlist: NewDoublyFuncDeep[T](),
	}
}

// SyncAny is a linked list container for any types T.
// It adds a synchornization layer to the linked list.
// All methods are safe for concurrent use.
// It uses RWMutex for synchronization, so read methods could be used in parallel.
type SyncAny[T any] struct {
	linkedlist LinkedList[T]
	mutex      sync.RWMutex
}

// Lock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (s *SyncAny[T]) Lock() {
	s.mutex.Lock()
}

// Unlock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (s *SyncAny[T]) Unlock() {
	s.mutex.Unlock()
}

// RLock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (s *SyncAny[T]) RLock() {
	s.mutex.RLock()
}

// RUnlock locks the mutex for writing.
// It can be used if some outside synchronization is needed.
func (s *SyncAny[T]) RUnlock() {
	s.mutex.RUnlock()
}

// NodeFirst returns the first node in the list.
// Returned node is not safe for concurrent use.
// If you would like to use plain Node, use RLock/RUnlock from the structure.
// See the details in the specific implementation.
func (s *SyncAny[T]) NodeFirst() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeFirst()
}

// NodeLast returns the last node in the list.
// Returned node is not safe for concurrent use.
// If you would like to use plain Node, use RLock/RUnlock from the structure.
// See the details in the specific implementation.
func (s *SyncAny[T]) NodeLast() Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeLast()
}

// NodeAt returns the node at the given index.
// Returned node is not safe for concurrent use.
// If you would like to use plain Node, use RLock/RUnlock from the structure.
// See the details in the specific implementation.
func (s *SyncAny[T]) NodeAt(index int) Node[T] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.NodeAt(index)
}

// ValueFirst returns the first value in the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) ValueFirst() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueFirst()
}

// ValueLast returns the last value in the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) ValueLast() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueLast()
}

// ValueAt returns the value at the given index.
// See the details in the specific implementation.
func (s *SyncAny[T]) ValueAt(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.ValueAt(index)
}

// IndexOf returns the first index of the given value.
// See the details in the specific implementation.
func (s *SyncAny[T]) IndexOf(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOf(value)
}

// IndexOfLast returns the last index of the given value.
// See the details in the specific implementation.
func (s *SyncAny[T]) IndexOfLast(value T) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.IndexOfLast(value)
}

// Contains returns true if the given value is in the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) Contains(value T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Contains(value)
}

// Length returns the number of elements in the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) Length() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Length()
}

// Empty returns true if the list is empty.
// See the details in the specific implementation.
func (s *SyncAny[T]) Empty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.linkedlist.Empty()
}

// AddFirst adds the given value to the front of the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) AddFirst(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddFirst(value)
}

// AddLast adds the given value to the end of the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) AddLast(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddLast(value)
}

// AddAt adds the given value at the given index.
// See the details in the specific implementation.
func (s *SyncAny[T]) AddAt(index int, value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.AddAt(index, value)
}

// PopFirst removes and returns the first value in the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) PopFirst() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopFirst()
}

// PopLast removes and returns the last value in the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) PopLast() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopLast()
}

// PopAt removes and returns the value at the given index.
// See the details in the specific implementation.
func (s *SyncAny[T]) PopAt(index int) (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.PopAt(index)
}

// RemoveNode removes the given node from the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) RemoveNode(node Node[T]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveNode(node)
}

// RemoveFirstBy removes the first value that equals the given value.
// See the details in the specific implementation.
func (s *SyncAny[T]) RemoveFirstBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveFirstBy(value)
}

// RemoveLastBy removes the last value that equals the given value.
// See the details in the specific implementation.
func (s *SyncAny[T]) RemoveLastBy(value T) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveLastBy(value)
}

// RemoveAllBy removes all values that equals the given value.
// See the details in the specific implementation.
func (s *SyncAny[T]) RemoveAllBy(value T) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.linkedlist.RemoveAllBy(value)
}

// Clear removes all nodes from the list.
// See the details in the specific implementation.
func (s *SyncAny[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.linkedlist.Clear()
}
