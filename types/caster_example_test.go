package types_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/types"
)

func ExampleCast() {
	fmt.Println(types.Cast[int](-242))
	fmt.Println(types.Cast[int](0))
	fmt.Println(types.Cast[int](3))
	fmt.Println(types.Cast[int](0x234))
	fmt.Println(types.Cast[int](0b101011))

	ir := new(int)
	*ir = 2532
	ir, ok := types.Cast[*int](ir)
	fmt.Printf("%T %d %#v\n", ir, *ir, ok)

	fmt.Println(types.Cast[string](""))
	fmt.Println(types.Cast[string]("test"))
	fmt.Println(types.Cast[string](nil))

	sr := new(string)
	*sr = "sr_test"
	sr, ok = types.Cast[*string](sr)
	fmt.Printf("%T %s %#v\n", sr, *sr, ok)

	sr, ok = types.Cast[*string](nil)
	fmt.Printf("%T %#v\n", sr, ok)

	// Output:
	// -242 true
	// 0 true
	// 3 true
	// 564 true
	// 43 true
	// *int 2532 true
	//  true
	// test true
	//  false
	// *string sr_test true
	// *string false
}

func ExampleCastOk() {
	fFalse := func() (any, bool) {
		return 234, false
	}

	fmt.Println(types.CastOk[string](fFalse()))
	fmt.Println(types.CastOk[int](fFalse()))
	fmt.Println(types.CastOk[float64](fFalse()))
	fmt.Println(types.CastOk[*string](fFalse()))
	fmt.Println(types.CastOk[*int](fFalse()))

	fTrue := func() (any, bool) {
		return 234, true
	}

	fmt.Println(types.CastOk[string](fTrue()))
	fmt.Println(types.CastOk[int](fTrue()))
	fmt.Println(types.CastOk[float64](fTrue()))
	fmt.Println(types.CastOk[*string](fTrue()))
	fmt.Println(types.CastOk[*int](fTrue()))

	// Output:
	// false
	// 0 false
	// 0 false
	// <nil> false
	// <nil> false
	//  false
	// 234 true
	// 0 false
	// <nil> false
	// <nil> false
}

func ExampleRef() {
	intVal := types.Ref[int](0)
	fmt.Println(*intVal)

	intVal = types.Ref[int](235)
	fmt.Println(*intVal)

	ir := new(int)
	*ir = 2532
	fmt.Printf("%T %d\n", types.Ref[*int](ir), *ir)

	strVal := types.Ref[string]("")
	fmt.Println(*strVal)

	strVal = types.Ref[string]("test")
	fmt.Println(*strVal)

	sr := new(string)
	*sr = "sr_test"
	fmt.Printf("%T %s\n", types.Ref[*string](sr), *sr)

	fmt.Printf("%T\n", types.Ref[*string](nil))

	// Output:
	// 0
	// 235
	// **int 2532
	//
	// test
	// **string sr_test
	// **string
}
