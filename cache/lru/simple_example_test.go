package lru_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/cache/lru"
)

func ExampleSimple() {
	// Create an simple LRU with a capacity of 3.
	lru := lru.NewSimple[string, int](3)

	// Fill the cache up to full capacity.
	lru.Set("a", 1)
	lru.Set("b", 2)
	lru.Set("c", 3)

	fmt.Println(lru.Length()) // Print size of the cache.
	// Print value of "a" and update this key to the most recently used.
	// After this operation, "a" will be the most recently used key and "b" will be the least recently used key.
	fmt.Println(lru.Value("a"))

	// Add new element to the cache. This element will remove the least recently used element.
	lru.Set("d", 4)

	fmt.Println(lru.Length())   // Print size of the cache.
	fmt.Println(lru.Peek("a"))  // Print value of "a" but not update this key to the most recently used.
	fmt.Println(lru.Value("b")) // Try to access Value "b" but this key is not in the cache.

	fmt.Println(lru.Contains("b")) // Check if the cache contains "b" key.
	fmt.Println(lru.Contains("a")) // Check if the cache contains "a" key.

	fmt.Println(lru.Pop("a")) // Remove "a" key from the cache and return the value.
	fmt.Println(lru.Pop("a")) // Remove "a" key from the cache and return the value.

	fmt.Println(lru.Length()) // Print size of the cache.

	lru.Remove("d") // Remove "d" key from the cache.

	fmt.Println(lru.Length()) // Print size of the cache.

	lru.Clear() // Clear the cache.

	fmt.Println(lru.Length()) // Print size of the cache.

	// Fill the cache up to full capacity.
	lru.Set("e", 5)
	lru.Set("f", 6)
	lru.Set("g", 7)
	// Add new element to the cache. This element will remove the least recently used element.
	lru.Set("h", 8)

	fmt.Println(lru.Length()) // Print size of the cache.

	// Print all keys in the cache.
	fmt.Println(lru.Peek("e"))
	fmt.Println(lru.Peek("f"))
	fmt.Println(lru.Peek("g"))
	fmt.Println(lru.Peek("h"))

	// Output:
	// 3
	// 1 true
	// 3
	// 1 true
	// 0 false
	// false
	// true
	// 1 true
	// 0 false
	// 2
	// 1
	// 0
	// 3
	// 0 false
	// 6 true
	// 7 true
	// 8 true
}
