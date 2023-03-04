package types

// Cast converts a value to a specified type.
// It returns default type's value, if input is nil.
func Cast[Type any](value any) (val Type, ok bool) {
	if value == nil {
		return val, false
	}

	val, ok = value.(Type)

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
func CastOk[Type any](value any, ok bool) (Type, bool) {
	if !ok {
		var val Type

		return val, false
	}

	return Cast[Type](value)
}
