package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

var (
	_ linkedlist.LinkedList[int]               = &linkedlist.SinglyAny[int]{}
	_ linkedlist.LinkedList[string]            = &linkedlist.SinglyAny[string]{}
	_ linkedlist.LinkedList[map[string]string] = &linkedlist.SinglyAny[map[string]string]{}

	_ linkedlist.Node[int]               = &linkedlist.SinglyNodeAny[int]{}
	_ linkedlist.Node[string]            = &linkedlist.SinglyNodeAny[string]{}
	_ linkedlist.Node[map[string]string] = &linkedlist.SinglyNodeAny[map[string]string]{}

	_ linkedlist.LinkedList[int]    = &linkedlist.SinglyComparable[int]{}
	_ linkedlist.LinkedList[string] = &linkedlist.SinglyComparable[string]{}

	_ linkedlist.Node[int]    = &linkedlist.SinglyNodeComparable[int]{}
	_ linkedlist.Node[string] = &linkedlist.SinglyNodeComparable[string]{}
)

func TestSinglyComparable(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyComparable[int]()
	}

	runTests(t, factory)
}

func TestSinglyComparableSync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyComparableSync[int]()
	}

	runTests(t, factory)
}

func TestSinglyAny(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyAny[int]()
	}

	runTests(t, factory)
}

func TestSinglyAnySync(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyAnySync[int]()
	}

	runTests(t, factory)
}

func TestSinglyAnyCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyAnyCmp(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

func TestSinglyAnySyncCmp(t *testing.T) {
	t.Parallel()

	factory := func() linkedlist.LinkedList[int] {
		return linkedlist.NewSinglyAnySyncCmp(func(a, b int) bool { return a == b })
	}

	runTests(t, factory)
}

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
