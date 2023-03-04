package slices_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleReverse() {
	fmt.Println("Nil value, instead of slice")

	var sliceNil []int

	slices.Reverse(sliceNil)
	fmt.Println(sliceNil)
	fmt.Println()

	// =-=-=-=-=-=-=

	fmt.Println("Empty slice")

	sliceEmpty := []int{}

	slices.Reverse(sliceEmpty)
	fmt.Println(sliceEmpty)
	fmt.Println()

	// =-=-=-=-=-=-=

	fmt.Println("One value in a slice")

	sliceOne := []int{1}

	slices.Reverse(sliceOne)
	fmt.Println(sliceOne)
	fmt.Println()

	// =-=-=-=-=-=-=

	fmt.Println("Two values in a slice")

	sliceTwo := []int{1, 2}

	fmt.Println("Before:", sliceTwo)
	slices.Reverse(sliceTwo)
	fmt.Println("After:", sliceTwo)
	fmt.Println()

	// =-=-=-=-=-=-=

	fmt.Println("Three values in a slice")

	sliceThree := []int{1, 2, 3}

	fmt.Println("Before:", sliceThree)
	slices.Reverse(sliceThree)
	fmt.Println("After:", sliceThree)
	fmt.Println()

	// Output:
	// Nil value, instead of slice
	// []
	//
	// Empty slice
	// []
	//
	// One value in a slice
	// [1]
	//
	// Two values in a slice
	// Before: [1 2]
	// After: [2 1]
	//
	// Three values in a slice
	// Before: [1 2 3]
	// After: [3 2 1]
}
