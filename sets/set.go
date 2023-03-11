package sets

import (
	"github.com/aohorodnyk/stl/types"
)

func New[Value comparable]() Set[Value] {
	return make(Set[Value])
}

func NewSize[Value comparable](size int) Set[Value] {
	return make(Set[Value], size)
}

type Set[Value comparable] map[Value]types.Empty

func (set Set[Value]) Contains(entry Value) bool {
	_, ok := set[entry]

	return ok
}

func (set Set[Value]) Add(entry Value) bool {
	if !set.Contains(entry) {
		set[entry] = types.Empty{}

		return true
	}

	return false
}

func (set Set[Value]) Remove(entry Value) bool {
	if set.Contains(entry) {
		delete(set, entry)

		return true
	}

	return false
}
