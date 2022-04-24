package queue_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/queue"
)

var (
	_ queue.Queue[int]      = &queue.Fixed[int]{}
	_ queue.Queue[string]   = &queue.Fixed[string]{}
	_ queue.Queue[testType] = &queue.Fixed[testType]{}
)

func TestFixed(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewFixed[int](10)

	size10Test(t, queueObj)
}

func TestFixedClear(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewFixed[int](3)

	clearTest(t, queueObj)
}

func size10Test(t *testing.T, queueObj queue.Queue[int]) {
	fullList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if queueObj.Length() != 0 {
		t.Errorf("Expected length of queue to be 0, got %d", queueObj.Length())
	}

	for idx, value := range fullList {
		if !queueObj.Push(value) {
			t.Errorf("Expected Push to return true, got false")
		}

		if queueObj.Length() != idx+1 {
			t.Errorf("Expected length of queue to be %d, got %d", idx+1, queueObj.Length())
		}

		if queueObj.Empty() {
			t.Errorf("Expected queue to not be empty")
		}

		if peeked, ok := queueObj.Peek(); !ok || peeked != fullList[0] {
			t.Errorf("Expected Peek to return %d, got %d", fullList[0], peeked)
		}
	}

	if queueObj.Push(123) {
		t.Errorf("Expected Push to return false, got true")
	}

	for idx, value := range fullList {
		if peeked, ok := queueObj.Peek(); !ok || peeked != value {
			t.Errorf("Expected Peek to return %d, got %d", value, peeked)
		}

		if popped, ok := queueObj.Pop(); !ok || popped != value {
			t.Errorf("Expected Pop to return %d, got %d", value, popped)
		}

		if queueObj.Length() != len(fullList)-idx-1 {
			t.Errorf("Expected length of queue to be %d, got %d", idx, queueObj.Length())
		}
	}

	if !queueObj.Push(123) {
		t.Errorf("Expected Push to return false, got true")
	}

	if !queueObj.Push(124) {
		t.Errorf("Expected Push to return false, got true")
	}

	if peeked, ok := queueObj.Peek(); !ok || peeked != 123 {
		t.Errorf("Expected Peek to return %d, got %d", 123, peeked)
	}

	if popped, ok := queueObj.Pop(); !ok || popped != 123 {
		t.Errorf("Expected Pop to return %d, got %d", 123, popped)
	}

	if peeked, ok := queueObj.Peek(); !ok || peeked != 124 {
		t.Errorf("Expected Peek to return %d, got %d", 124, peeked)
	}

	if popped, ok := queueObj.Pop(); !ok || popped != 124 {
		t.Errorf("Expected Pop to return %d, got %d", 124, popped)
	}
}
