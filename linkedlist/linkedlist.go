package linkedlist

type LinkedList[T any] interface {
	NodeFirst() (Node[T], bool)
	NodeLast() (Node[T], bool)
	NodeAt(int) (Node[T], bool)
	ValueFirst() (T, bool)
	ValueLast() (T, bool)
	ValueAt(int) (T, bool)
	IndexOf(T) int
	Contains(T) bool
	Length() int
	Empty() bool
	AddAt(int, T) bool
	AddFirst(T) bool
	AddLast(T) bool
	PopFirst() (T, bool)
	PopLast() (T, bool)
	PopAt(int) (T, bool)
	Clear()
}

type Node[T any] interface {
	Value() T
	Next() Node[T]
	Prev() Node[T]
}
