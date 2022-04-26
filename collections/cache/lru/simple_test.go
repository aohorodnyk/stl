package lru_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/collections/cache"
	"github.com/aohorodnyk/stl/collections/cache/lru"
)

var (
	_ cache.Cache[string, int]               = &lru.Simple[string, int]{}
	_ cache.Cache[string, map[string]string] = &lru.Simple[string, map[string]string]{}
)

func TestSimple(t *testing.T) {
	t.Parallel()

	// Create an simple LRU with a capacity of 3.
	lru := lru.NewSimple[string, int](3)

	// Fill the cache up to full capacity.
	lru.Set("a", 1)
	lru.Set("b", 2)
	lru.Set("c", 3)

	if lru.Length() != 3 {
		t.Errorf("expected length of the cache to be 3, got %d", lru.Length())
	}

	// Print value of "a" and update this key to the most recently used.
	// After this operation, "a" will be the most recently used key and "b" will be the least recently used key.
	if value, ok := lru.Value("a"); !ok || value != 1 {
		t.Errorf(`expected value of "a" to be 1, got %d`, value)
	}

	// Add new element to the cache. This element will remove the least recently used element.
	lru.Set("d", 4)

	if lru.Length() != 3 {
		t.Errorf("expected length of the cache to be 3, got %d", lru.Length())
	}

	if value, ok := lru.Peek("a"); !ok || value != 1 {
		t.Errorf(`expected value of "a" to be 1, got %d`, value)
	}

	if value, ok := lru.Value("b"); ok || value != 0 {
		t.Errorf(`expected value of "b" to be 0, got %d`, value)
	}

	if lru.Contains("b") {
		t.Errorf(`expected cache to not contain key "b"`)
	}

	if !lru.Contains("a") {
		t.Errorf(`expected cache to contain key "a"`)
	}

	if value, ok := lru.Pop("a"); !ok || value != 1 {
		t.Errorf(`expected value of "a" to be 1, got %d`, value)
	}

	if value, ok := lru.Pop("a"); ok || value != 0 {
		t.Errorf(`expected value of "a" to be 0, got %d`, value)
	}

	if lru.Length() != 2 {
		t.Errorf("expected length of the cache to be 2, got %d", lru.Length())
	}

	lru.Remove("d")

	if lru.Length() != 1 {
		t.Errorf("expected length of the cache to be 1, got %d", lru.Length())
	}

	lru.Clear() // Clear the cache.

	if lru.Length() != 0 {
		t.Errorf("expected length of the cache to be 0, got %d", lru.Length())
	}

	// Fill the cache up to full capacity.
	lru.Set("e", 5)
	lru.Set("f", 6)
	lru.Set("g", 7)
	// Add new element to the cache. This element will remove the least recently used element.
	lru.Set("h", 8)

	if lru.Length() != 3 {
		t.Errorf("expected length of the cache to be 3, got %d", lru.Length())
	}

	// Print all keys in the cache.
	if value, ok := lru.Peek("e"); ok || value != 0 {
		t.Errorf(`expected value of "e" to be 0, got %d`, value)
	}

	if value, ok := lru.Peek("f"); !ok || value != 6 {
		t.Errorf(`expected value of "f" to be 6, got %d`, value)
	}

	if value, ok := lru.Peek("g"); !ok || value != 7 {
		t.Errorf(`expected value of "g" to be 7, got %d`, value)
	}

	if value, ok := lru.Peek("h"); !ok || value != 8 {
		t.Errorf(`expected value of "h" to be 8, got %d`, value)
	}
}

func TestSimplePanic(t *testing.T) {
	t.Parallel()

	capacities := []struct {
		capacity int
		panic    bool
	}{
		{-1000, true},
		{-100, true},
		{-10, true},
		{-5, true},
		{-3, true},
		{-1, true},
		{0, true},
		{1, false},
		{3, false},
		{5, false},
		{10, false},
		{100, false},
		{1000, false},
	}

	for idx, prov := range capacities {
		prov := prov

		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r == nil && prov.panic {
					t.Errorf("The code did not panic")
				}

				if r := recover(); r != nil && !prov.panic {
					t.Errorf("The code paniced")
				}
			}()

			lru.NewSimple[string, int](prov.capacity)
		})
	}
}
