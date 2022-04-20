package reflect

import (
	"reflect"
)

// DeepEqual is the same as golang standard reflect.DeepEqual but with the ability to
// compare comparable types without reflection.
// This small improvements are made to improve performance in places where we need to compare any types.
// At the same time this small improvements improves the performance of the function up to 40%,
// when we can avoid using reflection.
// Unfortunately, theses additional checks decreases the performance for non-comparable types up to 2%.
// This DeepEqual calls DeepEqualCmp(?, a, b) and returns the result.
// If the parameters are not comparable or not will be resolved by the reflection function.
// This function is much faster than golang standard reflect.DeepEqual in a case
// when one of the parameters is comparable.
func DeepEqual(first, second any) bool {
	isNil := first == nil || second == nil
	isCmp := isNil || reflect.TypeOf(first).Comparable() || reflect.TypeOf(second).Comparable()

	return DeepEqualCmp(isCmp, first, second)
}

// DeepEqualCmp is the same as golang standard reflect.DeepEqual but with the ability to
// compare comparable types without reflection.
// This small improvements are made to improve performance in places where we need to compare any types.
// At the same time this small improvements improves the performance of the function up to 40%,
// when we can avoid using reflection.
// Unfortunately, theses additional checks decreases the performance for non-comparable types up to 2%.
// First parameter is a boolean that indicates if the both parameters are comparable.
// If first and second types are not equal the function will always return false,
// even if types are not comparable, but isCmp is true.
// If first and second types are equal, but not comparable and isCmp is true, the function will panic.
// If first and second types are equal, but not comparable and isCmp is false, the function will return false.
func DeepEqualCmp(isCmp bool, first, second any) bool {
	isNil := first == nil || second == nil
	isNotNil := first != nil || second != nil

	if isNil && isNotNil {
		return false
	}

	if isCmp {
		return first == second
	}

	return reflect.DeepEqual(first, second)
}
