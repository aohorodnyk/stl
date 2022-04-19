package linkedlist

//nolint:varnamelen // It's good name for the generic type.
type LinkedList[T any] interface {
	NodeFirst() Node[T]
	NodeLast() Node[T]
	NodeAt(int) Node[T]
	ValueFirst() (T, bool)
	ValueLast() (T, bool)
	ValueAt(int) (T, bool)
	IndexOf(T) int
	IndexOfLast(T) int
	Contains(T) bool
	Length() int
	Empty() bool
	AddAt(int, T) bool
	AddFirst(T) bool
	AddLast(T) bool
	PopFirst() (T, bool)
	PopLast() (T, bool)
	PopAt(int) (T, bool)
	RemoveNode(Node[T]) bool
	RemoveFirstBy(T) bool
	RemoveLastBy(T) bool
	RemoveAllBy(T) bool
	Clear()
}

type Node[T any] interface {
	Value() T
	Next() Node[T]
	Prev() Node[T]
}
