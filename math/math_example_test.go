package math_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/math"
)

func ExampleAbs() {
	fmt.Println(math.Abs(int8(math.MinInt8))) // This is the exception, we cannot convert it to the positive number.
	fmt.Println(math.Abs(int8(math.MinInt8 + 1)))
	fmt.Println(math.Abs(int8(127)))

	fmt.Println(math.Abs(int(math.MinInt))) // This is the exception, we cannot convert it to the positive number.
	fmt.Println(math.Abs(int(math.MinInt + 1)))
	fmt.Println(math.Abs(int(-5)))
	fmt.Println(math.Abs(int(0)))
	fmt.Println(math.Abs(int(7)))

	fmt.Println(math.Abs(float32(math.SmallestNonzeroFloat32)))
	fmt.Println(math.Abs(float32(math.MaxFloat32)))
	fmt.Println(math.Abs(float64(math.SmallestNonzeroFloat64)))
	fmt.Println(math.Abs(float64(math.MaxFloat64)))

	// Output:
	// -128
	// 127
	// 127
	// -9223372036854775808
	// 9223372036854775807
	// 5
	// 0
	// 7
	// 1e-45
	// 3.4028235e+38
	// 5e-324
	// 1.7976931348623157e+308
}
