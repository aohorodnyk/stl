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

func ExampleDelete() {
	container := map[string]int{
		"test":  1,
		"test2": 2,
		"test3": 3,
		"test5": 5,
		"test6": 6,
		"test7": 7,
	}
	fmt.Println(container)

	maps.Delete(container, "test2", "test3", "test4")
	fmt.Println(container)

	maps.Delete(container, []string{"test5", "test6"}...)
	fmt.Println(container)

	maps.Delete(container)
	fmt.Println(container)

	// Output:
	// map[test:1 test2:2 test3:3 test5:5 test6:6 test7:7]
	// map[test:1 test5:5 test6:6 test7:7]
	// map[test:1 test7:7]
	// map[test:1 test7:7]
}
