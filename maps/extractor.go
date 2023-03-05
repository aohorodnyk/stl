package maps

// Values extracts values from a map and returns it as a slice.
// The values order will be the same as we received by iteration by the map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}

	return values
}

// Keys extracts keys from a map and returns it as a slice.
// The key order will be the same as we received by iteration by the map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
