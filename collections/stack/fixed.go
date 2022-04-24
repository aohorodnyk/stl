package stack

// NewFixed returns a new fixed-size stack for any type T.
// The stack is initialized with the given capacity.
// The capacity must be greater than zero.
// nolint:varnamelen // It's good name for the generic type.
func NewFixed[T any](capacity int) *Fixed[T] {
	if capacity <= 0 {
		panic("capacity must be greater than 0")
	}

	return &Fixed[T]{
		data: make([]T, capacity),
	}
}

// Fixed is a stack implementation based on static slice.
// This stack is not safe for concurrent use.
// This stack is memory efficient and is suitable for use cases where the stack is has.
// It allocates the memory with the specified size during initialization
// and just use it without any additional allocations.
// All operations are performed in O(1) time.
// nolint:varnamelen // It's good name for the generic type.
type Fixed[T any] struct {
	data   []T
	length int
}

// Push adds a new element to the top of the stack.
// The latest pushed element is the first one to be popped.
// Returns true if the element was pushed successfully.
// Push is a O(1) operation.
func (d *Fixed[T]) Push(value T) bool {
	if d.length >= len(d.data) {
		return false
	}

	d.data[d.length] = value

	d.length++

	return true
}

// Pop removes the top element from the stack and returns it.
// If the stack is empty, Pop returns the default value for the type T and false as a second return value.
// Pop is a O(1) operation.
func (d *Fixed[T]) Pop() (T, bool) {
	var result T

	if d.length == 0 {
		return result, false
	}

	d.length--

	result = d.data[d.length]

	return result, true
}

// Peek returns the top element from the stack without removing it.
// If the stack is empty, Peek returns the default value for the type T and false as a second return value.
// Peek is a O(1) operation.
func (d *Fixed[T]) Peek() (T, bool) {
	if d.length == 0 {
		var tmp T

		return tmp, false
	}

	return d.data[d.length-1], true
}

// Empty returns true if the stack is empty.
// Empty is a O(1) operation.
func (d *Fixed[T]) Empty() bool {
	return d.length == 0
}

// Length returns the number of elements in the stack.
// Length is a O(1) operation.
func (d *Fixed[T]) Length() int {
	return d.length
}

// Length returns the number of elements in the stack.
// Length is a O(1) operation.
func (d *Fixed[T]) Clear() {
	d.length = 0
}
