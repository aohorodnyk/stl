package types

import "fmt"

func ExampleIsType() {
	fmt.Println(IsType[int](0))
	fmt.Println(IsType[int](1))
	fmt.Println(IsType[int](1.0))
	fmt.Println(IsType[int]("1"))
	fmt.Println(IsType[string](0))
	fmt.Println(IsType[string](1))
	fmt.Println(IsType[string]("1"))
	fmt.Println(IsType[*string]("1"))
	fmt.Println(IsType[*string](Ref("1")))

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
