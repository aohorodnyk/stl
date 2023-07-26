package maps

// Set sets value to a provided key in a provided container.
// It returns old value and did container contain old value before set.
// If there is not needs to know about previous value, just do `container[key] = value`, no needs to use this function.
func Set[K comparable, V any](container map[K]V, key K, value V) (old V, ok bool) {
	old, ok = container[key]
	container[key] = value

	return old, ok
}

// Delete deletes keys from a provided map.
// This function is just a wrapper around `delete` function, but with a support of variadic keys.
func Delete[K comparable, V any](container map[K]V, keys ...K) {
	for _, key := range keys {
		delete(container, key)
	}
}

// DeleteBy deletes keys from a map by a predicate.
// If predicate returns true, key will be deleted.
func DeleteBy[K comparable, V any](container map[K]V, predicate func(key K, value V) bool) {
	for key, value := range container {
		if predicate(key, value) {
			delete(container, key)
		}
	}
}
