package queue_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections/queue"
)

func ExampleFixed() {
	// Create a queue with a capacity of 3.
	fixedStack := queue.NewFixed[string](3)

	// Push 3 elements onto the queue.
	fmt.Println(fixedStack.Push("a"))
	fmt.Println(fixedStack.Push("b"))
	fmt.Println(fixedStack.Push("c"))

	// Try to push fourth element to the queue, but it will return false, because the queue is full.
	fmt.Println(fixedStack.Push("d"))

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Peek())

	// Peek the top element from the queue.
	fmt.Println(fixedStack.Pop())

	// Push the new third element to the queue, because of one has been already popped.
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
	// false
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
