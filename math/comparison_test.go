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

func TestMin_panicEmptyCall(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil || err != "Min function requires minimum one parameter" {
			t.Errorf("Expected panic, got \"%s\"", err)
		}
	}()

	math.Min[int]()
}

func TestMin_panicSlice(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil || err != "Min function requires minimum one parameter" {
			t.Errorf("Expected panic, got \"%s\"", err)
		}
	}()

	math.Min([]int{}...)
}

func TestMax_panicEmptyCall(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil || err != "Max function requires minimum one parameter" {
			t.Errorf("Expected panic, got \"%s\"", err)
		}
	}()

	math.Max[int]()
}

func TestMax_panicSlice(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil || err != "Max function requires minimum one parameter" {
			t.Errorf("Expected panic, got \"%s\"", err)
		}
	}()

	math.Max([]int{}...)
}
