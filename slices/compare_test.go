package slices_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/slices"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	type provType struct {
		slice1   []int
		slice2   []int
		expected bool
	}

	provider := []provType{
		{nil, nil, true},
		{[]int{}, nil, true},
		{[]int{}, []int{}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2, 3, 3}, false},
	}

	for _, prov := range provider {
		t.Run(fmt.Sprintf("%v", prov.slice1), func(t *testing.T) {
			t.Parallel()

			result := slices.Equal(prov.slice1, prov.slice2)
			if result != prov.expected {
				t.Errorf("expected: %v, got: %v", prov.expected, result)
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	t.Parallel()

	type customType struct {
		a int
	}

	type provType struct {
		slice1   []customType
		slice2   []customType
		expected bool
	}

	provider := []provType{
		{nil, nil, true},
		{[]customType{}, nil, true},
		{[]customType{}, []customType{}, true},
		{[]customType{{1}}, []customType{{1}}, true},
		{[]customType{{1}, {2}, {3}}, []customType{{1}, {2}, {3}}, true},
		{[]customType{{1}}, []customType{{1}, {2}}, false},
		{[]customType{{1}}, []customType{{1}, {2}, {3}}, false},
	}

	for _, prov := range provider {
		t.Run(fmt.Sprintf("%v", prov.slice1), func(t *testing.T) {
			t.Parallel()

			cmp := func(i int) bool {
				return prov.slice1[i].a == prov.slice2[i].a
			}

			result := slices.EqualFunc(prov.slice1, prov.slice2, cmp)
			if result != prov.expected {
				t.Errorf("expected: %v, got: %v", prov.expected, result)
			}
		})
	}
}
