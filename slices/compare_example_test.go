package slices_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleEqual() {
	leftInt := []int{1, 2, 3}
	rightInt := []int{1, 2, 3}
	fmt.Println(slices.Equal(leftInt, rightInt))

	leftString := []string{"a", "b", "c"}
	rightString := []string{"a", "c", "c"}
	fmt.Println(slices.Equal(leftString, rightString))

	leftString = []string{"a", "b", "c"}
	rightString = []string{"a", "b", "c"}
	fmt.Println(slices.Equal(leftString, rightString))

	leftString = []string{"a", "b", "c"}
	rightString = []string{"a", "b", "c", "d"}
	fmt.Println(slices.Equal(leftString, rightString))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleEqualFunc() {
	// Compare equal int slices.
	leftInt := []int{1, 2, 3}
	rightInt := []int{1, 2, 3}
	cmp := func(i int) bool {
		return leftInt[i] == rightInt[i]
	}
	fmt.Println(slices.EqualFunc(leftInt, rightInt, cmp))

	// Compare mismatch int slices.
	leftInt = []int{1, 2, 3}
	rightInt = []int{1, 2, 4}
	cmp = func(i int) bool {
		return leftInt[i] == rightInt[i]
	}
	fmt.Println(slices.EqualFunc(leftInt, rightInt, cmp))

	// Compare slices with custom types
	type customType struct {
		a int
	}

	leftCustom := []customType{
		{a: 1}, {a: 2}, {a: 3}, {a: 4},
	}
	rightCustom := []customType{
		{a: 1}, {a: 2}, {a: 3}, {a: 4},
	}
	cmp = func(i int) bool {
		return leftCustom[i].a == rightCustom[i].a
	}
	fmt.Println(slices.EqualFunc(leftCustom, rightCustom, cmp))

	// Output:
	// true
	// false
	// true
}
