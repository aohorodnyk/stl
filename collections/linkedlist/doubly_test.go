package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

var (
	_ linkedlist.LinkedList[int]               = &linkedlist.DoublyAny[int]{}
	_ linkedlist.LinkedList[string]            = &linkedlist.DoublyAny[string]{}
	_ linkedlist.LinkedList[map[string]string] = &linkedlist.DoublyAny[map[string]string]{}

	_ linkedlist.Node[int]               = &linkedlist.DoublyNodeAny[int]{}
	_ linkedlist.Node[string]            = &linkedlist.DoublyNodeAny[string]{}
	_ linkedlist.Node[map[string]string] = &linkedlist.DoublyNodeAny[map[string]string]{}

	_ linkedlist.LinkedList[int]    = &linkedlist.DoublyComparable[int]{}
	_ linkedlist.LinkedList[string] = &linkedlist.DoublyComparable[string]{}

	_ linkedlist.Node[int]    = &linkedlist.DoublyNodeComparable[int]{}
	_ linkedlist.Node[string] = &linkedlist.DoublyNodeComparable[string]{}
)

func TestDoublyComparable(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyComparable[int]()
	}

	runTests(t, factory)
}

func TestDoublyComparableSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyComparableSync[int]()
	}

	runTests(t, factory)
}

func TestDoublyAny(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyAny[int]()
	}

	runTests(t, factory)
}

func TestDoublyAnySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyAnySync[int]()
	}

	runTests(t, factory)
}

func TestDoublyAnyCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyAnyCmp(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestDoublyAnySyncCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyAnySyncCmp(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestDDoublyComparable(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyComparable[string]()
	}

	runDoublyTests(t, factory)
}

func TestDDoublyComparableSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyComparableSync[string]()
	}

	runDoublyTests(t, factory)
}

func TestDDoublyAny(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyAny[string]()
	}

	runDoublyTests(t, factory)
}

func TestDDoublyAnySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyAnySync[string]()
	}

	runDoublyTests(t, factory)
}

func TestDDoublyAnyCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyAnyCmp(func(a, b string) bool { return a == b })
	}

	runDoublyTests(t, factory)
}

func TestDDoublyAnySyncCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyAnySyncCmp(func(a, b string) bool { return a == b })
	}

	runDoublyTests(t, factory)
}

func runDoublyTests(t *testing.T, factory func() linkedlist.LinkedList[string]) {
	t.Run("TestNode", func(t *testing.T) {
		t.Parallel()

		ll := factory()
		nodeDoublyTest(t, ll)
	})
}

func nodeDoublyTest(t *testing.T, ll linkedlist.LinkedList[string]) {
	values := []string{"a1", "a2", "a3", "a4", "a5"}
	for i, v := range values {
		ll.AddAt(i, v)
	}

	if ll.Length() != len(values) {
		t.Errorf("Length() = %d, want %d", ll.Length(), len(values))
	}

	node := ll.NodeFirst()
	for _, v := range values {
		if node.Value() != v {
			t.Errorf("Node.Value = %s, want %s", node.Value(), v)
		}

		node = node.Next()
	}

	node = ll.NodeLast()
	for i := len(values) - 1; i >= 0; i-- {
		if node.Value() != values[i] {
			t.Errorf("Node.Value = %s, want %s", node.Value(), values[i])
		}

		node = node.Prev()
	}

	node = ll.NodeAt(3)
	if node.Value() != values[3] {
		t.Errorf("Node.Value = %s, want %s", node.Value(), values[3])
	}

	if !ll.RemoveNode(node) {
		t.Errorf("RemoveNode() = false, want true")
	}

	if ll.Length() != len(values)-1 {
		t.Errorf("Length() = %d, want %d", ll.Length(), len(values)-1)
	}

	node2 := ll.NodeAt(3)
	if node2.Value() == values[3] {
		t.Errorf("Node2.Value = %s, want %s", node2.Value(), values[3])
	}

	if node.Next().Value() != node2.Value() {
		t.Errorf("Node.Next().Value = %s, want %s", node.Next().Value(), node2.Value())
	}

	if node.Prev().Value() != node2.Prev().Value() {
		t.Errorf("Node2.Prev().Value = %s, want %s", node2.Prev().Value(), node.Value())
	}
}
