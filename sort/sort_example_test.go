package sort_test

import (
	"fmt"
	sortstd "sort"

	"github.com/aohorodnyk/stl/sort"
)

func ExampleSort() {
	// Sort strings in ascending order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.Sort(strs)
	fmt.Println(strs)

	// Sort integers in ascending order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	sort.Sort(ints)
	fmt.Println(ints)

	// Sort floats in ascending order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	sort.Sort(floats)
	fmt.Println(floats)

	// Output:
	// [a aa b bb c cc d]
	// [1 2 3 4 5 6 7 8 9 12]
	// [1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 12.12]
}

func ExampleReverse() {
	// Sort strings in reverse order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.Reverse(strs)
	fmt.Println(strs)

	// Sort integers in reverse order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	sort.Reverse(ints)
	fmt.Println(ints)

	// Sort floats in reverse order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	sort.Reverse(floats)
	fmt.Println(floats)

	// Output:
	// [d cc c bb b aa a]
	// [12 9 8 7 6 5 4 3 2 1]
	// [12.12 9.9 8.8 7.7 6.6 5.5 4.4 3.3 2.2 1.1]
}

func ExampleSorted() {
	strs := []string{"c", "a", "d", "b"}
	fmt.Println(sort.Sorted(strs))

	strs = []string{"a", "b", "c", "d"}
	fmt.Println(sort.Sorted(strs))

	ints := []int{5, 4, 3, 2}
	fmt.Println(sort.Sorted(ints))

	ints = []int{2, 3, 4, 5}
	fmt.Println(sort.Sorted(ints))

	floats := []float64{2.2, 3.3, 4.1, 4.2}
	fmt.Println(sort.Sorted(floats))

	// Output:
	// false
	// true
	// false
	// true
	// true
}

func ExampleSortedReverse() {
	strs := []string{"c", "a", "d", "b"}
	fmt.Println(sort.SortedReverse(strs))

	strs = []string{"d", "c", "b", "a"}
	fmt.Println(sort.SortedReverse(strs))

	ints := []int{2, 3, 4, 5}
	fmt.Println(sort.SortedReverse(ints))

	ints = []int{5, 4, 3, 2}
	fmt.Println(sort.SortedReverse(ints))

	floats := []float64{4.2, 4.1, 3.3, 2.1}
	fmt.Println(sort.SortedReverse(floats))

	// Output:
	// false
	// true
	// false
	// true
	// true
}

func ExampleStable() {
	strs := []string{"c", "a", "d", "b"}
	sort.Stable(strs)
	fmt.Println(strs)

	ints := []int{5, 4, 3, 2}
	sort.Stable(ints)
	fmt.Println(ints)

	floats := []float64{2.2, 3.3, 4.1, 4.2}
	sort.Stable(floats)
	fmt.Println(floats)

	// Output:
	// [a b c d]
	// [2 3 4 5]
	// [2.2 3.3 4.1 4.2]
}

func ExampleStableReverse() {
	// Sort strings in reverse order.
	strs := []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sort.StableReverse(strs)
	fmt.Println(strs)

	// Sort integers in reverse order.
	ints := []int{5, 3, 4, 2, 7, 1, 8, 12, 9, 6}
	sort.StableReverse(ints)
	fmt.Println(ints)

	// Sort floats in reverse order.
	floats := []float64{5.5, 3.3, 4.4, 2.2, 7.7, 1.1, 8.8, 12.12, 9.9, 6.6}
	sort.StableReverse(floats)
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
	sortstd.Sort(sort.SliceOrdered[string](strs))
	fmt.Println(strs)

	// Sotr strings in reverse order.
	strs = []string{"c", "a", "d", "aa", "cc", "b", "bb"}
	sortstd.Sort(sortstd.Reverse(sort.SliceOrdered[string](strs)))
	fmt.Println(strs)

	// Output:
	// [a aa b bb c cc d]
	// [d cc c bb b aa a]
}
