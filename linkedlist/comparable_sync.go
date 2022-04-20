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

func (c *ComparableSync[T]) NodeFirst() Node[T] {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.NodeFirst()
}

func (c *ComparableSync[T]) NodeLast() Node[T] {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.NodeLast()
}

func (c *ComparableSync[T]) NodeAt(index int) Node[T] {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.NodeAt(index)
}

func (c *ComparableSync[T]) ValueFirst() (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.ValueFirst()
}

func (c *ComparableSync[T]) ValueLast() (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.ValueLast()
}

func (c *ComparableSync[T]) ValueAt(index int) (T, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.ValueAt(index)
}

func (c *ComparableSync[T]) IndexOf(value T) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.IndexOf(value)
}

func (c *ComparableSync[T]) IndexOfLast(value T) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.IndexOfLast(value)
}

func (c *ComparableSync[T]) Contains(value T) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.Contains(value)
}

func (c *ComparableSync[T]) Length() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.Length()
}

func (c *ComparableSync[T]) Empty() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.linkedlist.Empty()
}

func (c *ComparableSync[T]) AddFirst(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.AddFirst(value)
}

func (c *ComparableSync[T]) AddLast(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.AddLast(value)
}

func (c *ComparableSync[T]) AddAt(index int, value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.AddAt(index, value)
}

func (c *ComparableSync[T]) PopFirst() (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.PopFirst()
}

func (c *ComparableSync[T]) PopLast() (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.PopLast()
}

func (c *ComparableSync[T]) PopAt(index int) (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.PopAt(index)
}

func (c *ComparableSync[T]) RemoveNode(node Node[T]) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveNode(node)
}

func (c *ComparableSync[T]) RemoveFirstBy(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveFirstBy(value)
}

func (c *ComparableSync[T]) RemoveLastBy(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveLastBy(value)
}

func (c *ComparableSync[T]) RemoveAllBy(value T) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.linkedlist.RemoveAllBy(value)
}

func (c *ComparableSync[T]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.linkedlist.Clear()
}
