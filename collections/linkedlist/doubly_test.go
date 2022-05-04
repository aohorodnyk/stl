package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

var (
	_ linkedlist.LinkedList[int]               = &linkedlist.DoublyFunc[int]{}
	_ linkedlist.LinkedList[string]            = &linkedlist.DoublyFunc[string]{}
	_ linkedlist.LinkedList[map[string]string] = &linkedlist.DoublyFunc[map[string]string]{}

	_ linkedlist.Node[int]               = &linkedlist.DoublyNodeAny[int]{}
	_ linkedlist.Node[string]            = &linkedlist.DoublyNodeAny[string]{}
	_ linkedlist.Node[map[string]string] = &linkedlist.DoublyNodeAny[map[string]string]{}

	_ linkedlist.LinkedList[int]    = &linkedlist.Doubly[int]{}
	_ linkedlist.LinkedList[string] = &linkedlist.Doubly[string]{}

	_ linkedlist.Node[int]    = &linkedlist.DoublyNode[int]{}
	_ linkedlist.Node[string] = &linkedlist.DoublyNode[string]{}
)

func TestDoubly(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoubly[int]()
	}

	runTests(t, factory)
}

func TestDoublySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublySync[int]()
	}

	runTests(t, factory)
}

func TestDoublyFunc(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyFunc(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestDoublyFuncSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewDoublyFuncSync(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestDDoubly(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoubly[string]()
	}

	runDoublyTests(t, factory)
}

func TestDDoublySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublySync[string]()
	}

	runDoublyTests(t, factory)
}

func TestDDoublyFunc(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyFunc(func(a, b string) bool { return a == b })
	}

	runDoublyTests(t, factory)
}

func TestDDoublyFuncSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewDoublyFuncSync(func(a, b string) bool { return a == b })
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
