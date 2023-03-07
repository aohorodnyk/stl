package maps_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/maps"
)

func ExampleKeyExists() {
	fmt.Println(maps.KeyExists(map[string]int{"a": 1}, "b"))
	fmt.Println(maps.KeyExists(map[string]int{"b": 2}, "b"))

	// First parameter can be nil.
	fmt.Println(maps.KeyExists[string, int](nil, "b"))

	// Output:
	// false
	// true
	// false
}

func ExampleSearch() {
	container := map[int]string{1: "one", 10: "ten", 20: "twenty"}

	fmt.Println(maps.Search(container, "ten"))

	// Output:
	// 10 true
}

func ExampleSearchFunc() {
	container := map[string]int{"one": 1, "two": 2, "three": 3}
	searchVal := 2
	cmp := func(key string) bool {
		return container[key] == searchVal
	}

	fmt.Println(maps.SearchFunc(container, cmp))

	// Output:
	// two true
}
