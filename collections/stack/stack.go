package stack

// Stack is a basic LIFO stack interface for any type T.
//nolint:varnamelen // It's good name for the generic type.
type Stack[T any] interface {
	// Push adds a new element to the top of the stack.
	// The latest pushed element is the first one to be popped.
	// Returns true if the element was pushed successfully.
	Push(T) bool
	// Pop removes the top element from the stack and returns it.
	// If the stack is empty, Pop returns the default value for the type T and false as a second return value.
	Pop() (T, bool)
	// Peek returns the top element from the stack without removing it.
	// If the stack is empty, Peek returns the default value for the type T and false as a second return value.
	Peek() (T, bool)
	// Empty returns true if the stack is empty.
	Empty() bool
	// Length returns the number of elements in the stack.
	Length() int
	// Clear removes all elements from the stack.
	Clear()
}
