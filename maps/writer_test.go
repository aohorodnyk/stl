package maps_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/maps"
)

func TestSet(t *testing.T) {
	t.Parallel()

	provider := []struct {
		input     map[int]int
		key       int
		value     int
		expValue  int
		expExists bool
	}{
		{map[int]int{}, 5, 3, 0, false},
		{map[int]int{1: 3, 5: 6, 8: 2}, 12, 8, 0, false},
		{map[int]int{5: 6}, 5, 3, 6, true},
		{map[int]int{1: 3, 5: 6, 8: 2}, 5, 3, 6, true},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestSet_%d", idx), func(t *testing.T) {
			t.Parallel()

			old, ok := maps.Set(prov.input, prov.key, prov.value)
			if ok != prov.expExists {
				t.Errorf("Unexpected ok, expected: %t, got: %t", prov.expExists, ok)
			}

			if old != prov.expValue {
				t.Errorf("Unexpected value, expected: %d, got: %d", prov.expValue, old)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	provider := []struct {
		input map[int]int
		keys  []int
		exp   map[int]int
	}{
		{nil, []int{5, 3}, nil},
		{map[int]int{}, []int{5, 3}, map[int]int{}},
		{map[int]int{1: 3, 5: 6, 8: 2}, []int{12, 8}, map[int]int{1: 3, 5: 6}},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestDelete_%d", idx), func(t *testing.T) {
			t.Parallel()

			maps.Delete(prov.input, prov.keys...)

			if reflect.DeepEqual(prov.input, prov.exp) == false {
				t.Errorf("Unexpected map, expected: %v, got: %v", prov.exp, prov.input)
			}
		})
	}
}
