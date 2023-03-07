package maps

// Copy copies all key and values from source map to destination.
// If destination already contains they key, it will be overwritten.
// If destination map contains some unique key, it will not be removed.
// Returns false one when destination is nil, otherwise returns true.
// This functions does not do DeepCopy, it just assigns values.
func Copy[Key comparable, Value any](dest map[Key]Value, src map[Key]Value) bool {
	if dest == nil {
		return false
	}

	for key, val := range src {
		dest[key] = val
	}

	return true
}

// Clone create a new map with the same size as a source map and assign values
// from the source map.
// This functions does not do DeepCopy, it just assigns values.
func Clone[Key comparable, Value any](src map[Key]Value) map[Key]Value {
	res := make(map[Key]Value, len(src))
	Copy(res, src)

	return res
}
