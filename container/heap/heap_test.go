package heap_test

import (
	"fmt"
	"github.com/aohorodnyk/stl/container/heap"
	syncStl "github.com/aohorodnyk/stl/sync"
	"sync"
	"testing"
)

func TestNewMinHeapOrdered(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		init  []int
		added []int
		want  []int
	}{
		{
			name:  "nil",
			init:  nil,
			added: nil,
			want:  nil,
		},
		{
			name:  "empty",
			init:  []int{},
			added: []int{},
			want:  []int{},
		},
		{
			name:  "init 1",
			init:  []int{1},
			added: []int{},
			want:  []int{1},
		},
		{
			name:  "init 1, 24, 734, 2, 453",
			init:  []int{1, 24, 734, 2, 453},
			added: nil,
			want:  []int{1, 2, 24, 453, 734},
		},
		{
			name:  "init 1, 3, 2, added 3, 3, 2",
			init:  []int{7, 5},
			added: []int{3, 3, 2},
			want:  []int{2, 3, 3, 5, 7},
		},
		{
			name:  "init 8, 234, 1, added 1, 8, 234",
			init:  nil,
			added: []int{8, 234, 1},
			want:  []int{1, 8, 234},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			h := heap.NewMinHeapOrdered(test.init)
			for _, v := range test.added {
				h.Push(v)
			}

			for h.Len() > 0 {
				got := h.Pop()
				if got != test.want[0] {
					t.Errorf("got %v, want %v", got, test.want[0])
				}
				test.want = test.want[1:]
			}

			if len(test.want) != 0 {
				t.Errorf("want %v, but h.Len() == %d", test.want, h.Len())
			}
		})
	}
}

func TestNewMaxHeapOrdered(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		init  []int
		added []int
		want  []int
	}{
		{
			name:  "empty",
			init:  []int{},
			added: []int{},
			want:  []int{},
		},
		{
			name:  "init 1",
			init:  []int{1},
			added: []int{},
			want:  []int{1},
		},
		{
			name:  "init 1, 24, 734, 2, 453",
			init:  []int{1, 24, 734, 2, 453},
			added: nil,
			want:  []int{734, 453, 24, 2, 1},
		},
		{
			name:  "init 7, 5 added 3, 2, 3",
			init:  []int{7, 5},
			added: []int{3, 2, 3},
			want:  []int{7, 5, 3, 3, 2},
		},
		{
			name:  "init nil, added 234, 8, 1",
			init:  nil,
			added: []int{8, 234, 1},
			want:  []int{234, 8, 1},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			h := heap.NewMaxHeapOrdered(test.init)
			for _, v := range test.added {
				h.Push(v)
			}

			for h.Len() > 0 {
				got := h.Pop()
				if got != test.want[0] {
					t.Errorf("got %v, want %v", got, test.want[0])
				}

				test.want = test.want[1:]
			}

			if len(test.want) != 0 {
				t.Errorf("want %v, but h.Len() == %d", test.want, h.Len())
			}
		})
	}
}

func ExampleNewSyncMaxHeapOrdered() {
	h := heap.NewSyncMaxHeapOrdered[int](nil)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			h.Push(i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Println(h.Pop())
	}

	// Output:
	// 999
	// 998
	// 997
	// 996
	// 995
}

func ExampleNewSyncMinHeapOrdered() {
	h := heap.NewSyncMinHeapOrdered[int](nil)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			h.Push(i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Println(h.Pop())
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleNewMaxHeapOrdered() {
	h := heap.NewMaxHeapOrdered[int]([]int{253, 1, 234, 2, 453, 734})
	h.Push(5)
	h.Push(10354)
	h.Push(15354)
	h.Push(344)

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	// Output:
	// 15354
	// 10354
	// 734
	// 453
	// 344
	// 253
	// 234
	// 5
	// 2
	// 1
}

func ExampleNewMinHeapOrdered() {
	h := heap.NewMinHeapOrdered[int]([]int{253, 1, 234, 2, 453, 734})
	h.Push(5)
	h.Push(10354)
	h.Push(15354)
	h.Push(344)

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	// Output:
	// 1
	// 2
	// 5
	// 234
	// 253
	// 344
	// 453
	// 734
	// 10354
	// 15354
}

func ExampleNewHeap() {
	std := heap.NewStdInterface([]int{253, 1, 234, 2, 453, 734})
	std.LessFunc = func(i, j int) bool {
		return std.Value(i) < std.Value(j)
	}

	h := heap.NewHeap[int](syncStl.LockerDummy{}, &std)
	h.Push(5)
	h.Push(33)

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	// Output:
	// 1
	// 2
	// 5
	// 33
	// 234
	// 253
	// 453
	// 734
}

func ExampleNewHeapSync() {
	std := heap.NewStdInterface[int](nil)
	std.LessFunc = func(i, j int) bool {
		return std.Value(i) < std.Value(j)
	}

	h := heap.NewHeapSync[int](&std)

	var wg sync.WaitGroup

	data := []int{253, 1, 234, 2, 453, 734, 5, 33}
	for _, v := range data {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			h.Push(v)
		}(v)
	}

	wg.Wait()

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	// Output:
	// 1
	// 2
	// 5
	// 33
	// 234
	// 253
	// 453
	// 734
}

func ExampleNewHeapSimple() {
	std := heap.NewStdInterface[int](nil)
	std.LessFunc = func(i, j int) bool {
		return std.Value(i) < std.Value(j)
	}

	h := heap.NewHeapSimple[int](&std)

	data := []int{253, 1, 234, 2, 453, 734, 5, 33}
	for _, v := range data {
		h.Push(v)
	}

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	// Output:
	// 1
	// 2
	// 5
	// 33
	// 234
	// 253
	// 453
	// 734
}
