package queue

// Queue is a basic FIFO queue interface for any type T.
//nolint:varnamelen // It's good name for the generic type.
type Queue[T any] interface {
	// Push adds a new element to the bottom of the queue.
	// The first pushed element is the first one to be popped.
	// Returns true if the element was pushed successfully.
	Push(T) bool
	// Pop removes the top element from the queue and returns it.
	// If the queue is empty, Pop returns the default value for the type T and false as a second return value.
	Pop() (T, bool)
	// Peek returns the top element from the queue without removing it.
	// If the queue is empty, Peek returns the default value for the type T and false as a second return value.
	Peek() (T, bool)
	// Empty returns true if the queue is empty.
	Empty() bool
	// Length returns the number of elements in the queue.
	Length() int
	// Clear removes all elements from the queue.
	Clear()
}
