package types_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/types"
)

func ExampleConvertType() {
	fmt.Println(types.ConvertType[int](-242))
	fmt.Println(types.ConvertType[int](0))
	fmt.Println(types.ConvertType[int](3))
	fmt.Println(types.ConvertType[int](0x234))
	fmt.Println(types.ConvertType[int](0b101011))

	ir := new(int)
	*ir = 2532
	ir, ok := types.ConvertType[*int](ir)
	fmt.Printf("%T %d %#v\n", ir, *ir, ok)

	fmt.Println(types.ConvertType[string](""))
	fmt.Println(types.ConvertType[string]("test"))
	fmt.Println(types.ConvertType[string](nil))

	sr := new(string)
	*sr = "sr_test"
	sr, ok = types.ConvertType[*string](sr)
	fmt.Printf("%T %s %#v\n", sr, *sr, ok)

	sr, ok = types.ConvertType[*string](nil)
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

func ExampleConvertTypeOk() {
	fFalse := func() (any, bool) {
		return 234, false
	}

	fmt.Println(types.ConvertTypeOk[string](fFalse()))
	fmt.Println(types.ConvertTypeOk[int](fFalse()))
	fmt.Println(types.ConvertTypeOk[float64](fFalse()))
	fmt.Println(types.ConvertTypeOk[*string](fFalse()))
	fmt.Println(types.ConvertTypeOk[*int](fFalse()))

	fTrue := func() (any, bool) {
		return 234, true
	}

	fmt.Println(types.ConvertTypeOk[string](fTrue()))
	fmt.Println(types.ConvertTypeOk[int](fTrue()))
	fmt.Println(types.ConvertTypeOk[float64](fTrue()))
	fmt.Println(types.ConvertTypeOk[*string](fTrue()))
	fmt.Println(types.ConvertTypeOk[*int](fTrue()))

	// Output:
	//  false
	// 234 false
	// 0 false
	// <nil> false
	// <nil> false
	//  false
	// 234 true
	// 0 false
	// <nil> false
	// <nil> false
}
