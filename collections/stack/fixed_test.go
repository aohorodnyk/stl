package stack_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/stack"
)

var (
	_ stack.Stack[int]               = &stack.Fixed[int]{}
	_ stack.Stack[string]            = &stack.Fixed[string]{}
	_ stack.Stack[map[string]string] = &stack.Fixed[map[string]string]{}
)

func TestFixed(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewFixed[int](10)

	size10Test(t, stackObj)
}

func TestFixedClear(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewFixed[int](3)

	clearTest(t, stackObj)
}

func size10Test(t *testing.T, stackObj stack.Stack[int]) {
	fullList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if stackObj.Length() != 0 {
		t.Errorf("Expected length of stack to be 0, got %d", stackObj.Length())
	}

	for idx, val := range fullList {
		if !stackObj.Push(val) {
			t.Errorf("Push(%d) failed", val)
		}

		if stackObj.Length() != idx+1 {
			t.Errorf("Length() = %d, want %d", stackObj.Length(), idx+1)
		}

		if peeked, ok := stackObj.Peek(); peeked != val || !ok {
			t.Errorf("Peek() = %d, %t, want %d, false", peeked, ok, val)
		}
	}

	if stackObj.Push(123) {
		t.Errorf("Push(%d) succeeded", 123)
	}

	if stackObj.Length() != 10 {
		t.Errorf("Length() = %d, want %d", stackObj.Length(), 10)
	}

	for idx := stackObj.Length() - 1; idx >= 0; idx-- {
		popped, ok := stackObj.Pop()
		if !ok {
			t.Errorf("Pop() failed")
		}

		if popped != fullList[idx] {
			t.Errorf("Pop() = %d, want %d", popped, fullList[idx])
		}

		if stackObj.Length() != idx {
			t.Errorf("Length() = %d, want %d", stackObj.Length(), idx)
		}
	}

	if val, ok := stackObj.Pop(); ok || val != 0 {
		t.Errorf("Pop() = %d, %t, want 0, false", val, ok)
	}

	if stackObj.Length() != 0 {
		t.Errorf("Length() = %d, want %d", stackObj.Length(), 10)
	}

	if val, ok := stackObj.Peek(); ok || val != 0 {
		t.Errorf("Peek() = %d, %t, want 0, false", val, ok)
	}
}
