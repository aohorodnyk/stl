package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

func TestSSinglyComparable(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyComparable[string]()
	}

	runSinglyTests(t, factory)
}

func TestSSinglyComparableSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyComparableSync[string]()
	}

	runSinglyTests(t, factory)
}

func TestSSinglyAny(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyAny[string]()
	}

	runSinglyTests(t, factory)
}

func TestSSinglyAnySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyAnySync[string]()
	}

	runSinglyTests(t, factory)
}

func TestSSinglyAnyCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyAnyCmp(func(a, b string) bool { return a == b })
	}

	runSinglyTests(t, factory)
}

func TestSSinglyAnySyncCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[string] {
		return linkedlist.NewSinglyAnySyncCmp(func(a, b string) bool { return a == b })
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

	node := ll.NodeFirst()
	node.Prev()
}
