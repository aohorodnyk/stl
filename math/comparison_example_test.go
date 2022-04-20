package math_test

import (
	"fmt"
	mathstd "math"

	"github.com/aohorodnyk/stl/math"
)

func ExampleMin() {
	fmt.Println(math.Min(1, 2))
	fmt.Println(math.Min(5, -242234324))
	fmt.Println(math.Min(byte(5), 15))
	fmt.Println(math.Min(uint(123), 267865765426754))
	fmt.Println(math.Min(46524, uint(124)))
	fmt.Println(math.Min(635235.25362, 0.000001))

	// Output:
	// 1
	// -242234324
	// 5
	// 123
	// 124
	// 1e-06
}

func ExampleMax() {
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Max(5, -242234324))
	fmt.Println(math.Max(byte(5), 15))
	fmt.Println(math.Max(uint(123), 267865765426754))
	fmt.Println(math.Max(46524, uint(124)))
	fmt.Println(math.Max(635235.25362, 0.000001))

	// Output:
	// 2
	// 5
	// 15
	// 267865765426754
	// 46524
	// 635235.25362
}

func ExampleMinMulti() {
	fmt.Println(math.MinMulti(15, 345, 12, -345, -53))
	fmt.Println(math.MinMulti(byte(5), 15, 3, 56))
	fmt.Println(math.MinMulti(uint(123), 267865765426754, 1, 346))
	fmt.Println(math.MinMulti(46524, uint(124)))
	fmt.Println(math.MinMulti(635235.25362, 0.000001, -324235.1242, -234., 345.0, 0.))

	ints := []int64{mathstd.MaxInt64, mathstd.MinInt64, 345346, 125, -346763525, 346456346}
	fmt.Println(math.MinMulti(ints...))

	ints = nil
	fmt.Println(math.MinMulti(ints...))

	// Output:
	// -345
	// 3
	// 1
	// 124
	// -324235.1242
	// -9223372036854775808
	// 0
}

func ExampleMaxMulti() {
	fmt.Println(math.MaxMulti(15, 353, 12, -345, -53))
	fmt.Println(math.MaxMulti(byte(5), 15, 3, 56))
	fmt.Println(math.MaxMulti(uint(123), 267865765426754, 1, 346))
	fmt.Println(math.MaxMulti(46524, int32(124), -235))
	fmt.Println(math.MaxMulti(635235.25362, 0.000001, -324235.1242, -234., 345.0, 0.))

	ints := []int64{mathstd.MaxInt64, mathstd.MinInt64, 345346, 125, -346763525, 346456346}
	fmt.Println(math.MinMulti(ints...))

	ints = nil
	fmt.Println(math.MinMulti(ints...))

	// Output:
	// 353
	// 56
	// 267865765426754
	// 46524
	// 635235.25362
	// -9223372036854775808
	// 0
}

func ExampleCompareMulti() {
	// Find the maximum value of the following set of numbers.
	cmpInt := func(a, b int) bool { return a > b }
	fmt.Println(math.CompareMulti(cmpInt, 15, 353, 12, -345, -53))

	// Find the maximum value of the following set of numbers.
	cmpInt = func(a, b int) bool { return a < b }
	fmt.Println(math.CompareMulti(cmpInt, 15, 353, 12, -345, -53))

	// Specify the custm type for further examples.
	type customType struct {
		field int
	}

	// Find the minimum value for the custom structure type.
	cmpCustom := func(first, second customType) bool { return first.field < second.field }
	fmt.Println(math.CompareMulti(cmpCustom, customType{field: 15}, customType{field: 353}, customType{field: 12}))

	// Find the maximum value for the custom structure type.
	cmpCustom = func(first, second customType) bool { return first.field > second.field }
	fmt.Println(math.CompareMulti(cmpCustom, customType{field: 15}, customType{field: 353}, customType{field: 12}))

	// Compare empty list of values for the custom structure type.
	fmt.Println(math.CompareMulti(cmpCustom))

	cmpRef := func(first, second *customType) bool { return first.field > second.field }
	// Compare empty list of values for the reference of the custom structure type.
	fmt.Println(math.CompareMulti(cmpRef))

	// Compare the one value for the reference of the custom structure type.
	fmt.Println(math.CompareMulti(cmpRef, &customType{field: 1}))
	// Compare the list of values for the reference of the custom structure type.
	fmt.Println(math.CompareMulti(cmpRef, &customType{field: 1}, &customType{field: 2}, &customType{field: -15}, &customType{field: -2}))

	// Output:
	// 353
	// -345
	// {12}
	// {353}
	// {0}
	// <nil>
	// &{1}
	// &{2}
}
