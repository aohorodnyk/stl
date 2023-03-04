package math_test

import (
	"fmt"
	mathstd "math"

	"github.com/aohorodnyk/stl/math"
)

func ExampleAbs() {
	fmt.Println("Edge case, min int that cannot be converted to a positive value of the same type")
	fmt.Println(math.Abs(mathstd.MinInt64))
	fmt.Println()

	fmt.Println("The smallest int value that can be converted to the max positive value")
	fmt.Println(math.Abs(mathstd.MinInt64 + 1))
	fmt.Println()

	fmt.Println("Zero")
	fmt.Println(math.Abs(0))
	fmt.Println()

	fmt.Println("Positive value")
	fmt.Println(math.Abs(234))
	fmt.Println()

	fmt.Println("Minimum int8 value that cannot convert")
	fmt.Println(math.Abs(int8(mathstd.MinInt8)))
	fmt.Println()

	fmt.Println("Minimum int8 value that can be converted")
	fmt.Println(math.Abs(int8(mathstd.MinInt8 + 1)))
	fmt.Println()

	fmt.Println("Zero int8")
	fmt.Println(math.Abs(int8(0)))
	fmt.Println()

	fmt.Println("Max int8")
	fmt.Println(math.Abs(int8(127)))
	fmt.Println()

	// Output:
	// Edge case, min int that cannot be converted to a positive value of the same type
	// -9223372036854775808
	//
	// The smallest int value that can be converted to the max positive value
	// 9223372036854775807
	//
	// Zero
	// 0
	//
	// Positive value
	// 234
	//
	// Minimum int8 value that cannot convert
	// -128
	//
	// Minimum int8 value that can be converted
	// 127
	//
	// Zero int8
	// 0
	//
	// Max int8
	// 127
}
