package linkedlist_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

func runTests(t *testing.T, factory func() linkedlist.LinkedList[int]) {
	t.Run("Test Simple", func(t *testing.T) {
		t.Parallel()

		list := factory()
		simpleTest(t, list)
	})

	t.Run("Test Multiple", func(t *testing.T) {
		t.Parallel()

		list := factory()
		multipleTest(t, list)
	})

	t.Run("Test Remove First", func(t *testing.T) {
		t.Parallel()

		list := factory()
		removeFirst(t, list)
	})

	t.Run("Test Remove Last", func(t *testing.T) {
		t.Parallel()

		list := factory()
		removeLast(t, list)
	})

	t.Run("Test Empty", func(t *testing.T) {
		t.Parallel()

		list := factory()
		emptyTest(t, list)
	})

	t.Run("Test Next", func(t *testing.T) {
		t.Parallel()

		list := factory()
		nextTest(t, list)
	})
}

func multipleTest(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddFirst(5234)
	list.AddFirst(532)
	list.AddFirst(5234)
	list.AddFirst(532)
	list.AddFirst(123)
	list.AddFirst(123)
	list.AddLast(123)
	list.AddLast(532)
	list.AddLast(5234)

	if list.Length() != 9 {
		t.Errorf("Expected length to be 9, got %d", list.Length())
	}

	if val := list.RemoveAllBy(123); val != 3 {
		t.Errorf("Expected to remove 3 items, got %d", val)
	}

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if val := list.RemoveAllBy(123); val != 0 {
		t.Errorf("Expected to remove 0 items, got %d", val)
	}

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if list.IndexOf(532) != 0 {
		t.Errorf("Expected the first index of 532 to be 0, got %d", list.IndexOf(532))
	}

	if list.IndexOfLast(532) != 4 {
		t.Errorf("Expected the last index of 532 to be 4, got %d", list.IndexOfLast(532))
	}

	if !list.RemoveFirstBy(532) {
		t.Errorf("Expected to remove first element with value 523")
	}

	if list.IndexOf(532) != 1 {
		t.Errorf("Expected the first index of 532 to be 1, got %d", list.IndexOf(532))
	}

	if list.IndexOfLast(532) != 3 {
		t.Errorf("Expected the last index of 532 to be 3, got %d", list.IndexOfLast(532))
	}

	if !list.RemoveLastBy(532) {
		t.Errorf("Expected to last element with value 523")
	}

	if list.IndexOf(532) != 1 {
		t.Errorf("Expected the first index of 532 to be 1, got %d", list.IndexOf(532))
	}

	if list.IndexOfLast(532) != 1 {
		t.Errorf("Expected the last index of 532 to be 3, got %d", list.IndexOfLast(532))
	}

	if list.IndexOf(5234) != 0 {
		t.Errorf("Expected the first index of 5234 to be 0, got %d", list.IndexOf(5234))
	}

	node := list.NodeAt(0)
	if node == nil {
		t.Errorf("Expected to find node at index 0")
	}

	if node.Value() != 5234 {
		t.Errorf("Expected the value of the first node to be 5234, got %d", node.Value())
	}

	if !list.RemoveNode(node) {
		t.Errorf("Expected to remove node")
	}

	if list.IndexOf(5234) != 1 {
		t.Errorf("Expected the first index of 5234 to be 1, got %d", list.IndexOf(5234))
	}

	if list.IndexOfLast(5234) != 2 {
		t.Errorf("Expected the first index of 5234 to be 1, got %d", list.IndexOfLast(5234))
	}
}

func removeLast(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(1)
	list.AddLast(2)

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if list.IndexOf(1) != 0 {
		t.Errorf("Expected the first index of 1 to be 0, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 4 {
		t.Errorf("Expected the last index of 1 to be 4, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveLastBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != 0 {
		t.Errorf("Expected the first index of 1 to be 0, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 2 {
		t.Errorf("Expected the last index of 1 to be 2, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveLastBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != 0 {
		t.Errorf("Expected the first index of 1 to be 0, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 0 {
		t.Errorf("Expected the last index of 1 to be 0, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveLastBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != -1 {
		t.Errorf("Expected the first index of 1 to be -1, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != -1 {
		t.Errorf("Expected the last index of 1 to be -1, got %d", list.IndexOfLast(1))
	}

	if list.RemoveLastBy(1) {
		t.Errorf("Expected to not remove an element with value 1, not elements kept")
	}

	for index := list.Length() - 1; index >= 0; index-- {
		if list.IndexOf(2) != 0 {
			t.Errorf("Expected the first index of 2 to be 0, got %d", list.IndexOf(2))
		}

		if list.IndexOfLast(2) != index {
			t.Errorf("Expected the last index of 2 to be %d, got %d", index, list.IndexOfLast(2))
		}

		if !list.RemoveLastBy(2) {
			t.Errorf("Expected to remove last element with value %d", index)
		}
	}

	if list.RemoveLastBy(2) {
		t.Errorf("Expected to not remove an element with value 2, not elements kept")
	}
}

func removeFirst(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(1)
	list.AddLast(2)

	if list.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", list.Length())
	}

	if list.IndexOf(1) != 0 {
		t.Errorf("Expected the first index of 1 to be 0, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 4 {
		t.Errorf("Expected the last index of 1 to be 4, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveFirstBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != 1 {
		t.Errorf("Expected the first index of 1 to be 1, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 3 {
		t.Errorf("Expected the last index of 1 to be 3, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveFirstBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != 2 {
		t.Errorf("Expected the first index of 1 to be 2, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != 2 {
		t.Errorf("Expected the last index of 1 to be 2, got %d", list.IndexOfLast(1))
	}

	if !list.RemoveFirstBy(1) {
		t.Errorf("Expected to remove last element with value 1")
	}

	if list.IndexOf(1) != -1 {
		t.Errorf("Expected the first index of 1 to be -1, got %d", list.IndexOf(1))
	}

	if list.IndexOfLast(1) != -1 {
		t.Errorf("Expected the last index of 1 to be -1, got %d", list.IndexOfLast(1))
	}

	for index := list.Length() - 1; index >= 0; index-- {
		if list.IndexOf(2) != 0 {
			t.Errorf("Expected the first index of 2 to be 0, got %d", list.IndexOf(2))
		}

		if list.IndexOfLast(2) != index {
			t.Errorf("Expected the last index of 2 to be %d, got %d", index, list.IndexOfLast(2))
		}

		if !list.RemoveLastBy(2) {
			t.Errorf("Expected to remove last element with value %d", index)
		}
	}

	if list.RemoveLastBy(2) {
		t.Errorf("Expected to not remove an element with value 2, not elements kept")
	}
}

func simpleTest(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddLast(-1)
	list.AddAt(1, 3)
	list.AddAt(100, 3)

	exp := []int{2, 3, 1, -1}

	if list.Length() != len(exp) {
		t.Errorf("Expected length %d, got %d", len(exp), list.Length())
	}

	for expIdx, expVal := range exp {
		if !list.Contains(expVal) {
			t.Errorf("Contains(%d) function. Expected true, got false", expVal)
		}

		if list.IndexOf(expVal) != expIdx {
			t.Errorf("IndexOf(%d) function. Expected %d, got %d", expVal, expIdx, list.IndexOf(expVal))
		}

		if val, ok := list.ValueAt(expIdx); !ok || val != expVal {
			t.Errorf("ValueAt(%d) function. Expected %d, got %d", expIdx, expVal, val)
		}

		if node := list.NodeAt(expIdx); node.Value() != expVal {
			t.Errorf("NodeAt(%d) function. Expected %d, got %d", expIdx, expVal, node.Value())
		}
	}

	if val, ok := list.ValueFirst(); !ok || val != exp[0] {
		t.Errorf("ValueFirst() function. Expected %d, got %d", exp[0], val)
	}

	if val, ok := list.ValueLast(); !ok || val != exp[len(exp)-1] {
		t.Errorf("ValueLast() function. Expected %d, got %d", exp[len(exp)-1], val)
	}

	if node := list.NodeFirst(); node.Value() != exp[0] {
		t.Errorf("NodeFirst() function. Expected %d, got %d", exp[0], node.Value())
	}

	if node := list.NodeLast(); node.Value() != exp[len(exp)-1] {
		t.Errorf("NodeLast() function. Expected %d, got %d", exp[len(exp)-1], node.Value())
	}

	if val, ok := list.PopLast(); !ok || val != exp[len(exp)-1] {
		t.Errorf("PopLast() function. Expected %d, got %d", exp[len(exp)-1], val)
	}

	if val, ok := list.ValueLast(); !ok || val != exp[len(exp)-2] {
		t.Errorf("ValueLast() function. Expected %d, got %d", exp[len(exp)-2], val)
	}

	if list.Length() != len(exp)-1 {
		t.Errorf("Expected length %d, got %d", len(exp)-1, list.Length())
	}

	if val, ok := list.PopAt(1); !ok || val != exp[1] {
		t.Errorf("PopAt(1) function. Expected %d, got %d", exp[1], val)
	}

	if list.Length() != len(exp)-2 {
		t.Errorf("Expected length %d, got %d", len(exp)-2, list.Length())
	}

	if val, ok := list.ValueAt(1); !ok || val != exp[2] {
		t.Errorf("ValueAt(1) function. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.PopFirst(); !ok || val != exp[0] {
		t.Errorf("PopFirst() function. Expected %d, got %d", exp[0], val)
	}

	if list.Length() != len(exp)-3 {
		t.Errorf("Expected length %d, got %d", len(exp)-3, list.Length())
	}

	if val, ok := list.ValueFirst(); !ok || val != exp[2] {
		t.Errorf("ValueFirst() function. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.ValueLast(); !ok || val != exp[2] {
		t.Errorf("ValueLast() function. Expected %d, got %d", exp[2], val)
	}

	if val, ok := list.PopFirst(); !ok || val != exp[2] {
		t.Errorf("PopFirst() function. Expected %d, got %d", exp[2], val)
	}

	if list.Length() != 0 {
		t.Errorf("Expected length %d, got %d", 0, list.Length())
	}

	if val, ok := list.ValueFirst(); ok || val != 0 {
		t.Errorf("ValueFirst() function. Expected %d, got %d", 0, val)
	}
}

func emptyTest(t *testing.T, list linkedlist.LinkedList[int]) {
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddLast(-1)
	list.AddAt(1, 3)

	if list.Length() != 4 {
		t.Errorf("Expected length %d, got %d", 4, list.Length())
	}

	if list.Empty() {
		t.Errorf("Expected empty list, got non-empty")
	}

	list.Clear()

	if list.Length() != 0 {
		t.Errorf("Expected length 0, got %d", list.Length())
	}

	if !list.Empty() {
		t.Errorf("Expected non-empty list, got empty")
	}

	if list.Contains(1) {
		t.Error("Contains function. Expected false, got true")
	}

	if list.IndexOf(1) != -1 {
		t.Errorf("IndexOf function. Expected -1, got %d", list.IndexOf(1))
	}

	if _, ok := list.ValueFirst(); ok {
		t.Error("ValueFirst function. Expected false, got true")
	}

	if _, ok := list.ValueLast(); ok {
		t.Error("ValueLast function. Expected false, got true")
	}

	if _, ok := list.ValueAt(0); ok {
		t.Error("ValueAt function. Expected false, got true")
	}

	if node := list.NodeFirst(); node == nil {
		t.Error("NodeFirst function. Expected false, got true")
	}

	if node := list.NodeLast(); node == nil {
		t.Error("NodeLast function. Expected false, got true")
	}

	if node := list.NodeAt(0); node == nil {
		t.Error("NodeAt function. Expected false, got true")
	}

	if _, ok := list.PopFirst(); ok {
		t.Error("PopFirst function. Expected false, got true")
	}

	if _, ok := list.PopLast(); ok {
		t.Error("PopLast function. Expected false, got true")
	}

	if _, ok := list.PopAt(0); ok {
		t.Error("PopAt function. Expected false, got true")
	}

	if val := list.RemoveAllBy(1); val != 0 {
		t.Errorf("RemoveAllBy function. Expected 0, got %d", val)
	}

	if list.RemoveFirstBy(1) {
		t.Error("RemoveFirstBy function. Expected false, got true")
	}

	if list.RemoveLastBy(1) {
		t.Error("RemoveLastBy function. Expected false, got true")
	}
}

func nextTest(t *testing.T, ll linkedlist.LinkedList[int]) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, value := range values {
		if !ll.AddAt(idx, value) {
			t.Errorf("AddAt(%d, %d) function. Expected true, got false", idx, value)
		}
	}

	if ll.Length() != len(values) {
		t.Errorf("Expected length %d, got %d", len(values), ll.Length())
	}

	node := ll.NodeFirst()
	if node == nil {
		t.Errorf("NodeFirst function. Expected non-nil, got nil")
	}

	counter := 0
	for node != nil {
		if node.Value() != values[counter] {
			t.Errorf("Node.Value() function. Expected %d, got %d", values[counter], node.Value())
		}

		node = node.Next()
		counter++
	}
}
