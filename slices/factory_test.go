package slices_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/stl/slices"
)

func ExampleMake2D() {
	matrix := slices.Make2D[int](3, 2)
	fmt.Println(matrix)

	matrix = slices.Make2D[int](3, 2)
	fmt.Println(matrix)

	matrix = slices.Make2D[int](3, 0)
	fmt.Println(matrix)

	// Output:
	// [[0 0] [0 0] [0 0]]
	// [[0 0] [0 0] [0 0]]
	// [[] [] []]
}

func ExampleMake3D() {
	matrix := slices.Make3D[int](3, 2, 1)
	fmt.Println(matrix)

	matrix = slices.Make3D[int](3, 2, 0)
	fmt.Println(matrix)

	// Output:
	// [[[0] [0]] [[0] [0]] [[0] [0]]]
	// [[[] []] [[] []] [[] []]]
}

func ExampleMake4D() {
	matrix := slices.Make4D[int](3, 2, 1, 1)
	fmt.Println(matrix)

	matrix = slices.Make4D[int](3, 2, 0, 1)
	fmt.Println(matrix)

	// Output:
	// [[[[0]] [[0]]] [[[0]] [[0]]] [[[0]] [[0]]]]
	// [[[] []] [[] []] [[] []]]
}

func TestMakeMatrix(t *testing.T) {
	t.Parallel()

	provider := []struct {
		rows, cols int
		exp        [][]int
	}{
		{0, 0, [][]int{}},
		{1, 0, [][]int{{}}},
		{0, 1, [][]int{}},
		{1, 1, [][]int{{0}}},
		{2, 2, [][]int{{0, 0}, {0, 0}}},
		{3, 2, [][]int{{0, 0}, {0, 0}, {0, 0}}},
	}

	for _, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("rows=%d,cols=%d", prov.rows, prov.cols), func(t *testing.T) {
			t.Parallel()

			matrix := slices.Make2D[int](prov.rows, prov.cols)

			if !reflect.DeepEqual(matrix, prov.exp) {
				t.Errorf("expected %v, got %v", prov.exp, matrix)
			}
		})
	}
}

func Test3D(t *testing.T) {
	t.Parallel()

	provider := []struct {
		rows, cols, tubes int
		exp               [][][]int
	}{
		{0, 0, 0, [][][]int{}},
		{1, 0, 0, [][][]int{{}}},
		{0, 1, 0, [][][]int{}},
		{0, 0, 1, [][][]int{}},
		{1, 1, 0, [][][]int{{{}}}},
		{1, 0, 1, [][][]int{{}}},
		{0, 1, 1, [][][]int{}},
		{1, 1, 1, [][][]int{{{0}}}},
		{2, 2, 2, [][][]int{{{0, 0}, {0, 0}}, {{0, 0}, {0, 0}}}},
		{3, 2, 1, [][][]int{{{0}, {0}}, {{0}, {0}}, {{0}, {0}}}},
	}

	for _, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("rows=%d,cols=%d,tubes=%d", prov.rows, prov.cols, prov.tubes), func(t *testing.T) {
			t.Parallel()

			matrix := slices.Make3D[int](prov.rows, prov.cols, prov.tubes)

			if !reflect.DeepEqual(matrix, prov.exp) {
				t.Errorf("expected %v, got %v", prov.exp, matrix)
			}
		})
	}
}

func Test4D(t *testing.T) {
	t.Parallel()

	provider := []struct {
		rows, cols, tubes, cells int
		exp                      [][][][]int
	}{
		{0, 0, 0, 0, [][][][]int{}},
		{1, 0, 0, 0, [][][][]int{{}}},
		{0, 1, 0, 0, [][][][]int{}},
		{0, 0, 1, 0, [][][][]int{}},
		{0, 0, 0, 1, [][][][]int{}},
		{1, 1, 0, 0, [][][][]int{{{}}}},
		{1, 0, 1, 0, [][][][]int{{}}},
		{1, 0, 0, 1, [][][][]int{{}}},
		{0, 1, 1, 0, [][][][]int{}},
		{0, 1, 0, 1, [][][][]int{}},
		{0, 0, 1, 1, [][][][]int{}},
		{1, 1, 1, 0, [][][][]int{{{{}}}}},
		{1, 1, 0, 1, [][][][]int{{{}}}},
		{1, 0, 1, 1, [][][][]int{{}}},
		{0, 1, 1, 1, [][][][]int{}},
		{1, 1, 1, 1, [][][][]int{{{{0}}}}},
		{2, 2, 2, 2, [][][][]int{{{{0, 0}, {0, 0}}, {{0, 0}, {0, 0}}}, {{{0, 0}, {0, 0}}, {{0, 0}, {0, 0}}}}},
		{3, 2, 1, 1, [][][][]int{{{{0}}, {{0}}}, {{{0}}, {{0}}}, {{{0}}, {{0}}}}},
	}

	for _, prov := range provider {
		prov := prov

		testName := fmt.Sprintf("rows=%d,cols=%d,tubes=%d,cells=%d", prov.rows, prov.cols, prov.tubes, prov.cells)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			matrix := slices.Make4D[int](prov.rows, prov.cols, prov.tubes, prov.cells)

			if !reflect.DeepEqual(matrix, prov.exp) {
				t.Errorf("expected %v, got %v", prov.exp, matrix)
			}
		})
	}
}
