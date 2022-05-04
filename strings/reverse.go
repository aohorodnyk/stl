package strings

import "github.com/aohorodnyk/stl/slices"

// Reverse reverses a string.
func Reverse(str string) string {
	runes := []rune(str)
	slices.Reverse(runes)

	return string(runes)
}
