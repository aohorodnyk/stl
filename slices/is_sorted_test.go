package slices_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleIsSorted() {
	fmt.Println("Empty int:", slices.IsSorted([]int{}))
	fmt.Println("Sorted int:", slices.IsSorted([]int{1, 34}))
	fmt.Println("Unsorted int:", slices.IsSorted([]int{5, 25, 13}))

	fmt.Println()

	fmt.Println("Empty string:", slices.IsSorted([]string{}))
	fmt.Println("Sorted string:", slices.IsSorted([]string{"a", "b"}))
	fmt.Println("Unsorted string:", slices.IsSorted([]string{"a", "d", "b"}))

	// Output:
	// Empty int: true
	// Sorted int: true
	// Unsorted int: false
	//
	// Empty string: true
	// Sorted string: true
	// Unsorted string: false
}

func TestIsSorted(t *testing.T) {
	t.Parallel()

	provider := []struct {
		slice []int
		exp   bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 3}, true},
		{[]int{1, 3, 5}, true},
		{[]int{1, 5}, true},
		{[]int{1, 5, 4}, false},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestIsSorted_%d", idx), func(t *testing.T) {
			t.Parallel()

			cur := slices.IsSorted(prov.slice)
			if cur != prov.exp {
				t.Errorf("Expected %t, got %t", prov.exp, cur)
			}
		})
	}
}
