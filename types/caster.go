package types

// Cast converts a value to a specified type.
// It returns default type's value, if input is nil.
func Cast[T any](value any) (val T, ok bool) {
	if value == nil {
		return val, false
	}

	val, ok = value.(T)

	return val, ok
}

// CastOk converts a value to a specified type with checking second ok parameter.
// If the second parameter is false, the function returns the default value and false.
// It returns default type's value, if input is nil.
// This function is needed for this use case:
//
//	func main() {
//	  f1 := func() (any, bool) {
//	    return 1, true
//	  }
//
//	  val, ok := CastOk[int](f1())
//	  fmt.Println(val, ok)
//	}
//	Output: (1, true)
func CastOk[T any](value any, ok bool) (T, bool) {
	if !ok {
		var val T

		return val, false
	}

	return Cast[T](value)
}

// Ref returns a pointer to the value passed in.
// Useful for getting a pointer to any value.
func Ref[T any](value T) *T {
	return &value
}
