package cache_test

import "github.com/aohorodnyk/stl/collections/cache"

var _ cache.Cache[int, int] = &cache.Sync[int, int]{}
