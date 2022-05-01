package cache_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections/cache"
	"github.com/aohorodnyk/stl/collections/cache/lru"
)

var _ cache.Cache[int, int] = &cache.Sync[int, int]{}

func ExampleSync() {
	c := cache.NewSync[string, string](lru.NewSimple[string, string](10))
	c.Set("key", "value")

	value, ok := c.Value("key")
	fmt.Println(value, ok)

	// Output:
	// value true
}
