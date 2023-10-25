package types_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/types"
)

func ExampleValue() {
	returnErr := func() (string, error) {
		return "returnErr", nil
	}

	returnBool := func() (string, bool) {
		return "returnBool", false
	}

	returnErrAndBool := func() (string, error, bool) {
		return "returnErrAndBool", nil, false
	}

	// This is how you can use Value to ignore the second parameter.
	// It's useful for the cases when you need to pass a value
	// to a function that requires a parameter of a specific type, but another function returns
	// two parameters like: (value, error), (value, bool).
	fmt.Println(types.Value(returnErr()))
	fmt.Println(types.Value(returnBool()))
	fmt.Println(types.Value(returnErrAndBool()))

	// Output:
	// returnErr
	// returnBool
	// returnErrAndBool
}
