package maps_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/maps"
	"github.com/aohorodnyk/stl/slices"
)

func ExampleKeys() {
	mapStr := map[string]int{
		"b": 2,
		"a": 1,
		"c": 3,
	}
	keysStr := maps.Keys(mapStr)
	slices.Sort(keysStr)
	fmt.Println(keysStr)

	mapEmpty := map[string]int{}
	fmt.Println(maps.Keys(mapEmpty))

	mapEmpty = nil
	fmt.Println(maps.Keys(mapEmpty))

	// Output:
	// [a b c]
	// []
	// []
}

func ExampleValues() {
	mapStr := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
	}
	valuesStr := maps.Values(mapStr)
	slices.Sort(valuesStr)
	fmt.Println(valuesStr)

	mapEmpty := map[string]int{}
	fmt.Println(maps.Values(mapEmpty))

	mapEmpty = nil
	fmt.Println(maps.Values(mapEmpty))

	// Output:
	// [1 2 3]
	// []
	// []
}
