package reflect

import (
	"reflect"
)

func DeepEqual(a, b any) bool {
	comparable := reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable()

	return DeepEqualCmp(comparable, a, b)
}

func DeepEqualCmp(comparable bool, a, b any) bool {
	if comparable {
		return a == b
	}

	return reflect.DeepEqual(a, b)
}
