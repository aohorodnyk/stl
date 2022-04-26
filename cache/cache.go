package cache

// Cache is an interface for a cache.
// It is used to store and retrieve data.
type Cache[Key comparable, Value any] interface {
	// Contains returns true if the key is in the cache.
	Contains(Key) bool
	// Value retrieves a value from the cache.
	// If the key is not found, the false returned in second return value.
	Value(Key) (Value, bool)
	// Set sets a value in the cache.
	Set(Key, Value)
	// Pop retrieves a value from the cache and removes it.
	// If the key is not found, the false returned in second return value.
	Pop(Key) (Value, bool)
	// Remove removes a value from the cache.
	Remove(Key)
	// Length returns the number of items in the cache.
	Length() int
	// Clear clears the cache.
	Clear()
}
