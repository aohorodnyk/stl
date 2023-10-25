package strings

import "github.com/aohorodnyk/stl/slices"

// Reverse reverses any string-inherited types.
func Reverse[S ~string](str S) S {
	runes := []rune(str)
	slices.Reverse(runes)

	return S(runes)
}
