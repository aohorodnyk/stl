package math_test

import (
	"fmt"
	mathstd "math"
	"testing"

	"github.com/aohorodnyk/stl/math"
)

func TestAbs_int8(t *testing.T) {
	t.Parallel()

	provider := []struct{ num, exp int8 }{
		{mathstd.MinInt8, mathstd.MinInt8},
		{mathstd.MinInt8 + 1, mathstd.MaxInt8},
		{0, 0},
		{-5, 5},
		{mathstd.MaxInt8, mathstd.MaxInt8},
	}

	for _, provVal := range provider {
		t.Run(fmt.Sprintf("TestAbs_%T_%d", provVal.num, provVal.num), func(t *testing.T) {
			t.Parallel()

			cur := math.Abs(provVal.num)
			if cur != provVal.exp {
				t.Errorf("Expected %d, got %d", provVal.exp, cur)
			}
		})
	}
}

func TestAbs_int16(t *testing.T) {
	t.Parallel()

	provider := []struct{ num, exp int16 }{
		{mathstd.MinInt16, mathstd.MinInt16},
		{mathstd.MinInt16 + 1, mathstd.MaxInt16},
		{0, 0},
		{-5, 5},
		{mathstd.MaxInt16, mathstd.MaxInt16},
	}

	for _, provVal := range provider {
		t.Run(fmt.Sprintf("TestAbs_%T_%d", provVal.num, provVal.num), func(t *testing.T) {
			t.Parallel()

			cur := math.Abs(provVal.num)
			if cur != provVal.exp {
				t.Errorf("Expected %d, got %d", provVal.exp, cur)
			}
		})
	}
}

func TestAbs_int32(t *testing.T) {
	t.Parallel()

	provider := []struct{ num, exp int32 }{
		{mathstd.MinInt32, mathstd.MinInt32},
		{mathstd.MinInt32 + 1, mathstd.MaxInt32},
		{0, 0},
		{-5, 5},
		{mathstd.MaxInt32, mathstd.MaxInt32},
	}

	for _, provVal := range provider {
		t.Run(fmt.Sprintf("TestAbs_%T_%d", provVal.num, provVal.num), func(t *testing.T) {
			t.Parallel()

			cur := math.Abs(provVal.num)
			if cur != provVal.exp {
				t.Errorf("Expected %d, got %d", provVal.exp, cur)
			}
		})
	}
}

func TestAbs_int64(t *testing.T) {
	t.Parallel()

	provider := []struct{ num, exp int64 }{
		{mathstd.MinInt64, mathstd.MinInt64},
		{mathstd.MinInt64 + 1, mathstd.MaxInt64},
		{0, 0},
		{-5, 5},
		{mathstd.MaxInt64, mathstd.MaxInt64},
	}

	for _, provVal := range provider {
		t.Run(fmt.Sprintf("TestAbs_%T_%d", provVal.num, provVal.num), func(t *testing.T) {
			t.Parallel()

			cur := math.Abs(provVal.num)
			if cur != provVal.exp {
				t.Errorf("Expected %d, got %d", provVal.exp, cur)
			}
		})
	}
}
