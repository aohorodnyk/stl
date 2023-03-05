package maps_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/maps"
)

func ExampleEqual() {
	mapStrLeft := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	mapStrRight := map[string]int{
		"b": 2,
		"c": 3,
		"a": 1,
	}
	fmt.Println(maps.Equal(mapStrLeft, mapStrRight))

	mapStrLeft = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	mapStrRight = map[string]int{
		"b": 2,
		"s": 3,
		"a": 1,
	}
	fmt.Println(maps.Equal(mapStrLeft, mapStrRight))

	// Output:
	// true
	// false
}

func ExampleEqualFunc() {
	type customType struct {
		a int
	}

	mapStrLeft := map[string]customType{
		"b": {a: 2},
		"a": {a: 1},
		"c": {a: 3},
	}
	mapStrRight := map[string]customType{
		"c": {a: 3},
		"b": {a: 2},
		"a": {a: 1},
	}
	cmp := func(i string) bool {
		return mapStrLeft[i].a == mapStrRight[i].a
	}
	fmt.Println(maps.EqualFunc(mapStrLeft, mapStrRight, cmp))

	// Output:
	// true
}
