package linkedlist

// LinkedList is a linked list interface that defines the methods.
// These methods are implemented in Singly and Doubly linked lists.
// All linked lists can be simply replaced without changing the code.
//nolint:varnamelen // It's good name for the generic type.
type LinkedList[T any] interface {
	// NodeFirst returns the first node in the list.
	NodeFirst() Node[T]
	// NodeLast returns the last node in the list.
	NodeLast() Node[T]
	// NodeAt returns a node by the index from the list.
	NodeAt(int) Node[T]
	// ValueFirst returns the first value in the list.
	ValueFirst() (T, bool)
	// ValueLast returns the last value in the list.
	ValueLast() (T, bool)
	// ValueAt returns a value by the index from the list.
	ValueAt(int) (T, bool)
	// IndexOf returns the first index of the value in the list.
	IndexOf(T) int
	// IndexOfLast returns the last index of the value in the list.
	IndexOfLast(T) int
	// Contains returns true if the value is in the list.
	Contains(T) bool
	// Length returns the length of the list.
	Length() int
	// Empty returns true if the list is empty.
	Empty() bool
	// AddAt adds a value to the list at the index.
	AddAt(int, T) bool
	// AddFirst adds a value to the first index of the list.
	AddFirst(T) bool
	// AddLast adds a value to the last index of the list.
	AddLast(T) bool
	// PopFirst removes the first node from the list and returns the value.
	PopFirst() (T, bool)
	// PopLast removes the last node from the list and returns the value.
	PopLast() (T, bool)
	// PopAt removes the node at the index and returns the value.
	PopAt(int) (T, bool)
	// RemoveNode removes the exact node from the list.
	RemoveNode(Node[T]) bool
	// RemoveFirstBy removes the first node that contains the value.
	RemoveFirstBy(T) bool
	// RemoveLastBy removes the last node that contains the value.
	RemoveLastBy(T) bool
	// RemoveAllBy removes all nodes that contains the value.
	RemoveAllBy(T) bool
	// Clear removes all nodes from the list.
	Clear()
}

// Node is a node interface for the LinkedList interface.
// These methods are implemented in Singly and Doubly linked lists.
// All linked lists can be simply replaced without changing the code.
type Node[T any] interface {
	Value() T
	Next() Node[T]
	Prev() Node[T]
}
