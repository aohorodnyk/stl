package math_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/math"
)

func ExampleMin() {
	fmt.Println("Min with one parameter")
	fmt.Println(math.Min(2512))
	fmt.Println()

	fmt.Println("Min with two parameters")
	fmt.Println(math.Min(2512, 212))
	fmt.Println()

	fmt.Println("Min with three parameters")
	fmt.Println(math.Min(2512, 21, 212))
	fmt.Println()

	fmt.Println("Min with five parameters")
	fmt.Println(math.Min(55, 2512, -234, 212, 346234))
	fmt.Println()

	fmt.Println("Min from slice")
	fmt.Println(math.Min([]int{83, 423, 457, 4573, 234}...))
	fmt.Println()

	fmt.Println("Min from slice of int16")
	fmt.Println(math.Min([]int16{83, 423, 457, 4573, 234}...))
	fmt.Println()

	// Output:
	// Min with one parameter
	// 2512
	//
	// Min with two parameters
	// 212
	//
	// Min with three parameters
	// 21
	//
	// Min with five parameters
	// -234
	//
	// Min from slice
	// 83
	//
	// Min from slice of int16
	// 83
}

func ExampleMax() {
	fmt.Println("Max with one parameter")
	fmt.Println(math.Max(2512))
	fmt.Println()

	fmt.Println("Max with two parameters")
	fmt.Println(math.Max(32, 212))
	fmt.Println()

	fmt.Println("Max with three parameters")
	fmt.Println(math.Max(1, 21, 20))
	fmt.Println()

	fmt.Println("Max with five parameters")
	fmt.Println(math.Max(55, 2512, -234, 212, 346234))
	fmt.Println()

	fmt.Println("Max from slice")
	fmt.Println(math.Max([]int{83, 423, 457, 4573, 234}...))
	fmt.Println()

	// Output:
	// Max with one parameter
	// 2512
	//
	// Max with two parameters
	// 212
	//
	// Max with three parameters
	// 21
	//
	// Max with five parameters
	// 346234
	//
	// Max from slice
	// 4573
}
