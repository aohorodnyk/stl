package types

// Value returns the given value by ignoring the second parameter.
// This function is useful for the cases when you need to pass a value
// to a function that requires a parameter of a specific type, but another function returns
// two parameters like: (value, error), (value, bool).
func Value[T any](value T, _ ...any) T {
	return value
}
