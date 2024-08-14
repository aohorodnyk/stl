package heap

import (
	heapStd "container/heap"
	"sync"

	"github.com/aohorodnyk/stl/types"
)

type heap[V any] struct {
	locker sync.Locker
	data   *StdInterface[V]
}

func (h heap[V]) Fix(idx int) {
	h.locker.Lock()
	defer h.locker.Unlock()

	heapStd.Fix(h.data, idx)
}

func (h heap[V]) Len() int {
	h.locker.Lock()
	defer h.locker.Unlock()

	return h.data.Len()
}

func (h heap[V]) Pop() V {
	h.locker.Lock()
	defer h.locker.Unlock()

	valAny := heapStd.Pop(h.data)
	val, _ := types.Cast[V](valAny)

	return val
}

func (h heap[V]) Push(v V) {
	h.locker.Lock()
	defer h.locker.Unlock()

	heapStd.Push(h.data, v)
}

func (h heap[V]) Remove(i int) V {
	h.locker.Lock()
	defer h.locker.Unlock()

	valAny := heapStd.Remove(h.data, i)
	val, _ := types.Cast[V](valAny)

	return val
}
