package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/linkedlist"
)

// Verify that DoublyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[int] = &linkedlist.DoublyComparable[int]{}

// Verify that DoublyComparable implements LinkedList interface.
var _ linkedlist.Node[int] = &linkedlist.DoublyNodeComparable[int]{}

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[int] = &linkedlist.SinglyComparable[int]{}

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[int] = &linkedlist.ComparableSync[int]{}

// Verify that SinglyNodeComparable implements LinkedList interface.
var _ linkedlist.Node[int] = &linkedlist.SinglyNodeComparable[int]{}

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[map[string]string] = &linkedlist.SinglyAny[map[string]string]{}

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[map[string]string] = &linkedlist.SinglyAnySync[map[string]string]{}

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.Node[map[string]string] = &linkedlist.SinglyNodeAny[map[string]string]{}

func TestDoublyComparable(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparable[int]()
	simpleTest(t, list)
}

func TestDoublyComparableRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparable[int]()
	removeLast(t, list)
}

func TestDoublyComparableEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparable[int]()
	emptyTest(t, list)
}

func TestDoublyComparableMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparable[int]()
	multipleTest(t, list)
}

func TestSinglyComparable(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparable[int]()
	simpleTest(t, list)
}

func TestSinglyComparableMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparable[int]()
	multipleTest(t, list)
}

func TestSinglyComparableRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparable[int]()
	removeLast(t, list)
}

func TestSinglyComparableEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparable[int]()
	emptyTest(t, list)
}

func TestSinglyComparableSyncRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparableSync[int]()
	removeLast(t, list)
}

func TestSinglyComparableSync(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparableSync[int]()
	simpleTest(t, list)
}

func TestSinglyComparableSyncEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparableSync[int]()
	emptyTest(t, list)
}

func TestSinglyComparableSyncMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyComparableSync[int]()
	multipleTest(t, list)
}

func TestDoublyComparableSyncRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparableSync[int]()
	removeLast(t, list)
}

func TestDoublyComparableSync(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparableSync[int]()
	simpleTest(t, list)
}

func TestDoublyComparableSyncEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparableSync[int]()
	emptyTest(t, list)
}

func TestDoublyComparableSyncMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewDoublyComparableSync[int]()
	multipleTest(t, list)
}

func TestSinglyAny(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAny[int]()
	simpleTest(t, list)
}

func TestSinglyAnyRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAny[int]()
	removeLast(t, list)
}

func TestSinglyAnyMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAny[int]()
	multipleTest(t, list)
}

func TestSinglyAnyEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAny[int]()
	emptyTest(t, list)
}

func TestSinglyAnySyncRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySync[int]()
	removeLast(t, list)
}

func TestSinglyAnySync(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySync[int]()
	simpleTest(t, list)
}

func TestSinglyAnySyncEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySync[int]()
	emptyTest(t, list)
}

func TestSinglyAnySyncMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySync[int]()
	multipleTest(t, list)
}

func TestSinglyAnyCmp(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnyCmp(func(a, b int) bool { return a == b })
	simpleTest(t, list)
}

func TestSinglyAnyCmpEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnyCmp(func(a, b int) bool { return a == b })
	emptyTest(t, list)
}

func TestSinglyAnyCmpRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnyCmp(func(a, b int) bool { return a == b })
	removeLast(t, list)
}

func TestSinglyAnyCmpMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnyCmp(func(a, b int) bool { return a == b })
	multipleTest(t, list)
}

func TestSinglyAnySyncCmp(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySyncCmp(func(a, b int) bool { return a == b })
	simpleTest(t, list)
}

func TestSinglyAnySyncCmpEmpty(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySyncCmp(func(a, b int) bool { return a == b })
	emptyTest(t, list)
}

func TestSinglyAnySyncCmpMultiple(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySyncCmp(func(a, b int) bool { return a == b })
	multipleTest(t, list)
}

func TestSinglyAnySyncCmpRemoveLast(t *testing.T) {
	t.Parallel()

	list := linkedlist.NewSinglyAnySyncCmp(func(a, b int) bool { return a == b })
	removeLast(t, list)
}

func multipleTest(t *testing.T, list linkedlist.LinkedList[int]) {
	t.Helper()

	list.AddFirst(5234)
	list.AddFirst(532)
	list.AddFirst(5234)
	list.AddFirst(532)
	list.AddFirst(123)
	list.AddFirst(123)
	list.AddLast(123)
	list.AddLast(532)
	list.AddLast(5234)

	if list.Length() != 9 {
		t.Errorf("Expected length to be 9, got %d", list.Length())
	}

	if !list.RemoveAllBy(123) {
		t.Errorf("Expected to remove all elements with value 123")
	}

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if list.RemoveAllBy(123) {
		t.Errorf("Expected to not remove any elements")
	}

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if list.IndexOf(532) != 0 {
		t.Errorf("Expected the first index of 532 to be 0, got %d", list.IndexOf(532))
	}

	if list.IndexOfLast(532) != 4 {
		t.Errorf("Expected the last index of 532 to be 4, got %d", list.IndexOfLast(532))
	}

	if !list.RemoveFirstBy(532) {
		t.Errorf("Expected to remove first element with value 523")
	}

	if list.IndexOf(532) != 1 {
		t.Errorf("Expected the first index of 532 to be 1, got %d", list.IndexOf(532))
	}

	if list.IndexOfLast(532) != 3 {
		t.Errorf("Expected the last index of 532 to be 3, got %d", list.IndexOfLast(532))
	}

	if !list.RemoveLastBy(532) {
		t.Errorf("Expected to last element with value 523")
	}

	if list.IndexOf(532) != 1 {
		t.Errorf("Expected the first index of 532 to be 1, got %d", list.IndexOf(532))
	}

	if list.IndexOfLast(532) != 1 {
		t.Errorf("Expected the last index of 532 to be 3, got %d", list.IndexOfLast(532))
	}

	if list.IndexOf(5234) != 0 {
		t.Errorf("Expected the first index of 5234 to be 0, got %d", list.IndexOf(5234))
	}

	node := list.NodeAt(0)
	if node == nil {
		t.Errorf("Expected to find node at index 0")
	}

	if node.Value() != 5234 {
		t.Errorf("Expected the value of the first node to be 5234, got %d", node.Value())
	}

	if !list.RemoveNode(node) {
		t.Errorf("Expected to remove node")
	}

	if list.IndexOf(5234) != 1 {
		t.Errorf("Expected the first index of 5234 to be 1, got %d", list.IndexOf(5234))
	}

	if list.IndexOfLast(5234) != 2 {
		t.Errorf("Expected the first index of 5234 to be 1, got %d", list.IndexOfLast(5234))
	}
}

func removeLast(t *testing.T, list linkedlist.LinkedList[int]) {
	t.Helper()

	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(1)
	list.AddLast(2)

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if list.IndexOf(1) != 0 {
		t.Errorf("Expected the first index of 1 to be 0, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 4 {
		t.Errorf("Expected the last index of 1 to be 4, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveLastBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != 0 {
		t.Errorf("Expected the first index of 1 to be 0, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 2 {
		t.Errorf("Expected the last index of 1 to be 4, got %d", list.IndexOfLast(1))
	}
}

func simpleTest(t *testing.T, list linkedlist.LinkedList[int]) {
	t.Helper()

	list.AddFirst(1)
	list.AddFirst(2)
	list.AddLast(-1)
	list.AddAt(1, 3)
	list.AddAt(100, 3)

	exp := []int{2, 3, 1, -1}

	if list.Length() != len(exp) {
		t.Errorf("Expected length %d, got %d", len(exp), list.Length())
	}

	for expIdx, expVal := range exp {
		if !list.Contains(expVal) {
			t.Errorf("Contains(%d) function. Expected true, got false", expVal)
		}

		if list.IndexOf(expVal) != expIdx {
			t.Errorf("IndexOf(%d) function. Expected %d, got %d", expVal, expIdx, list.IndexOf(expVal))
		}

		if val, ok := list.ValueAt(expIdx); !ok || val != expVal {
			t.Errorf("ValueAt(%d) function. Expected %d, got %d", expIdx, expVal, val)
		}

		if node := list.NodeAt(expIdx); node.Value() != expVal {
			t.Errorf("NodeAt(%d) function. Expected %d, got %d", expIdx, expVal, node.Value())
		}
	}

	if val, ok := list.ValueFirst(); !ok || val != exp[0] {
		t.Errorf("ValueFirst() function. Expected %d, got %d", exp[0], val)
	}

	if val, ok := list.ValueLast(); !ok || val != exp[len(exp)-1] {
		t.Errorf("ValueLast() function. Expected %d, got %d", exp[len(exp)-1], val)
	}

	if node := list.NodeFirst(); node.Value() != exp[0] {
		t.Errorf("NodeFirst() function. Expected %d, got %d", exp[0], node.Value())
	}

	if node := list.NodeLast(); node.Value() != exp[len(exp)-1] {
		t.Errorf("NodeLast() function. Expected %d, got %d", exp[len(exp)-1], node.Value())
	}

	if val, ok := list.PopLast(); !ok || val != exp[len(exp)-1] {
		t.Errorf("PopLast() function. Expected %d, got %d", exp[len(exp)-1], val)
	}

	if val, ok := list.ValueLast(); !ok || val != exp[len(exp)-2] {
		t.Errorf("ValueLast() function. Expected %d, got %d", exp[len(exp)-2], val)
	}

	if list.Length() != len(exp)-1 {
		t.Errorf("Expected length %d, got %d", len(exp)-1, list.Length())
	}

	if val, ok := list.PopAt(1); !ok || val != exp[1] {
		t.Errorf("PopAt(1) function. Expected %d, got %d", exp[1], val)
	}

	if list.Length() != len(exp)-2 {
		t.Errorf("Expected length %d, got %d", len(exp)-2, list.Length())
	}

	if val, ok := list.ValueAt(1); !ok || val != exp[2] {
		t.Errorf("ValueAt(1) function. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.PopFirst(); !ok || val != exp[0] {
		t.Errorf("PopFirst() function. Expected %d, got %d", exp[0], val)
	}

	if list.Length() != len(exp)-3 {
		t.Errorf("Expected length %d, got %d", len(exp)-3, list.Length())
	}

	if val, ok := list.ValueFirst(); !ok || val != exp[2] {
		t.Errorf("ValueFirst() function. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.ValueLast(); !ok || val != exp[2] {
		t.Errorf("ValueLast() function. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.PopFirst(); !ok || val != exp[2] {
		t.Errorf("PopFirst() function. Expected %d, got %d", exp[2], val)
	}

	if list.Length() != 0 {
		t.Errorf("Expected length %d, got %d", 0, list.Length())
	}

	if val, ok := list.ValueFirst(); ok || val != 0 {
		t.Errorf("ValueFirst() function. Expected %d, got %d", 0, val)
	}
}

func emptyTest(t *testing.T, list linkedlist.LinkedList[int]) {
	t.Helper()

	list.AddFirst(1)
	list.AddFirst(2)
	list.AddLast(-1)
	list.AddAt(1, 3)

	node := list.NodeFirst()

	if list.Length() != 4 {
		t.Errorf("Expected length %d, got %d", 4, list.Length())
	}

	if list.Empty() {
		t.Errorf("Expected empty list, got non-empty")
	}

	list.Clear()

	if list.Length() != 0 {
		t.Errorf("Expected length 0, got %d", list.Length())
	}

	if !list.Empty() {
		t.Errorf("Expected non-empty list, got empty")
	}

	if list.Contains(1) {
		t.Error("Contains function. Expected false, got true")
	}

	if list.IndexOf(1) != -1 {
		t.Errorf("IndexOf function. Expected -1, got %d", list.IndexOf(1))
	}

	if _, ok := list.ValueFirst(); ok {
		t.Error("ValueFirst function. Expected false, got true")
	}

	if _, ok := list.ValueLast(); ok {
		t.Error("ValueLast function. Expected false, got true")
	}

	if _, ok := list.ValueAt(0); ok {
		t.Error("ValueAt function. Expected false, got true")
	}

	if node := list.NodeFirst(); node == nil {
		t.Error("NodeFirst function. Expected false, got true")
	}

	if node := list.NodeLast(); node == nil {
		t.Error("NodeLast function. Expected false, got true")
	}

	if node := list.NodeAt(0); node == nil {
		t.Error("NodeAt function. Expected false, got true")
	}

	if _, ok := list.PopFirst(); ok {
		t.Error("PopFirst function. Expected false, got true")
	}

	if _, ok := list.PopLast(); ok {
		t.Error("PopLast function. Expected false, got true")
	}

	if _, ok := list.PopAt(0); ok {
		t.Error("PopAt function. Expected false, got true")
	}

	if list.RemoveAllBy(1) {
		t.Error("RemoveAllBy function. Expected false, got true")
	}

	if list.RemoveFirstBy(1) {
		t.Error("RemoveFirstBy function. Expected false, got true")
	}

	if list.RemoveLastBy(1) {
		t.Error("RemoveLastBy function. Expected false, got true")
	}

	if list.RemoveNode(node) {
		t.Error("RemoveNode function. Expected false, got true")
	}
}
