package reflect_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/reflect"
)

func ExampleDeepEqual() {
	// Comparable types and matched values.
	fmt.Println(reflect.DeepEqual(1, 1))
	fmt.Println(reflect.DeepEqual(nil, nil))
	fmt.Println(reflect.DeepEqual("asdfewf", "asdfewf"))
	// Non-comparable types and matched values.
	fmt.Println(reflect.DeepEqual(map[string]string{"a": "1"}, map[string]string{"a": "1"}))
	fmt.Println(reflect.DeepEqual(map[string]string{}, map[string]string{}))

	// Mismatch types.
	fmt.Println(reflect.DeepEqual(1, "1"))
	fmt.Println(reflect.DeepEqual(48, "1"))
	fmt.Println(reflect.DeepEqual(1, 1.0))
	fmt.Println(reflect.DeepEqual(map[string]string{"a": "1"}, "1"))
	fmt.Println(reflect.DeepEqual(map[string]string{}, nil))
	fmt.Println(reflect.DeepEqual(nil, map[string]string{}))
	// Mismatch values.
	fmt.Println(reflect.DeepEqual(1, 2))
	fmt.Println(reflect.DeepEqual(1.0, 1.1))
	fmt.Println(reflect.DeepEqual(map[string]string{"a": "1"}, map[string]string{"a": "2"}))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}

func ExampleDeepEqualCmp() {
	// Comparable types and matched values.
	fmt.Println(reflect.DeepEqualCmp(true, 1, 1))
	fmt.Println(reflect.DeepEqualCmp(false, 1, 1))
	fmt.Println(reflect.DeepEqualCmp(true, nil, nil))
	fmt.Println(reflect.DeepEqualCmp(false, nil, nil))
	fmt.Println(reflect.DeepEqualCmp(true, "asdfewf", "asdfewf"))
	fmt.Println(reflect.DeepEqualCmp(false, "asdfewf", "asdfewf"))
	fmt.Println(reflect.DeepEqualCmp(false, map[string]string{"a": "1"}, map[string]string{"a": "1"}))

	fmt.Println(reflect.DeepEqualCmp(true, 1, 1.0))
	fmt.Println(reflect.DeepEqualCmp(false, 1, 1.0))
	fmt.Println(reflect.DeepEqualCmp(true, 1, "1"))
	fmt.Println(reflect.DeepEqualCmp(false, 1, "1"))
	fmt.Println(reflect.DeepEqualCmp(true, 48, "1"))
	fmt.Println(reflect.DeepEqualCmp(false, 48, "1"))
	fmt.Println(reflect.DeepEqualCmp(false, map[string]string{"a": "1"}, map[string]string{"a": "2"}))
	fmt.Println(reflect.DeepEqualCmp(true, map[string]string{"a": "1"}, []string{"a", "b", "c"}))
	fmt.Println(reflect.DeepEqualCmp(true, map[string]string{"a": "1"}, nil))
	fmt.Println(reflect.DeepEqualCmp(false, map[string]string{"a": "1"}, nil))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}
