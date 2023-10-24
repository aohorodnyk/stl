package slices_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/slices"
)

func TestSortInt(t *testing.T) {
	t.Parallel()

	provider := []map[string][]int{
		{
			"slice":    nil,
			"expected": nil,
		},
		{
			"slice":    []int{},
			"expected": []int{},
		},
		{
			"slice":    []int{1},
			"expected": []int{1},
		},
		{
			"slice":    []int{2, 1},
			"expected": []int{1, 2},
		},
		{
			"slice":    []int{2, 1, 6, 12, 5, 11, 12, 10, 16},
			"expected": []int{1, 2, 5, 6, 10, 11, 12, 12, 16},
		},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestSortInt_%d", idx), func(t *testing.T) {
			t.Parallel()

			slice := prov["slice"]
			expected := prov["expected"]

			slices.Sort(slice)

			if !reflect.DeepEqual(slice, expected) {
				t.Errorf("Expected %v, got %v", expected, slice)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	t.Parallel()

	provider := []map[string][]string{
		{
			"slice":    nil,
			"expected": nil,
		},
		{
			"slice":    []string{},
			"expected": []string{},
		},
		{
			"slice":    []string{"a"},
			"expected": []string{"a"},
		},
		{
			"slice":    []string{"b", "a"},
			"expected": []string{"a", "b"},
		},
		{
			"slice":    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			"expected": []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		},
		{
			"slice":    []string{"b", "a", "f", "d", "e", "e", "h", "i", "c", "g", "j"},
			"expected": []string{"a", "b", "c", "d", "e", "e", "f", "g", "h", "i", "j"},
		},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestSortString_%d", idx), func(t *testing.T) {
			t.Parallel()

			slice := prov["slice"]
			expected := prov["expected"]

			slices.Sort(slice)

			if !reflect.DeepEqual(slice, expected) {
				t.Errorf("Expected %v, got %v", expected, slice)
			}
		})
	}
}

func TestReverseInt(t *testing.T) {
	t.Parallel()

	provider := []map[string][]int{
		{
			"slice":    nil,
			"expected": nil,
		},
		{
			"slice":    []int{},
			"expected": []int{},
		},
		{
			"slice":    []int{1},
			"expected": []int{1},
		},
		{
			"slice":    []int{1, 2},
			"expected": []int{2, 1},
		},
		{
			"slice":    []int{2, 1},
			"expected": []int{2, 1},
		},
		{
			"slice":    []int{2, 1, 6, 12, 5, 11, 12, 10, 16},
			"expected": []int{16, 12, 12, 11, 10, 6, 5, 2, 1},
		},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestReverseInt_%d", idx), func(t *testing.T) {
			t.Parallel()

			slice := prov["slice"]
			expected := prov["expected"]

			slices.SortReverse(slice)

			if !reflect.DeepEqual(slice, expected) {
				t.Errorf("Expected %v, got %v", expected, slice)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	t.Parallel()

	provider := []map[string][]string{
		{
			"slice":    nil,
			"expected": nil,
		},
		{
			"slice":    []string{},
			"expected": []string{},
		},
		{
			"slice":    []string{"a"},
			"expected": []string{"a"},
		},
		{
			"slice":    []string{"b", "a"},
			"expected": []string{"b", "a"},
		},
		{
			"slice":    []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"},
			"expected": []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"},
		},
		{
			"slice":    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			"expected": []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"},
		},
		{
			"slice":    []string{"b", "a", "f", "d", "e", "e", "h", "i", "c", "g", "j"},
			"expected": []string{"j", "i", "h", "g", "f", "e", "e", "d", "c", "b", "a"},
		},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestReverseString_%d", idx), func(t *testing.T) {
			t.Parallel()

			slice := prov["slice"]
			expected := prov["expected"]

			slices.SortReverse(slice)

			if !reflect.DeepEqual(slice, expected) {
				t.Errorf("Expected %v, got %v", expected, slice)
			}
		})
	}
}
