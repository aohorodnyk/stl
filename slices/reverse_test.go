package slices_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleReverse() {
	var slice []int

	slices.Reverse(slice)
	fmt.Println(slice)

	strs := []string{"c", "a", "d", "b"}
	slices.Reverse(strs)
	fmt.Println(strs)

	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	slices.Reverse(ints)
	fmt.Println(ints)

	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	slices.Reverse(floats)
	fmt.Println(floats)

	// Output:
	// []
	// [b d a c]
	// [6 9 12 8 1 7 2 4 3 5]
	// [6.6 9.9 12.12 8.8 1.1 7.7 2.2 4.4 3.3 5.5]
}
