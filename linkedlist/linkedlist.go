package linkedlist

// LinkedList is a linked list interface that defines the methods.
// These methods are implemented in Singly and Doubly linked lists.
// All linked lists can be simply replaced without changing the code.
//nolint:varnamelen // It's good name for the generic type.
type LinkedList[T any] interface {
	// NodeFirst returns the first node in the list.
	// If the list is empty, it returns nil.
	NodeFirst() Node[T]
	// NodeLast returns the last node in the list.
	// If the list is empty, it returns nil.
	NodeLast() Node[T]
	// NodeAt returns a node by the index from the list.
	// If the index is out of range, it returns nil.
	NodeAt(int) Node[T]
	// ValueFirst returns the first value in the list.
	// If the list is empty, it returns false.
	ValueFirst() (T, bool)
	// ValueLast returns the last value in the list.
	// If the list is empty, it returns false.
	ValueLast() (T, bool)
	// ValueAt returns a value by the index from the list.
	// It returns false if the index is out of range.
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
	// It return true if value was added.
	AddAt(int, T) bool
	// AddFirst adds a value to the first index of the list.
	// It return true if value was added.
	// It is the same as AddAt(0, value).
	AddFirst(T) bool
	// AddLast adds a value to the last index of the list.
	// It return true if value was added.
	// It is the same as AddAt(Length(), value).
	AddLast(T) bool
	// PopFirst removes the first node from the list and returns the value.
	// The second parameter will return true if node was removed from the list.
	PopFirst() (T, bool)
	// PopLast removes the last node from the list and returns the value.
	// The second parameter will return true if node was removed from the list.
	PopLast() (T, bool)
	// PopAt removes the node at the index and returns the value.
	// The second parameter will return true if node was removed from the list.
	PopAt(int) (T, bool)
	// RemoveNode removes the exact node from the list.
	// It returns true if the node was found and removed.
	RemoveNode(Node[T]) bool
	// RemoveFirstBy removes the first node that contains the value.
	// It returns true if the value was found and removed.
	RemoveFirstBy(T) bool
	// RemoveLastBy removes the last node that contains the value.
	// It returns true if the value was found and removed.
	RemoveLastBy(T) bool
	// RemoveAllBy removes all nodes that contains the value.
	// It returns the number of removed nodes.
	RemoveAllBy(T) int
	// Clear removes all nodes from the list.
	Clear()
}

// Node is a node interface for the LinkedList interface.
// These methods are implemented in Singly and Doubly linked lists.
// All linked lists can be simply replaced without changing the code.
//nolint:varnamelen // It's good name for the generic type.
type Node[T any] interface {
	// Value returns the value of the node.
	Value() T
	// Next returns the next node.
	// If the node is the last node, it returns nil.
	Next() Node[T]
	// Prev returns the previous node.
	// If the node is the first node, it returns nil.
	// Previous node cannot be implemented for the Singly linked lists.
	Prev() Node[T]
}
