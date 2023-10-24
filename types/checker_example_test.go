package types_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/types"
)

func ExampleIsType() {
	fmt.Println(types.IsType[int](0))
	fmt.Println(types.IsType[int](1))
	fmt.Println(types.IsType[int](1.0))
	fmt.Println(types.IsType[int]("1"))
	fmt.Println(types.IsType[string](0))
	fmt.Println(types.IsType[string](1))
	fmt.Println(types.IsType[string]("1"))
	fmt.Println(types.IsType[*string]("1"))
	fmt.Println(types.IsType[*string](types.Ref("1")))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
	// true
	// false
	// true
}
