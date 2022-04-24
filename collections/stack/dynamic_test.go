package stack_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/stack"
)

var (
	_ stack.Stack[int]      = &stack.Dynamic[int]{}
	_ stack.Stack[string]   = &stack.Dynamic[string]{}
	_ stack.Stack[testType] = &stack.Dynamic[testType]{}
)

func TestDynamicClear(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewDynamic[int]()

	clearTest(t, stackObj)
}

func TestDynamicSize10(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewDynamic[string]()

	size10DynamicTest(t, stackObj)
}

func size10DynamicTest(t *testing.T, stackObj stack.Stack[string]) {
	t.Helper()

	fullList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for idx, val := range fullList {
		if !stackObj.Push(val) {
			t.Errorf("Push(%s) failed", val)
		}

		if stackObj.Length() != idx+1 {
			t.Errorf("Length() = %d, want %d", stackObj.Length(), idx+1)
		}

		if stackObj.Empty() {
			t.Error("Empty() = true, want false")
		}

		if peeked, ok := stackObj.Peek(); peeked != val || !ok {
			t.Errorf("Peek() = %s, want %s", peeked, val)
		}
	}

	if !stackObj.Push("test") {
		t.Error("Push(test) failed")
	}

	if popped, ok := stackObj.Pop(); popped != "test" || !ok {
		t.Errorf("Pop() = %s, want test", popped)
	}

	for idx := len(fullList) - 1; idx >= 0; idx-- {
		if popped, ok := stackObj.Pop(); popped != fullList[idx] || !ok {
			t.Errorf("Pop() = %s, want %s", popped, fullList[idx])
		}

		if stackObj.Length() != idx {
			t.Errorf("Length() = %d, want %d", stackObj.Length(), idx)
		}
	}
}
