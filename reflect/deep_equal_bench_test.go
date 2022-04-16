package reflect_test

import (
	"reflect"
	"testing"

	reflectstl "github.com/aohorodnyk/stl/reflect"
)

func BenchmarkDeepEqualStandardInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(i, i)
	}
}

func BenchmarkDeepEqualStandardMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(map[string]int{"tset": i}, map[string]int{"tset": i})
	}
}

func BenchmarkDeepEqualInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectstl.DeepEqual(i, i)
	}
}

func BenchmarkDeepEqualMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectstl.DeepEqual(map[string]int{"tset": i}, map[string]int{"tset": i})
	}
}

func BenchmarkDeepEqualCmpInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectstl.DeepEqualCmp(true, i, i)
	}
}

func BenchmarkDeepEqualCmpMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectstl.DeepEqualCmp(false, map[string]int{"tset": i}, map[string]int{"tset": i})
	}
}
