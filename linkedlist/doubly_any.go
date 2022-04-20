package linkedlist

import (
	reflectstd "reflect"

	"github.com/aohorodnyk/stl/math"
	"github.com/aohorodnyk/stl/reflect"
)

func NewDoublyAny[T any]() *DoublyAny[T] {
	var tmp T

	return &DoublyAny[T]{
		comparable: reflectstd.TypeOf(tmp).Comparable(),
	}
}

func NewDoublyAnyDeep[T any]() *DoublyAny[T] {
	return &DoublyAny[T]{
		comparable: false,
	}
}

func NewDoublyAnyCmp[T any](cmp func(T, T) bool) *DoublyAny[T] {
	if cmp == nil {
		panic("Comparable function is required")
	}

	return &DoublyAny[T]{
		cmp: cmp,
	}
}

type DoublyAny[T any] struct {
	head       *DoublyNodeAny[T]
	tail       *DoublyNodeAny[T]
	length     int
	cmp        func(T, T) bool
	comparable bool
}

func (d *DoublyAny[T]) NodeFirst() Node[T] {
	return d.NodeAt(0)
}

func (d *DoublyAny[T]) NodeLast() Node[T] {
	return d.NodeAt(math.Max(d.length-1, 0))
}

func (d *DoublyAny[T]) NodeAt(index int) Node[T] {
	return d.nodeAt(index)
}

func (d *DoublyAny[T]) ValueFirst() (T, bool) {
	return d.ValueAt(0)
}

func (d *DoublyAny[T]) ValueLast() (T, bool) {
	return d.ValueAt(math.Max(d.length-1, 0))
}

func (d *DoublyAny[T]) ValueAt(index int) (T, bool) {
	if node := d.nodeAt(index); node != nil {
		return node.Value(), true
	}

	var result T

	return result, false
}

func (d *DoublyAny[T]) IndexOf(value T) int {
	pointer := d.head
	for index := 0; pointer != nil; index++ {
		if d.compare(pointer.value, value) {
			return index
		}

		pointer = pointer.next
	}

	return -1
}

func (d *DoublyAny[T]) IndexOfLast(value T) int {
	result := -1
	pointer := d.head

	for index := 0; pointer != nil; index++ {
		if d.compare(pointer.value, value) {
			result = index
		}

		pointer = pointer.next
	}

	return result
}

func (d *DoublyAny[T]) Contains(value T) bool {
	return d.IndexOf(value) != -1
}

func (d *DoublyAny[T]) Length() int {
	return d.length
}

func (d *DoublyAny[T]) Empty() bool {
	return d.length == 0
}

func (d *DoublyAny[T]) AddFirst(value T) bool {
	return d.AddAt(0, value)
}

func (d *DoublyAny[T]) AddLast(value T) bool {
	return d.AddAt(d.length, value)
}

func (d *DoublyAny[T]) AddAt(index int, value T) bool {
	if index == 0 {
		d.length++
		d.head = &DoublyNodeAny[T]{
			value: value,
			next:  d.head,
		}

		if d.head.next == nil {
			d.tail = d.head
		} else {
			d.head.next.prev = d.head
		}

		return true
	}

	node := d.nodeAt(index - 1)
	if node == nil {
		return false
	}

	d.length++

	node.next = &DoublyNodeAny[T]{
		value: value,
		next:  node.next,
		prev:  node,
	}

	if node.next.next == nil {
		d.tail = node.next
	}

	return true
}

func (d *DoublyAny[T]) PopFirst() (T, bool) {
	return d.PopAt(0)
}

func (d *DoublyAny[T]) PopLast() (T, bool) {
	return d.PopAt(math.Max(d.length-1, 0))
}

func (d *DoublyAny[T]) PopAt(index int) (T, bool) {
	var result T

	node := d.nodeAt(math.Max(index, 0))
	if node != nil {
		result = node.value
		d.removeNode(node)

		return result, true
	}

	return result, false
}

func (d *DoublyAny[T]) RemoveNode(node Node[T]) bool {
	if node == nil {
		return false
	}

	pointer := d.head
	for pointer != nil {
		if pointer == node {
			d.removeNode(pointer)

			return true
		}

		pointer = pointer.next
	}

	return false
}

func (d *DoublyAny[T]) RemoveFirstBy(value T) bool {
	if d.head == nil {
		return false
	}

	pointer := d.head
	for pointer != nil {
		if d.compare(pointer.value, value) {
			d.removeNode(pointer)

			return true
		}

		pointer = pointer.next
	}

	return false
}

func (d *DoublyAny[T]) RemoveLastBy(value T) bool {
	if d.tail == nil {
		return false
	}

	pointer := d.tail
	for pointer != nil {
		if d.compare(pointer.value, value) {
			d.removeNode(pointer)

			return true
		}

		pointer = pointer.prev
	}

	return false
}

func (d *DoublyAny[T]) RemoveAllBy(value T) bool {
	var removed bool

	if d.head == nil {
		return removed
	}

	pointer := d.head
	for pointer != nil {
		if d.compare(pointer.value, value) {
			pointer = d.removeNode(pointer)

			removed = true
		} else {
			pointer = pointer.next
		}
	}

	return removed
}

func (d *DoublyAny[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.length = 0
}

func (d *DoublyAny[T]) removeNode(node *DoublyNodeAny[T]) *DoublyNodeAny[T] {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		d.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		d.tail = node.prev
	}

	d.length--

	d.syncEnds()

	if node.prev != nil {
		return node.prev
	}

	return node.next
}

func (d *DoublyAny[T]) nodeAt(index int) *DoublyNodeAny[T] {
	if index < 0 || index >= d.length {
		return nil
	}

	var node *DoublyNodeAny[T]

	if middle := d.length / 2; index <= middle {
		node = d.head
		for i := 0; i < index && node != nil; i++ {
			node = node.next
		}
	} else {
		node = d.tail
		for i := d.length - 1; i > index && node != nil; i-- {
			node = node.prev
		}
	}

	if node == nil {
		return nil
	}

	return node
}

func (d *DoublyAny[T]) syncEnds() {
	if d.head == nil || d.tail == nil {
		d.head = nil
		d.tail = nil
		d.length = 0
	}
}

func (d *DoublyAny[T]) compare(a, b T) bool {
	if d.cmp != nil {
		return d.cmp(a, b)
	}

	return reflect.DeepEqual(a, b)
}

type DoublyNodeAny[T any] struct {
	value T
	next  *DoublyNodeAny[T]
	prev  *DoublyNodeAny[T]
}

func (d *DoublyNodeAny[T]) Value() T {
	return d.value
}

func (d *DoublyNodeAny[T]) Next() Node[T] {
	if d.next == nil {
		return nil
	}

	return d.next
}

func (d *DoublyNodeAny[T]) Prev() Node[T] {
	if d.prev == nil {
		return nil
	}

	return d.prev
}
