package linkedlist_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections/linkedlist"
)

func ExampleLinkedList() {
	// Create a new doubly linked list with int type.
	list := linkedlist.NewDoubly[int]()

	fmt.Println("Length:", list.Length())
	fmt.Println("AddFirst:", list.AddFirst(1))
	fmt.Println("AddFirst:", list.AddFirst(4))
	fmt.Println("AddLast:", list.AddLast(9))
	fmt.Println("AddLast:", list.AddLast(15))
	fmt.Println("AddAt:", list.AddAt(3, 6))

	nodeInt := list.NodeFirst()
	for nodeInt != nil {
		fmt.Println("Value:", nodeInt.Value())

		nodeInt = nodeInt.Next()
	}

	nodeInt = list.NodeLast()
	for nodeInt != nil {
		fmt.Println("Value:", nodeInt.Value())

		nodeInt = nodeInt.Prev()
	}

	fmt.Println("IndexOf:", list.IndexOf(6))

	value, ok := list.PopFirst()
	fmt.Println("PopFirst:", value, ok)

	value, ok = list.PopLast()
	fmt.Println("PopLast:", value, ok)

	value, ok = list.PopAt(1)
	fmt.Println("PopAt:", value, ok)

	fmt.Println("Length:", list.Length())

	nodeInt = list.NodeFirst()
	for nodeInt != nil {
		fmt.Println("Value:", nodeInt.Value())

		nodeInt = nodeInt.Next()
	}

	nodeInt = list.NodeLast()
	for nodeInt != nil {
		fmt.Println("Value:", nodeInt.Value())

		nodeInt = nodeInt.Prev()
	}

	// Output:
	// Length: 0
	// AddFirst: true
	// AddFirst: true
	// AddLast: true
	// AddLast: true
	// AddAt: true
	// Value: 4
	// Value: 1
	// Value: 9
	// Value: 6
	// Value: 15
	// Value: 15
	// Value: 6
	// Value: 9
	// Value: 1
	// Value: 4
	// IndexOf: 3
	// PopFirst: 4 true
	// PopLast: 15 true
	// PopAt: 9 true
	// Length: 2
	// Value: 1
	// Value: 6
	// Value: 6
	// Value: 1
}
