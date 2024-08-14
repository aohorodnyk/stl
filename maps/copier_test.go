package maps_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/maps"
)

func TestCopy(t *testing.T) {
	t.Parallel()

	provider := []struct {
		dest map[string]int
		src  map[string]int
		exp  map[string]int
	}{
		{
			map[string]int{"one": 1000},
			map[string]int{"one": 1, "eleven": 11, "twenty": 20},
			map[string]int{"one": 1, "eleven": 11, "twenty": 20},
		},
		{
			map[string]int{"one": 1000, "three": 3},
			map[string]int{"one": 1, "eleven": 11, "twenty": 20},
			map[string]int{"one": 1, "eleven": 11, "twenty": 20, "three": 3},
		},
		{
			map[string]int{},
			map[string]int{"one": 1, "eleven": 11, "twenty": 20},
			map[string]int{"one": 1, "eleven": 11, "twenty": 20},
		},
		{
			map[string]int{"three": 3},
			map[string]int{},
			map[string]int{"three": 3},
		},
		{
			map[string]int{"three": 3},
			nil,
			map[string]int{"three": 3},
		},
		{
			nil,
			map[string]int{"one": 1},
			nil,
		},
	}

	for idx, prov := range provider {
		t.Run(fmt.Sprintf("TestCopy_%d", idx), func(t *testing.T) {
			t.Parallel()

			maps.Copy(prov.dest, prov.src)

			if !reflect.DeepEqual(prov.dest, prov.exp) {
				t.Errorf("Destination is not equal to source. dest: %v, exp: %v", prov.dest, prov.exp)
			}
		})
	}
}

func TestClone(t *testing.T) {
	t.Parallel()

	provider := []struct {
		src map[string]string
		exp map[string]string
	}{
		{nil, map[string]string{}},
		{map[string]string{}, map[string]string{}},
		{map[string]string{"one": "one"}, map[string]string{"one": "one"}},
		{map[string]string{"one": "one", "two": "three"}, map[string]string{"one": "one", "two": "three"}},
	}

	for idx, prov := range provider {
		t.Run(fmt.Sprintf("TestClone_%d", idx), func(t *testing.T) {
			t.Parallel()

			dest := maps.Clone(prov.src)

			if !reflect.DeepEqual(dest, prov.exp) {
				t.Errorf("Cloned is not equal to source. dest: %v, exp: %v", dest, prov.exp)
			}
		})
	}
}
