package types

import (
	"fmt"
	"testing"
)

func helper(m *testing.T) {
	m.Helper()
}

func TestIsType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		isType func(value any) bool
		value  any
		exp    bool
	}{
		{
			isType: IsType[int],
			value:  0,
			exp:    true,
		},
		{
			isType: IsType[int],
			value:  1,
			exp:    true,
		},
		{
			isType: IsType[int],
			value:  1.0,
			exp:    false,
		},
		{
			isType: IsType[int],
			value:  "1",
			exp:    false,
		},
		{
			isType: IsType[string],
			value:  0,
			exp:    false,
		},
		{
			isType: IsType[string],
			value:  1,
			exp:    false,
		},
	}

	for idx, prov := range tests {
		prov := prov

		t.Run(fmt.Sprintf("TestIsType_%d", idx), func(t *testing.T) {
			t.Parallel()

			if prov.isType(prov.value) != prov.exp {
				t.Errorf("IsType() = %v, want %v", prov.isType(prov.value), prov.exp)
			}
		})
	}
}
