package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/linkedlist"
)

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[int] = &linkedlist.SinglyComparable[int]{}

// Verify that SinglyNodeComparable implements LinkedList interface.
var _ linkedlist.Node[int] = &linkedlist.SinglyNodeComparable[int]{}

// Verify that SinglyComparable implements LinkedList interface.
var _ linkedlist.LinkedList[int] = &linkedlist.SinglyComparableSync[int]{}

func TestSinglyComparable(t *testing.T) {
	list := linkedlist.NewSinglyComparable[int]()
	simpleTest(t, list)
}

func TestSinglyComparableSync(t *testing.T) {
	list := linkedlist.NewSinglyComparableSync[int]()
	simpleTest(t, list)
}

func TestSinglyComparableEmpty(t *testing.T) {
	list := linkedlist.NewSinglyComparable[int]()
	emptyTest(t, list)
}

func TestSinglyComparableSyncEmpty(t *testing.T) {
	list := linkedlist.NewSinglyComparableSync[int]()
	emptyTest(t, list)
}

func simpleTest(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddLast(-1)
	list.AddAt(1, 3)

	exp := []int{2, 3, 1, -1}

	if list.Length() != len(exp) {
		t.Errorf("Expected length %d, got %d", len(exp), list.Length())
	}

	for expIdx, expVal := range exp {
		if !list.Contains(expVal) {
			t.Errorf("Contains(%d) fucntion. Expected true, got false", expVal)
		}

		if list.IndexOf(expVal) != expIdx {
			t.Errorf("IndexOf(%d) fucntion. Expected %d, got %d", expVal, expIdx, list.IndexOf(expVal))
		}

		if val, ok := list.ValueAt(expIdx); val != expVal || !ok {
			t.Errorf("ValueAt(%d) fucntion. Expected %d, got %d", expIdx, expVal, val)
		}

		if node, ok := list.NodeAt(expIdx); node.Value() != expVal || !ok {
			t.Errorf("NodeAt(%d) fucntion. Expected %d, got %d", expIdx, expVal, node.Value())
		}
	}

	if val, ok := list.ValueFirst(); val != exp[0] || !ok {
		t.Errorf("ValueFirst() fucntion. Expected %d, got %d", exp[0], val)
	}

	if val, ok := list.ValueLast(); val != exp[len(exp)-1] || !ok {
		t.Errorf("ValueLast() fucntion. Expected %d, got %d", exp[len(exp)-1], val)
	}

	if node, ok := list.NodeFirst(); node.Value() != exp[0] || !ok {
		t.Errorf("NodeFirst() fucntion. Expected %d, got %d", exp[0], node.Value())
	}

	if node, ok := list.NodeLast(); node.Value() != exp[len(exp)-1] || !ok {
		t.Errorf("NodeLast() fucntion. Expected %d, got %d", exp[len(exp)-1], node.Value())
	}

	if val, ok := list.PopLast(); val != exp[len(exp)-1] || !ok {
		t.Errorf("PopLast() fucntion. Expected %d, got %d", exp[len(exp)-1], val)
	}

	if val, ok := list.ValueLast(); val != exp[len(exp)-2] || !ok {
		t.Errorf("ValueLast() fucntion. Expected %d, got %d", exp[len(exp)-2], val)
	}

	if list.Length() != len(exp)-1 {
		t.Errorf("Expected length %d, got %d", len(exp)-1, list.Length())
	}

	if val, ok := list.PopAt(1); val != exp[1] || !ok {
		t.Errorf("PopAt(1) fucntion. Expected %d, got %d", exp[1], val)
	}

	if list.Length() != len(exp)-2 {
		t.Errorf("Expected length %d, got %d", len(exp)-2, list.Length())
	}

	if val, ok := list.ValueAt(1); val != exp[2] || !ok {
		t.Errorf("ValueAt(1) fucntion. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.PopFirst(); val != exp[0] || !ok {
		t.Errorf("PopFirst() fucntion. Expected %d, got %d", exp[0], val)
	}

	if list.Length() != len(exp)-3 {
		t.Errorf("Expected length %d, got %d", len(exp)-3, list.Length())
	}

	if val, ok := list.ValueFirst(); val != exp[2] || !ok {
		t.Errorf("ValueFirst() fucntion. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.ValueLast(); val != exp[2] || !ok {
		t.Errorf("ValueLast() fucntion. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.PopFirst(); val != exp[2] || !ok {
		t.Errorf("PopFirst() fucntion. Expected %d, got %d", exp[2], val)
	}

	if list.Length() != 0 {
		t.Errorf("Expected length %d, got %d", 0, list.Length())
	}

	if val, ok := list.ValueFirst(); val != 0 || ok {
		t.Errorf("ValueFirst() fucntion. Expected %d, got %d", 0, val)
	}
}

func emptyTest(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddLast(-1)
	list.AddAt(1, 3)

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
		t.Error("Contains fucntion. Expected false, got true")
	}

	if list.IndexOf(1) != -1 {
		t.Errorf("IndexOf fucntion. Expected -1, got %d", list.IndexOf(1))
	}

	if _, ok := list.ValueFirst(); ok {
		t.Error("ValueFirst fucntion. Expected false, got true")
	}

	if _, ok := list.ValueLast(); ok {
		t.Error("ValueLast fucntion. Expected false, got true")
	}

	if _, ok := list.ValueAt(0); ok {
		t.Error("ValueAt fucntion. Expected false, got true")
	}

	if _, ok := list.NodeFirst(); ok {
		t.Error("NodeFirst fucntion. Expected false, got true")
	}

	if _, ok := list.NodeLast(); ok {
		t.Error("NodeLast fucntion. Expected false, got true")
	}

	if _, ok := list.NodeAt(0); ok {
		t.Error("NodeAt fucntion. Expected false, got true")
	}

	if _, ok := list.PopFirst(); ok {
		t.Error("PopFirst fucntion. Expected false, got true")
	}

	if _, ok := list.PopLast(); ok {
		t.Error("PopLast fucntion. Expected false, got true")
	}

	if _, ok := list.PopAt(0); ok {
		t.Error("PopAt fucntion. Expected false, got true")
	}
}
