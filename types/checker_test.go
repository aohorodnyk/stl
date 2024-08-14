package types_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/types"
)

func TestIsType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		IsType func(value any) bool
		value  any
		exp    bool
	}{
		{
			IsType: types.IsType[int],
			value:  0,
			exp:    true,
		},
		{
			IsType: types.IsType[int],
			value:  1,
			exp:    true,
		},
		{
			IsType: types.IsType[int],
			value:  1.0,
			exp:    false,
		},
		{
			IsType: types.IsType[int],
			value:  "1",
			exp:    false,
		},
		{
			IsType: types.IsType[string],
			value:  0,
			exp:    false,
		},
		{
			IsType: types.IsType[string],
			value:  1,
			exp:    false,
		},
	}

	for idx, prov := range tests {
		t.Run(fmt.Sprintf("Testtypes.IsType_%d", idx), func(t *testing.T) {
			t.Parallel()

			if prov.IsType(prov.value) != prov.exp {
				t.Errorf("types.IsType() = %v, want %v", prov.IsType(prov.value), prov.exp)
			}
		})
	}
}
