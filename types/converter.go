package types

// Convert converts a value to a specified type.
func ConvertType[Type any](value any) (val Type, ok bool) {
	if value == nil {
		return val, false
	}

	val, ok = value.(Type)

	return val, ok
}

// ConvertOk converts a value to a specified type with checking second ok parameter.
// If the second parameter is false, the function returns the default value and false.
// This function is needed for this use case:
//   func main() {
//     f1 := func() (any, bool) {
//       return 1, true
//     }
//
//     val, ok := ConvertOk(f1())
//     fmt.Println(val, ok)
//   }
//   Output: (1, true)
func ConvertTypeOk[Type any](value any, ok bool) (Type, bool) {
	var val Type

	if value == nil {
		return val, false
	}

	val, okType := value.(Type)
	if ok {
		ok = okType
	}

	return val, ok
}
