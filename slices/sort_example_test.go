package slices_test

import (
	"fmt"
	"sort"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleSort() {
	// Sort strings in ascending order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	slices.Sort(strs)
	fmt.Println(strs)

	// Sort integers in ascending order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	slices.Sort(ints)
	fmt.Println(ints)

	// Sort floats in ascending order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	slices.Sort(floats)
	fmt.Println(floats)

	// Output:
	// [a aa b bb c cc d]
	// [1 2 3 4 5 6 7 8 9 12]
	// [1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 12.12]
}

func ExampleSortReverse() {
	// Sort strings in reverse order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	slices.SortReverse(strs)
	fmt.Println(strs)

	// Sort integers in reverse order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	slices.SortReverse(ints)
	fmt.Println(ints)

	// Sort floats in reverse order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	slices.SortReverse(floats)
	fmt.Println(floats)

	// Output:
	// [d cc c bb b aa a]
	// [12 9 8 7 6 5 4 3 2 1]
	// [12.12 9.9 8.8 7.7 6.6 5.5 4.4 3.3 2.2 1.1]
}

func ExampleStable() {
	strs := []string{"c", "a", "d", "b"}
	slices.Stable(strs)
	fmt.Println(strs)

	ints := []int{5, 4, 3, 2}
	slices.Stable(ints)
	fmt.Println(ints)

	floats := []float64{2.2, 3.3, 4.1, 4.2}
	slices.Stable(floats)
	fmt.Println(floats)

	// Output:
	// [a b c d]
	// [2 3 4 5]
	// [2.2 3.3 4.1 4.2]
}

func ExampleStableReverse() {
	// Sort strings in reverse order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	slices.StableReverse(strs)
	fmt.Println(strs)

	// Sort integers in reverse order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	slices.StableReverse(ints)
	fmt.Println(ints)

	// Sort floats in reverse order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	slices.StableReverse(floats)
	fmt.Println(floats)

	// Output:
	// [d cc c bb b aa a]
	// [12 9 8 7 6 5 4 3 2 1]
	// [12.12 9.9 8.8 7.7 6.6 5.5 4.4 3.3 2.2 1.1]
}

func ExampleSliceOrdered() {
	// There are examples of using SliceOrdered in standard sort package.
	// Sort strings in ascending order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.Sort(slices.SliceOrdered[string](strs))
	fmt.Println(strs)

	// Sotr strings in reverse order.
	strs = []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.Sort(sort.Reverse(slices.SliceOrdered[string](strs)))
	fmt.Println(strs)

	// Output:
	// [a aa b bb c cc d]
	// [d cc c bb b aa a]
}
