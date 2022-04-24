package stack_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections/stack"
)

func ExampleFixed() {
	// Create a stack with a capacity of 3.
	stackFixed := stack.NewFixed[int](3)

	// Push 3 elements onto the stack.
	fmt.Println(stackFixed.Push(1))
	fmt.Println(stackFixed.Push(2))
	fmt.Println(stackFixed.Push(3))

	// Try to push fourth element to the stack, but it will return false, because the stack is full.
	fmt.Println(stackFixed.Push(4))

	// Peek the top element from the stack.
	fmt.Println(stackFixed.Peek())

	// Pop the top element off the stack.
	fmt.Println(stackFixed.Pop())

	// Peek the top element from the stack.
	fmt.Println(stackFixed.Peek())

	// Pop the top element off the stack.
	fmt.Println(stackFixed.Pop())

	// Peek the top element from the stack.
	fmt.Println(stackFixed.Peek())

	// Pop the top element off the stack.
	fmt.Println(stackFixed.Pop())

	// Peek the top element from the stack.
	fmt.Println(stackFixed.Peek())

	// Pop the top element off the stack.
	// The stack should be empty.
	fmt.Println(stackFixed.Pop())

	// Output:
	// true
	// true
	// true
	// false
	// 3 true
	// 3 true
	// 2 true
	// 2 true
	// 1 true
	// 1 true
	// 0 false
	// 0 false
}
