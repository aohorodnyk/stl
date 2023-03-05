package slices

// Equal returns true if the two slices are equal, false otherwise.
// This is linear comparing function with O(n) time complexity and O(1) space complexity.
func Equal[T comparable](left, right []T) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

// EqualFunc returns true if the two slices are equal, false otherwise.
// This is linear comparing function with O(n) time complexity and O(1) space complexity.
func EqualFunc[T any](left, right []T, cmp func(int) bool) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if !cmp(i) {
			return false
		}
	}

	return true
}
