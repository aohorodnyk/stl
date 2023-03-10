package math_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/math"
)

func TestMin(t *testing.T) {
	t.Parallel()

	provider := []struct {
		nums []int
		exp  int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{25}, 25},
		{[]int{25, 3}, 3},
		{[]int{523, 32, 625, -235, 23}, -235},
		{[]int{-1245, 523, 32, 625, -235, 23}, -1245},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestMin_%d", idx), func(t *testing.T) {
			t.Parallel()

			cur := math.Min(prov.nums...)
			if cur != prov.exp {
				t.Errorf("Expected %d, got %d", prov.exp, cur)
			}
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	provider := []struct {
		nums []int
		exp  int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{25}, 25},
		{[]int{25, 3}, 25},
		{[]int{523, 32, 625, -235, 23}, 625},
		{[]int{-1245, 523, 32, 625, -235, 23}, 625},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestMin_%d", idx), func(t *testing.T) {
			t.Parallel()

			cur := math.Max(prov.nums...)
			if cur != prov.exp {
				t.Errorf("Expected %d, got %d", prov.exp, cur)
			}
		})
	}
}
