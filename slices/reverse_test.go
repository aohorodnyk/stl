package slices_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/slices"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	provider := []struct{ input, exp []int }{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestReverse_%d", idx), func(t *testing.T) {
			t.Parallel()

			slices.Reverse(prov.input)
			if !reflect.DeepEqual(prov.input, prov.exp) {
				t.Errorf("Expected %v, got %v", prov.exp, prov.input)
			}
		})
	}
}
