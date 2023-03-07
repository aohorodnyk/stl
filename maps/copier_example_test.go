package maps_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/maps"
)

func ExampleCopy() {
	dest := map[string]int{"three": 3, "one": 1000}
	src := map[string]int{"one": 1, "eleven": 11, "twenty": 20}

	copied := maps.Copy(dest, src)

	fmt.Println(copied)
	fmt.Println(dest)

	// Output:
	// true
	// map[eleven:11 one:1 three:3 twenty:20]
}

func ExampleClone() {
	src := map[string]int{"one": 1, "eleven": 11, "twenty": 20}

	dest := maps.Clone(src)

	fmt.Println(dest)

	// Output:
	// map[eleven:11 one:1 twenty:20]
}
