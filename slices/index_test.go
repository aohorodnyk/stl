package slices_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/slices"
)

func TestIndex(t *testing.T) {
	t.Parallel()

	type provType struct {
		slice    []int
		target   int
		expected int
	}

	provider := []provType{
		{nil, 1, -1},
		{[]int{}, 1, -1},
		{[]int{1, 5, 3, 7, 5}, 3, 2},
		{[]int{1, 4, 3, 8, 5}, 6, -1},
	}

	for _, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("%v", prov.slice), func(t *testing.T) {
			t.Parallel()

			result := slices.Index(prov.slice, prov.target)
			if result != prov.expected {
				t.Errorf("expected: %v, got: %v", prov.expected, result)
			}
		})
	}
}

func TestIndexFunc(t *testing.T) {
	t.Parallel()

	type provType struct {
		slice    [][]int
		target   []int
		expected int
	}

	provider := []provType{
		{nil, nil, -1},
		{[][]int{}, nil, -1},
		{[][]int{}, []int{}, -1},
		{[][]int{{2, 1}, {4, 2}}, []int{1, 2}, -1},
		{[][]int{{1, 2}, {4, 2}}, []int{1, 2}, 0},
		{[][]int{{4, 2}, {1, 2}, {0, 6}}, []int{1, 2}, 1},
	}

	for _, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("%v", prov.slice), func(t *testing.T) {
			t.Parallel()

			cmp := func(i int) bool {
				return reflect.DeepEqual(prov.slice[i], prov.target)
			}

			result := slices.IndexFunc(prov.slice, cmp)
			if result != prov.expected {
				t.Errorf("expected: %v, got: %v", prov.expected, result)
			}
		})
	}
}
