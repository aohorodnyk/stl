package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

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
	t.Helper()

	t.Run("TestNode", func(t *testing.T) {
		t.Parallel()

		ll := factory()
		nodeDoublyTest(t, ll)
	})
}

func nodeDoublyTest(t *testing.T, ll linkedlist.LinkedList[string]) {
	t.Helper()

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
}
