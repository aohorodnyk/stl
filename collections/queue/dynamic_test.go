package queue_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/queue"
)

var (
	_ queue.Queue[int]      = &queue.Dynamic[int]{}
	_ queue.Queue[string]   = &queue.Dynamic[string]{}
	_ queue.Queue[testType] = &queue.Dynamic[testType]{}
)

func TestDynamicSize10(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewDynamic[string]()

	size10DynamicTest(t, queueObj)
}

func size10DynamicTest(t *testing.T, queueObj queue.Queue[string]) {
	fullList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for idx, val := range fullList {
		if !queueObj.Push(val) {
			t.Errorf("Push(%s) failed", val)
		}

		if queueObj.Length() != idx+1 {
			t.Errorf("Length() = %d, want %d", queueObj.Length(), idx+1)
		}

		if queueObj.Empty() {
			t.Error("Empty() = true, want false")
		}

		if peeked, ok := queueObj.Peek(); peeked != fullList[0] || !ok {
			t.Errorf("Peek() = %s, want %s", fullList[0], val)
		}
	}

	if !queueObj.Push("test") {
		t.Error("Push(test) failed")
	}

	for idx, val := range fullList {
		if peeked, ok := queueObj.Peek(); peeked != val || !ok {
			t.Errorf("Pop() = %s, want %s", peeked, val)
		}

		if popped, ok := queueObj.Pop(); popped != val || !ok {
			t.Errorf("Pop() = %s, want %s", popped, val)
		}

		if queueObj.Length() != len(fullList)-idx {
			t.Errorf("Length() = %d, want %d", queueObj.Length(), len(fullList)-idx-1)
		}
	}

	if queueObj.Empty() {
		t.Error("Empty() = true, want false")
	}

	if peeked, ok := queueObj.Peek(); peeked != "test" || !ok {
		t.Errorf("Peek() = %s, want test", peeked)
	}

	if popped, ok := queueObj.Pop(); popped != "test" || !ok {
		t.Errorf("Pop() = %s, want test", popped)
	}

	if queueObj.Length() != 0 {
		t.Errorf("Length() = %d, want 0", queueObj.Length())
	}

	if !queueObj.Empty() {
		t.Error("Empty() = false, want true")
	}
}
