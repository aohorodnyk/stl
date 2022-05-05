package maps

// Equal compares two maps with comparable values.
// It returns true if the maps are equal, false otherwise.
// This is linear comparing function with O(n) time complexity and O(1) space complexity.
func Equal[K comparable, V comparable](left, right map[K]V) bool {
	cmp := func(key K) bool {
		vl := left[key]
		vr, okr := right[key]

		if !okr {
			return false
		}

		return vl == vr
	}

	return EqualFunc(left, right, cmp)
}

// EqualFunc compares two maps with any values by a custom comparable function.
// It returns true if the maps are equal, false otherwise.
// This is linear comparing function with O(n) time complexity and O(1) space complexity.
func EqualFunc[K comparable, T any](left, right map[K]T, cmp func(K) bool) bool {
	if len(left) != len(right) {
		return false
	}

	for key := range left {
		if !cmp(key) {
			return false
		}
	}

	return true
}
