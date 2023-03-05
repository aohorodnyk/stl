package slices_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleIndex() {
	// Search in slice of ints.
	sliceInt := []int{1, 2, 3, 4, 5}
	targetInt := 3

	fmt.Println(slices.Index(sliceInt, targetInt))

	// Search in slice of strings.
	sliceStr := []string{"a", "b", "d", "c", "e"}
	targetStr := "c"

	fmt.Println(slices.Index(sliceStr, targetStr))

	// Output:
	// 2
	// 3
}

func ExampleIndexFunc() {
	// Search in slice of ints.
	sliceInt := []int{1, 2, 3, 4, 5}
	targetInt := 3
	cmp := func(i int) bool {
		return sliceInt[i] == targetInt
	}

	fmt.Println(slices.IndexFunc(sliceInt, cmp))

	// Search in slice of strings.
	sliceStr := []string{"a", "b", "d", "c", "e"}
	targetStr := "c"
	cmp = func(i int) bool {
		return sliceStr[i] == targetStr
	}

	fmt.Println(slices.IndexFunc(sliceStr, cmp))

	// Search in slice of custom structure types.
	type customType struct {
		val int
	}

	sliceCustom := []customType{
		{1}, {1}, {2}, {6}, {3}, {4}, {5},
	}
	targetCustom := customType{4}
	cmp = func(i int) bool {
		return sliceCustom[i].val == targetCustom.val
	}

	fmt.Println(slices.IndexFunc(sliceCustom, cmp))

	// Output:
	// 2
	// 3
	// 5
}
