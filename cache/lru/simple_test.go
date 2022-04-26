package lru_test

import (
	"testing"

	"github.com/aohorodnyk/stl/cache"
	"github.com/aohorodnyk/stl/cache/lru"
)

var (
	_ cache.Cache[string, int]               = &lru.Simple[string, int]{}
	_ cache.Cache[string, map[string]string] = &lru.Simple[string, map[string]string]{}
)

func TestSimple(t *testing.T) {
	t.Parallel()

	lru := lru.NewSimple[string, int](3)

	lru.Set("a", 1)
	lru.Set("b", 2)
	lru.Set("c", 3)

	if lru.Length() != 3 {
		t.Errorf("Length() = %d, want %d", lru.Length(), 3)
	}

	if value, ok := lru.Value("a"); !ok || value != 1 {
		t.Errorf("Value(a) = %d, %t, want %d, %t", value, ok, 1, true)
	}

	lru.Set("d", 4)

	if lru.Length() != 3 {
		t.Errorf("Length() = %d, want %d", lru.Length(), 3)
	}

	if value, ok := lru.Value("a"); !ok || value != 1 {
		t.Errorf("Value(a) = %d, %t, want %d, %t", value, ok, 1, true)
	}

	if value, ok := lru.Value("b"); ok || value != 0 {
		t.Errorf("Value(b) = %d, %t, want %d, %t", value, ok, 0, false)
	}
}
