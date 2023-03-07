package types_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/types"
)

func ExampleEmpty() {
	printSetStr := func(set map[string]struct{}) {
		fmt.Println(set)
	}

	setStruct := map[string]struct{}{"a": {}, "b": {}}
	printSetStr(setStruct)

	setEmpty := map[string]struct{}{"c": types.Empty{}, "d": types.Empty{}}
	printSetStr(setEmpty)

	// Output:
	// map[a:{} b:{}]
	// map[c:{} d:{}]
}
