package types_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/types"
)

func TestCast_int(t *testing.T) {
	t.Parallel()

	provider := []struct {
		inp any
		exp int
		ok  bool
	}{
		{0, 0, true},
		{-235234, -235234, true},
		{math.MinInt, math.MinInt, true},
		{math.MaxInt, math.MaxInt, true},
		{234.33, 0, false},
		{-623.0, 0, false},
		{"", 0, false},
		{new(int), 0, false},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestCast_int_%d", idx), func(t *testing.T) {
			t.Parallel()

			cur, ok := types.Cast[int](prov.inp)
			if prov.ok != ok {
				t.Errorf("Ok expected true, but got %t", ok)
			} else if cur != prov.exp {
				t.Errorf("Converted value is expected %d, but got %d", prov.exp, cur)
			}
		})
	}
}

func TestCast_string(t *testing.T) {
	t.Parallel()

	provider := []struct {
		inp any
		exp string
		ok  bool
	}{
		{"", "", true},
		{"24", "24", true},
		{"text long", "text long", true},
		{new(string), "", false},
		{234, "", false},
		{55.3, "", false},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestCast_int_%d", idx), func(t *testing.T) {
			t.Parallel()

			cur, ok := types.Cast[string](prov.inp)
			if prov.ok != ok {
				t.Errorf("Ok expected true, but got %t", ok)
			} else if cur != prov.exp {
				t.Errorf("Converted value is expected %s, but got %s", prov.exp, cur)
			}
		})
	}
}

func TestCast(t *testing.T) {
	t.Parallel()

	valInt := 234
	curIntPointer, ok := types.Cast[*int](&valInt)

	if !ok {
		t.Errorf("Ok expected true, but got %t", ok)
	}

	if !reflect.DeepEqual(curIntPointer, &valInt) {
		t.Errorf("Converted value is expected %d, but got %d", valInt, *curIntPointer)
	}

	valStr := "my text"
	curStrPointer, ok := types.Cast[*string](&valStr)

	if !ok {
		t.Errorf("Ok expected true, but got %t", ok)
	}

	if !reflect.DeepEqual(curStrPointer, &valStr) {
		t.Errorf("Converted value is expected %s, but got %s", valStr, *curStrPointer)
	}
}

func TestCastOk(t *testing.T) {
	t.Parallel()

	callback := func(val any, ok bool) (any, bool) {
		return val, ok
	}

	valInt := 532
	curIntPointer, ok := types.CastOk[*int](callback(&valInt, true))

	if !ok {
		t.Errorf("Ok expected true, but got %t", ok)
	}

	if !reflect.DeepEqual(curIntPointer, &valInt) {
		t.Errorf("Converted value is expected %d, but got %d", valInt, *curIntPointer)
	}

	curIntPointer, ok = types.CastOk[*int](callback(&valInt, false))

	if ok {
		t.Errorf("Ok expected false, but got %t", ok)
	}

	if curIntPointer != nil {
		t.Errorf("Converted value is expected nil, but got %d", *curIntPointer)
	}
}

func TestRef(t *testing.T) {
	t.Parallel()

	valInt := 234
	curIntPointer := types.Ref[int](valInt)

	if valInt != *curIntPointer {
		t.Errorf("Converted int is expected %d, but got %d", valInt, *curIntPointer)
	}

	valStr := "my text"
	curStrPointer := types.Ref[string](valStr)

	if valStr != *curStrPointer {
		t.Errorf("Converted string is expected %s, but got %s", valStr, *curStrPointer)
	}

	curNilIntPointer := types.Ref[*int](nil)
	if *curNilIntPointer != nil {
		t.Errorf("Converted nil is expected 0, but got %d", curNilIntPointer)
	}
}
