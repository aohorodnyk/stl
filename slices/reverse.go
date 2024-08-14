package slices

// Reverse function reverses IN PLACE of any type of slice.
// Performance: O(n/2) -> O(n).
// Memory: O(1).
//
// This function is safe to use with null and empty slices.
func Reverse[T any](slice []T) {
	size := len(slice)
	for index := range size / 2 {
		slice[index], slice[size-1-index] = slice[size-1-index], slice[index]
	}
}
