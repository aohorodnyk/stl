package lru

import (
	"github.com/aohorodnyk/stl/collections/linkedlist"
)

// NewSimple returns new instance of SimpleLRU cache.
func NewSimple[Key comparable, Value any](capacity int) *Simple[Key, Value] {
	if capacity <= 0 {
		panic("capacity must be greater than zero")
	}

	return &Simple[Key, Value]{
		evictList: linkedlist.NewDoublyAnyDeep[entity[Key, Value]](),
		data:      make(map[Key]linkedlist.Node[entity[Key, Value]], capacity),
		capacity:  capacity,
	}
}

// Simple is a simple LRU cache.
// The simple cache is a cache that stores the most recently used items.
// The most frequently used item will be evicted when the cache is full.
// It is not safe for concurrent use.
// This implementation has O(1) performance complexity for all operations.
type Simple[Key comparable, Value any] struct {
	evictList linkedlist.LinkedList[entity[Key, Value]]
	data      map[Key]linkedlist.Node[entity[Key, Value]]
	length    int
	capacity  int
}

// Length returns the number of items in the cache.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Length() int {
	return lru.length
}

// Contains returns true if the cache contains the given key.
// This method does not update the access time of the item.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Contains(key Key) bool {
	_, ok := lru.data[key]

	return ok
}

// Peek returns the value for the given key without updating the access time.
// If the key is not in the cache, Peek returns false in the seconds return value.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Peek(key Key) (Value, bool) {
	node, ok := lru.data[key]
	if !ok {
		var tmp Value

		return tmp, false
	}

	return node.Value().value, true
}

// Value returns the value associated with the key.
// If the key does not exist, the second return value will be false.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Value(key Key) (Value, bool) {
	node, ok := lru.data[key]
	if !ok {
		var tmp Value

		return tmp, false
	}

	lru.evictList.RemoveNode(node)
	lru.evictList.AddFirst(node.Value())

	lru.data[key] = lru.evictList.NodeFirst()

	return lru.data[key].Value().value, true
}

// Set sets the value of the key.
// If the key exists, it overrides the previous value.
// If the key does not exist, it is added to the cache.
// If the cache is full, the least recently used item is evicted.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Set(key Key, value Value) {
	entityNew := entity[Key, Value]{
		key:   key,
		value: value,
	}

	node, ok := lru.data[key]
	if ok {
		lru.evictList.RemoveNode(node)
		lru.evictList.AddFirst(entityNew)

		lru.data[key] = lru.evictList.NodeFirst()

		return
	}

	if lru.length == lru.capacity {
		if entityOld, ok := lru.evictList.PopLast(); ok {
			delete(lru.data, entityOld.key)

			lru.length--
		}
	}

	lru.evictList.AddFirst(entityNew)
	lru.data[key] = lru.evictList.NodeFirst()
	lru.length++
}

// Pop returns value for the given key.
// If the key is not found, it returns false in the second return value.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Pop(key Key) (Value, bool) {
	node, ok := lru.data[key]
	if !ok {
		var tmp Value

		return tmp, false
	}

	lru.evictList.RemoveNode(node)
	delete(lru.data, key)

	lru.length--

	return node.Value().value, true
}

// Remove removes the value associated with the key.
// It returns true if the value was removed, false otherwise.
// This method has O(1) performance complexity.
func (lru *Simple[Key, Value]) Remove(key Key) {
	node, ok := lru.data[key]
	if !ok {
		return
	}

	lru.evictList.RemoveNode(node)
	delete(lru.data, key)

	lru.length--
}

// Clear removes all items from the cache.
// This method has O(1) performance complexity.
// But it allocates memory for the new map with required preallocated capacity.
func (lru *Simple[Key, Value]) Clear() {
	lru.data = make(map[Key]linkedlist.Node[entity[Key, Value]], lru.capacity)
	lru.length = 0

	lru.evictList.Clear()
}
