package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/linkedlist"
)

func TestSinglyAnyCustomCmp(t *testing.T) {
	type test struct {
		a, b int
	}

	cmp := func(a, b any) bool {
		ta := a.(test)
		tb := b.(test)

		return ta.a == tb.b || ta.b == tb.a
	}

	ll := linkedlist.NewSinglyAnyCmp[test](cmp)
	ll.AddFirst(test{a: 1, b: 5})

	if ll.IndexOf(test{a: 6, b: 1}) != 0 {
		t.Errorf("Expected index of 0, got %d", ll.IndexOf(test{a: 6, b: 1}))
	}

	if ll.IndexOf(test{a: 5, b: 3}) != 0 {
		t.Errorf("Expected index of 0, got %d", ll.IndexOf(test{a: 5, b: 3}))
	}

	if ll.IndexOf(test{a: 3, b: 3}) != -1 {
		t.Errorf("Expected index of -1, got %d", ll.IndexOf(test{a: 3, b: 3}))
	}
}
