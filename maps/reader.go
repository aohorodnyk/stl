package maps

// KeyExists checks is key exists in a map.
// This function does not compare values.
// To delete key from a map, `delete` function must be used.
func KeyExists[K comparable, V any](container map[K]V, key K) bool {
	_, ok := container[key]

	return ok
}

// Search searches value in a provided map.
// Performance complexity is O(n), memory complexity is O(1).
// It supports ONLY comparable value, otherwise we cannot compare values without reflect.
func Search[K comparable, V comparable](container map[K]V, value V) (K, bool) {
	for key, val := range container {
		if val == value {
			return key, true
		}
	}

	var key K

	return key, false
}

// SearchFunc searches value in a provided map.
// Performance complexity is O(n), memory complexity is O(1).
func SearchFunc[K comparable, V any](container map[K]V, cmp func(K) bool) (K, bool) {
	for key := range container {
		if cmp(key) {
			return key, true
		}
	}

	var key K

	return key, false
}
