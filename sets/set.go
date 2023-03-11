package sets

import (
	"github.com/aohorodnyk/stl/types"
)

func New[V comparable]() Set[V] {
	return make(Set[V])
}

func NewSize[V comparable](size int) Set[V] {
	return make(Set[V], size)
}

type Set[V comparable] map[V]types.Empty

func (set Set[V]) Contains(entry V) bool {
	_, ok := set[entry]

	return ok
}

func (set Set[V]) Add(entry V) bool {
	if !set.Contains(entry) {
		set[entry] = types.Empty{}

		return true
	}

	return false
}

func (set Set[V]) Remove(entry V) bool {
	if set.Contains(entry) {
		delete(set, entry)

		return true
	}

	return false
}
