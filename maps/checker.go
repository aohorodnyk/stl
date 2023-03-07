package maps

// KeyExists checks is key exists in a map.
// This function does not compare values.
// To delete key from a map, `delete` function must be used.
func KeyExists[Key comparable, Value any](container map[Key]Value, key Key) bool {
	_, ok := container[key]

	return ok
}

// Search searches value in a provided map.
// Performance complexity is O(n), memory complaxity is O(1).
// It supports ONLY comparable value, otherwise we cannot compare values without reflect.
func Search[Key comparable, Value comparable](container map[Key]Value, value Value) (Key, bool) {
	for key, val := range container {
		if val == value {
			return key, true
		}
	}

	var key Key

	return key, false
}

// Search searches value in a provided map.
// Performance complexity is O(n), memory complaxity is O(1).
func SearchFunc[Key comparable, Value any](container map[Key]Value, cmp func(Key) bool) (Key, bool) {
	for key := range container {
		if cmp(key) {
			return key, true
		}
	}

	var key Key

	return key, false
}
