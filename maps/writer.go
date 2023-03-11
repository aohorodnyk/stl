package maps

// Set sets value to a provided key in a provided container.
// It returns old value and did container contain old value before set.
// If there is not needs to know about previous value, just do `container[key] = value`, no needs to use this function.
func Set[Key comparable, Value any](container map[Key]Value, key Key, value Value) (old Value, ok bool) {
	old, ok = container[key]
	container[key] = value

	return old, ok
}
