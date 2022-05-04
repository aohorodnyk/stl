package linkedlist_test

import "github.com/aohorodnyk/stl/collections/linkedlist"

var (
	_ linkedlist.LinkedList[int]    = &linkedlist.Sync[int]{}
	_ linkedlist.LinkedList[string] = &linkedlist.Sync[string]{}
)
