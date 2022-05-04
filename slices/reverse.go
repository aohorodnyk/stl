package slices

// Reverse is a helper function that simplifies the API for standard library.
// Complexity: O(n) performance, O(1) space.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
