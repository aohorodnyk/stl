package cache

import "sync"

// NewSync returns new instance of Sync cache.
// The Sync container is safe for concurrent use.
// It covers the implementation from cache parameter.
// See the details in the specific implementation of cache from a marameter.
func NewSync[Key comparable, Value any](cache Cache[Key, Value]) *Sync[Key, Value] {
	return &Sync[Key, Value]{
		cache: cache,
	}
}

// Sync is a generic container that emplements the Cache interface with concurrent safe operations.
// To cover the any implementation of the Cache interface, use `NewSync(cache)` function.
type Sync[Key comparable, Value any] struct {
	cache Cache[Key, Value]
	mu    sync.Mutex
}

// Contains returns true if the key is in the cache.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Contains(key Key) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.cache.Contains(key)
}

// Peek returns the value for the given key without updating the access time.
// If the key is not in the cache, Peek returns false in the seconds return value.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Peek(key Key) (value Value, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.cache.Peek(key)
}

// Value returns the value associated with the key.
// If the key does not exist, the second return value will be false.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Value(key Key) (value Value, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.cache.Value(key)
}

// Set sets the value of the key.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Set(key Key, value Value) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache.Set(key, value)
}

// Pop returns value for the given key.
// If the key is not found, it returns false in the second return value.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Pop(key Key) (value Value, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.cache.Pop(key)
}

// Remove removes the value associated with the key.
// It returns true if the value was removed, false otherwise.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Remove(key Key) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache.Remove(key)
}

// Length returns the number of items in the cache.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Length() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.cache.Length()
}

// Clear removes all items from the cache.
// This method has O(1) performance complexity.
// See the details in the specific implementation.
func (s *Sync[Key, Value]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache.Clear()
}
