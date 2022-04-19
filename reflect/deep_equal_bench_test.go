package reflect_test

import (
	"reflect"
	"testing"

	reflectstl "github.com/aohorodnyk/stl/reflect"
)

func BenchmarkDeepEqualStandardInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nolint:gocritic //This test is expected to have the same params.
		if reflect.DeepEqual(i, i) {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkDeepEqualStandardMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(map[string]int{"test": i}, map[string]int{"test": i}) {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkDeepEqualInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !reflectstl.DeepEqual(i, i) {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkDeepEqualMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !reflectstl.DeepEqual(map[string]int{"test": i}, map[string]int{"test": i}) {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkDeepEqualCmpInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !reflectstl.DeepEqualCmp(true, i, i) {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkEqualInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nolint:gocritic //This test is expected to have the same params.
		if i != i {
			b.Fatal("not equal")
		}
	}
}

func BenchmarkDeepEqualCmpMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !reflectstl.DeepEqualCmp(false, map[string]int{"test": i}, map[string]int{"test": i}) {
			b.Fatal("not equal")
		}
	}
}
