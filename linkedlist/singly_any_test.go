package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/linkedlist"
)

func TestSinglyAnyCustomCmp(t *testing.T) {
	t.Parallel()

	type test struct {
		a, b int
	}

	cmp := func(a, b test) bool {
		return a.a == b.b || a.b == b.a
	}

	list := linkedlist.NewSinglyAnyCmp(cmp)
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
