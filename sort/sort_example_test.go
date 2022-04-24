package sort_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/sort"
)

func ExampleOrdered() {
	// Sort strings in ascending order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.Ordered(strs)
	fmt.Println(strs)

	// Sort integers in ascending order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	sort.Ordered(ints)
	fmt.Println(ints)

	// Sort floats in ascending order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	sort.Ordered(floats)
	fmt.Println(floats)

	// Output:
	// [a aa b bb c cc d]
	// [1 2 3 4 5 6 7 8 9 12]
	// [1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 12.12]
}

func ExampleOrderedReverse() {
	// Sort strings in reverse order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.OrderedReverse(strs)
	fmt.Println(strs)

	// Sort integers in reverse order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	sort.OrderedReverse(ints)
	fmt.Println(ints)

	// Sort floats in reverse order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	sort.OrderedReverse(floats)
	fmt.Println(floats)

	// Output:
	// [d cc c bb b aa a]
	// [12 9 8 7 6 5 4 3 2 1]
	// [12.12 9.9 8.8 7.7 6.6 5.5 4.4 3.3 2.2 1.1]
}

func ExampleLessFunc() {
	// Sort strings in ascending order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.LessFunc(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	fmt.Println(strs)

	// Sort integers in ascending order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	sort.LessFunc(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})
	fmt.Println(ints)

	// Sort floats in ascending order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	sort.LessFunc(floats, func(i, j int) bool {
		return floats[i] < floats[j]
	})
	fmt.Println(floats)

	type customType struct {
		value int
	}

	// Sort custom types in ascending order.
	customs := []customType{
		{value: 5},
		{value: 3},
		{value: 4},
		{value: 2},
	}
	sort.LessFunc(customs, func(i, j int) bool {
		return customs[i].value < customs[j].value
	})
	fmt.Println(customs)

	// Output:
	// [a aa b bb c cc d]
	// [1 2 3 4 5 6 7 8 9 12]
	// [1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 12.12]
	// [{2} {3} {4} {5}]
}
