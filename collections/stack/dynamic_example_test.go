package stack_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections/stack"
)

func ExampleDynamic() {
	// Create a stackDynamic stack.
	stackDynamic := stack.NewDynamic[int]()

	// Push 3 elements onto the stack.
	fmt.Println(stackDynamic.Push(1))
	fmt.Println(stackDynamic.Push(2))
	fmt.Println(stackDynamic.Push(3))

	// Peek the top element from the stack.
	fmt.Println(stackDynamic.Peek())

	// Pop the top element off the stack.
	fmt.Println(stackDynamic.Pop())

	// Peek the top element from the stack.
	fmt.Println(stackDynamic.Peek())

	// Pop the top element off the stack.
	fmt.Println(stackDynamic.Pop())

	// Peek the top element from the stack.
	fmt.Println(stackDynamic.Peek())

	// Pop the top element off the stack.
	fmt.Println(stackDynamic.Pop())

	// Peek the top element from the stack.
	fmt.Println(stackDynamic.Peek())

	// Pop the top element off the stack.
	// The stack should be empty.
	fmt.Println(stackDynamic.Pop())

	// Output:
	// true
	// true
	// true
	// 3 true
	// 3 true
	// 2 true
	// 2 true
	// 1 true
	// 1 true
	// 0 false
	// 0 false
}
