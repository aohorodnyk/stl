package slices

// Index returns the index of the first element in the slice that equals the given target.
// If no such element is found, -1 is returned.
// This is linear search with O(n) time complexity and O(1) space complaxity.
func Index[T comparable](slice []T, target T) int {
	for i := range slice {
		if slice[i] == target {
			return i
		}
	}

	return -1
}

// IndexFunc returns the index of the first element in the slice that satisfies the given predicate.
// If no such element is found, -1 is returned.
// This is linear search with O(n) time complexity and O(1) space complaxity.
func IndexFunc[T any](slice []T, cmp func(int) bool) int {
	for i := range slice {
		if cmp(i) {
			return i
		}
	}

	return -1
}
