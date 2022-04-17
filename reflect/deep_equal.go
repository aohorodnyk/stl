package reflect

import (
	"reflect"
)

func DeepEqual(a, b any) bool {
	isCmp := reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable()

	return DeepEqualCmp(isCmp, a, b)
}

func DeepEqualCmp(isCmp bool, a, b any) bool {
	if isCmp {
		return a == b
	}

	return reflect.DeepEqual(a, b)
}
