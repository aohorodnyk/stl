package linkedlist_test

import "github.com/aohorodnyk/stl/collections/linkedlist"

var (
	_ linkedlist.LinkedList[int]               = &linkedlist.SyncAny[int]{}
	_ linkedlist.LinkedList[string]            = &linkedlist.SyncAny[string]{}
	_ linkedlist.LinkedList[map[string]string] = &linkedlist.SyncAny[map[string]string]{}
)
