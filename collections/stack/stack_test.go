package stack_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/stack"
)

func clearTest(t *testing.T, stackObj stack.Stack[int]) {
	if !stackObj.Empty() {
		t.Errorf("Expected stack to be empty, got %t", stackObj.Empty())
	}

	stackObj.Push(213)
	stackObj.Push(123)
	stackObj.Push(321)

	if stackObj.Length() != 3 {
		t.Errorf("Length() = %d, want %d", stackObj.Length(), 3)
	}

	if stackObj.Empty() {
		t.Errorf("Expected stack to not be empty, got %t", stackObj.Empty())
	}

	stackObj.Clear()

	if stackObj.Length() != 0 {
		t.Errorf("Length() = %d, want %d", stackObj.Length(), 0)
	}

	if val, ok := stackObj.Peek(); ok || val != 0 {
		t.Errorf("Peek() = %d, %t, want 0, false", val, ok)
	}

	if val, ok := stackObj.Pop(); ok || val != 0 {
		t.Errorf("Pop() = %d, %t, want 0, false", val, ok)
	}

	stackObj.Push(111)
	stackObj.Push(222)

	if stackObj.Length() != 2 {
		t.Errorf("Length() = %d, want %d", stackObj.Length(), 2)
	}

	if val, ok := stackObj.Pop(); !ok || val != 222 {
		t.Errorf("Pop() = %d, %t, want 222, true", val, ok)
	}

	stackObj.Push(333)

	if stackObj.Length() != 2 {
		t.Errorf("Length() = %d, want %d", stackObj.Length(), 2)
	}

	if val, ok := stackObj.Pop(); !ok || val != 333 {
		t.Errorf("Pop() = %d, %t, want 333, true", val, ok)
	}
}
