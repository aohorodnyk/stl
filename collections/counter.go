package collections

import (
	"math"
	"sort"

	mathstl "github.com/aohorodnyk/stl/math"
)

// NewCounter creates a new counter with a specified generic type.
func NewCounter[T comparable]() Counter[T] {
	return Counter[T]{
		counter: make(map[T]uint),
	}
}

// NewCounterSize creates a new counter with a specified generic type and predefined size for the initial map.
func NewCounterSize[T comparable](size int) Counter[T] {
	return Counter[T]{
		counter: make(map[T]uint, size),
	}
}

// CounterValue is a type for a counter value.
// This type is used in the `ToSlice()` method.
type CounterValue[T comparable] struct {
	Key   T
	Count uint
}

// Counter is a counter with a specified generic type.
// It is a map with keys and counts per a key.
type Counter[T comparable] struct {
	counter map[T]uint
	total   uint
}

// Add adds counts to the counter's key.
func (c *Counter[T]) Add(key T, cnt uint) {
	if cnt == 0 {
		return
	}

	c.counter[key] += cnt
	c.total += cnt
}

// Sub subtracts counts from the counter's key.
func (c *Counter[T]) Sub(key T, cnt uint) uint {
	curCnt, ok := c.counter[key]
	if !ok || cnt == 0 {
		return 0
	}

	if curCnt <= cnt {
		c.total -= curCnt
		delete(c.counter, key)

		return curCnt
	}

	c.counter[key] -= cnt
	c.total -= cnt

	return cnt
}

// Inc adds 1 to the counter's key.
func (c *Counter[T]) Inc(key T) {
	c.counter[key]++
	c.total++
}

// Dec subtracts 1 from the counter's key.
func (c *Counter[T]) Dec(key T) bool {
	if c.counter[key] == 1 {
		delete(c.counter, key)
		c.total--

		return true
	}

	if c.counter[key] > 0 {
		c.counter[key]--
		c.total--

		return true
	}

	return false
}

// Counted returns the count number of the counter's key.
func (c Counter[T]) Counted(key T) uint {
	return c.counter[key]
}

// Sum returns the sum of all values in the counter.
func (c *Counter[T]) Sum() uint {
	return c.total
}

// Length returns the number of elements in the counter.
func (c *Counter[T]) Length() int {
	return len(c.counter)
}

// Max returns the key with the highest count.
// If there are no values in the counter, it returns nil.
// If there are multiple keys with the same count, the last one is returned.
func (c Counter[T]) Max() (T, uint) {
	var maxKey T

	if len(c.counter) == 0 {
		return maxKey, 0
	}

	var maxValue uint

	for key, val := range c.counter {
		if val > maxValue {
			maxValue = val
			maxKey = key
		}
	}

	return maxKey, maxValue
}

// MaxAll returns the key and value of the maximum value in the counter.
func (c Counter[T]) MaxAll() (map[T]uint, uint) {
	if len(c.counter) == 0 {
		return nil, 0
	}

	_, max := c.Max()

	maxAll := make(map[T]uint)

	for key, val := range c.counter {
		if val == max {
			maxAll[key] = val
		}
	}

	return maxAll, max
}

// Bottom returns the top number of elements in the counter.
// If number is less than or equal zero, it returns nil.
func (c Counter[T]) Top(number int) []CounterValue[T] {
	if number < 0 {
		return nil
	}

	topNum := mathstl.Min(number, len(c.counter))
	top := c.ToSlice()

	return top[len(top)-topNum:]
}

// Bottom returns the bottom number of elements in the counter.
// If number is less than or equal zero, it returns nil.
func (c Counter[T]) Bottom(number int) []CounterValue[T] {
	if number < 0 {
		return nil
	}

	topNum := mathstl.Min(number, len(c.counter))
	top := c.ToSlice()

	return top[:topNum]
}

// Min returns the minimum value in the counter with a key.
// If there are no values in the counter, it returns nil.
// If there are multiple keys with the same count, the last one is returned.
func (c Counter[T]) Min() (T, uint) {
	var minKey T

	if len(c.counter) == 0 {
		return minKey, 0
	}

	var minValue uint = math.MaxUint

	for key, val := range c.counter {
		if val < minValue {
			minValue = val
			minKey = key
		}
	}

	return minKey, minValue
}

// MinAll returns all keys with the minimum value.
func (c Counter[T]) MinAll() (map[T]uint, uint) {
	if len(c.counter) == 0 {
		return nil, 0
	}

	_, min := c.Min()
	minAll := make(map[T]uint)

	for key, val := range c.counter {
		if val == min {
			minAll[key] = val
		}
	}

	return minAll, min
}

// ToSlice returns slice of CounterValues sorted in ascending order by counts.
func (c Counter[T]) ToSlice() []CounterValue[T] {
	if len(c.counter) == 0 {
		return nil
	}

	top := make([]CounterValue[T], 0, len(c.counter))

	for key, val := range c.counter {
		cntVal := CounterValue[T]{key, val}
		top = append(top, cntVal)
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].Count < top[j].Count
	})

	return top
}

// Remove an element from the counter by the given key.
func (c *Counter[T]) Remove(key T) uint {
	curCnt, ok := c.counter[key]
	if !ok {
		return 0
	}

	delete(c.counter, key)
	c.total -= curCnt

	return curCnt
}

// Reset clears all the values in the counter.
func (c *Counter[T]) Reset() {
	c.counter = make(map[T]uint)
	c.total = 0
}
