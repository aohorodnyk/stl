/*
linkedlist package contains a linked list data structure implementation.
This package contains two different implementations of the linked list. Singly and Doubly linked lists.
	The Singly linked list's node contains only 1 link to the next value.
	The Doubly linked list's node contains two links to the next and previous values.

Every linked list has four different implementations:
	1. Just list name - the linked list that supports comparable types
	2. Sync - the Comparable implementation with concurrent access
	3. Func - the linked list that supports any types, but compare them in search by value methods though reflect.DeepEqual method or a custom comparable function can be passed
	4. AnySync - the Any implementation with concurrent access

To the example on how to use the linked list, please refer to the linkedlist_example_test.go file.
Or see https://pkg.go.dev/github.com/aohorodnyk/stl/collections/linkedlist#example-LinkedList, but use needed factory method, that is needed.

Any and Comparable implementations are needed to speed-up the implementation of the linked list's search methods.
The comparable implementation in 16 times faster than the Any implementation in the search methods with comparable types.
Two separate implementations were created, because of there is not posibility to pass "[T any]" to "[T comparable]", for example:
		func cmp[T comparable](a, b T) bool {
				return a == b
		}
		func cmpAny[T any](a, b T) bool {
				if reflect.TypeOf(a).Comparable() {
						return cmp(a, b)
				}

				return reflect.DeepEqual(a, b)
		}

If I'll find a way to pass "[T any]" to "[T comparable]", I'll remove the Comparable implementations.

To create a new linked list, use New functions.
*/
package linkedlist
