package maps_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/maps"
)

func ExampleSet() {
	fmt.Println("map[string]int")

	containerStrInt := map[string]int{}
	fmt.Println(maps.Set(containerStrInt, "test", 3))
	fmt.Println(maps.Set(containerStrInt, "test", 12))

	fmt.Println("map[string]string")

	containerStrStr := map[string]string{}
	fmt.Println(maps.Set(containerStrStr, "test", "string1"))
	fmt.Println(maps.Set(containerStrStr, "test", "string2"))

	// Output:
	// map[string]int
	// 0 false
	// 3 true
	// map[string]string
	//  false
	// string1 true
}
