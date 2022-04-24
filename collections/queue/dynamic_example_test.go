package queue_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections/queue"
)

func ExampleDynamic() {
	// Create a new dynamic queue.
	fixedStack := queue.NewFixed[string](3)

	// Push 3 elements onto the queue.
	fmt.Println(fixedStack.Push("a"))
	fmt.Println(fixedStack.Push("b"))
	fmt.Println(fixedStack.Push("c"))

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Peek())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Pop())

	// Push the new third element to the queue.
	fmt.Println(fixedStack.Push("e"))

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Peek())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Pop())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Peek())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Pop())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Peek())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Pop())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Peek())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Pop())

	// Output:
	// true
	// true
	// true
	// a true
	// a true
	// true
	// b true
	// b true
	// c true
	// c true
	// e true
	// e true
	//  false
	//  false
}
