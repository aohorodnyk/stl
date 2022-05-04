package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

var (
	_ linkedlist.LinkedList[int]               = &linkedlist.Singly[int]{}
	_ linkedlist.LinkedList[string]            = &linkedlist.Singly[string]{}
	_ linkedlist.LinkedList[map[string]string] = &linkedlist.Singly[map[string]string]{}

	_ linkedlist.Node[int]               = &linkedlist.SinglyNode[int]{}
	_ linkedlist.Node[string]            = &linkedlist.SinglyNode[string]{}
	_ linkedlist.Node[map[string]string] = &linkedlist.SinglyNode[map[string]string]{}

	_ linkedlist.LinkedList[int]    = &linkedlist.Singly[int]{}
	_ linkedlist.LinkedList[string] = &linkedlist.Singly[string]{}

	_ linkedlist.Node[int]    = &linkedlist.SinglyNode[int]{}
	_ linkedlist.Node[string] = &linkedlist.SinglyNode[string]{}
)

func TestSingly(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSingly[int]()
	}

	runTests(t, factory)
}

func TestSinglySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglySync[int]()
	}

	runTests(t, factory)
}

func TestSinglyFunc(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyFunc(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestSinglyFuncSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyFuncSync(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestSSingly(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSingly[string]()
	}

	runSinglyTests(t, factory)
}

func TestSSinglySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglySync[string]()
	}

	runSinglyTests(t, factory)
}

func TestSSinglyFunc(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyFunc(func(a, b string) bool { return a == b })
	}

	runSinglyTests(t, factory)
}

func TestSSinglyFuncSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyFuncSync(func(a, b string) bool { return a == b })
	}

	runSinglyTests(t, factory)
}

func runSinglyTests(t *testing.T, factory func() linkedlist.LinkedList[string]) {
	t.Run("TestNode", func(t *testing.T) {
		t.Parallel()

		ll := factory()
		prevPanicSinglyTest(t, ll)
	})
}

func prevPanicSinglyTest(t *testing.T, ll linkedlist.LinkedList[string]) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	ll.AddFirst("a")
	ll.AddFirst("b")
	ll.AddFirst("c")
	ll.AddFirst("d")

	node := ll.NodeAt(1)
	if !ll.RemoveNode(node) {
		t.Errorf("RemoveNode failed")
	}

	if ll.Length() != 3 {
		t.Errorf("Length failed")
	}

	if ll.NodeAt(1).Value() != node.Next().Value() {
		t.Errorf("NodeAt failed")
	}

	node.Prev()
}

func TestSinglyFuncCustom(t *testing.T) {
	t.Parallel()

	type test struct {
		a, b int
	}

	cmp := func(a, b test) bool {
		return a.a == b.b || a.b == b.a
	}

	list := linkedlist.NewSinglyFunc(cmp)
	list.AddFirst(test{a: 1, b: 5})

	if list.IndexOf(test{a: 6, b: 1}) != 0 {
		t.Errorf("Expected index of 0, got %d", list.IndexOf(test{a: 6, b: 1}))
	}

	if list.IndexOf(test{a: 5, b: 3}) != 0 {
		t.Errorf("Expected index of 0, got %d", list.IndexOf(test{a: 5, b: 3}))
	}

	if list.IndexOf(test{a: 3, b: 3}) != -1 {
		t.Errorf("Expected index of -1, got %d", list.IndexOf(test{a: 3, b: 3}))
	}
}
