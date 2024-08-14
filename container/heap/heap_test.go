package heap_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/aohorodnyk/stl/container/heap"
	syncStl "github.com/aohorodnyk/stl/sync"
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			heapOrdered := heap.NewMinHeapOrdered(test.init)
			for _, v := range test.added {
				heapOrdered.Push(v)
			}

			for heapOrdered.Len() > 0 {
				got := heapOrdered.Pop()
				if got != test.want[0] {
					t.Errorf("got %v, want %v", got, test.want[0])
				}

				test.want = test.want[1:]
			}

			if len(test.want) != 0 {
				t.Errorf("want %v, but heapOrdered.Len() == %d", test.want, heapOrdered.Len())
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			heapOrdered := heap.NewMaxHeapOrdered(test.init)
			for _, v := range test.added {
				heapOrdered.Push(v)
			}

			for heapOrdered.Len() > 0 {
				got := heapOrdered.Pop()
				if got != test.want[0] {
					t.Errorf("got %v, want %v", got, test.want[0])
				}

				test.want = test.want[1:]
			}

			if len(test.want) != 0 {
				t.Errorf("want %v, but heapOrdered.Len() == %d", test.want, heapOrdered.Len())
			}
		})
	}
}

func ExampleNewSyncMaxHeapOrdered() {
	heapOrdered := heap.NewSyncMaxHeapOrdered[int](nil)

	var wg sync.WaitGroup

	for idx := range 1000 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			heapOrdered.Push(i)
		}(idx)
	}

	wg.Wait()

	for range 5 {
		fmt.Println(heapOrdered.Pop())
	}

	// Output:
	// 999
	// 998
	// 997
	// 996
	// 995
}

func ExampleNewSyncMinHeapOrdered() {
	heapOrdered := heap.NewSyncMinHeapOrdered[int](nil)

	var wg sync.WaitGroup

	for idx := range 1000 {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			heapOrdered.Push(idx)
		}(idx)
	}

	wg.Wait()

	for range 5 {
		fmt.Println(heapOrdered.Pop())
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleNewMaxHeapOrdered() {
	heapOrdered := heap.NewMaxHeapOrdered[int]([]int{253, 1, 234, 2, 453, 734})
	heapOrdered.Push(5)
	heapOrdered.Push(10354)
	heapOrdered.Push(15354)
	heapOrdered.Push(344)

	for heapOrdered.Len() > 0 {
		fmt.Println(heapOrdered.Pop())
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
	heapOrdered := heap.NewMinHeapOrdered[int]([]int{253, 1, 234, 2, 453, 734})
	heapOrdered.Push(5)
	heapOrdered.Push(10354)
	heapOrdered.Push(15354)
	heapOrdered.Push(344)

	for heapOrdered.Len() > 0 {
		fmt.Println(heapOrdered.Pop())
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

	heapSync := heap.NewHeapSync[int](&std)

	var wg sync.WaitGroup

	data := []int{253, 1, 234, 2, 453, 734, 5, 33}
	for _, val := range data {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			heapSync.Push(v)
		}(val)
	}

	wg.Wait()

	for heapSync.Len() > 0 {
		fmt.Println(heapSync.Pop())
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

	heapSimple := heap.NewHeapSimple[int](&std)

	data := []int{253, 1, 234, 2, 453, 734, 5, 33}
	for _, v := range data {
		heapSimple.Push(v)
	}

	for heapSimple.Len() > 0 {
		fmt.Println(heapSimple.Pop())
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
