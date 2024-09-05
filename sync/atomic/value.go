package atomic

import "sync/atomic"

// A Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from [Value.Load].
// Once [Value.Store] has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value[T any] struct {
	std atomic.Value
}

// Load returns the value set by the most recent Store.
// It returns nil if there has been no call to Store for this Value.
func (v *Value[T]) Load() (value *T) {
	valueI := v.std.Load()
	tmp, ok := valueI.(*T)

	if !ok {
		return value
	}

	return tmp
}

// Store sets the value of the [Value] v to val.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value[T]) Store(value *T) {
	v.std.Store(value)
}

// Swap stores new into Value and returns the previous value. It returns nil if
// the Value is empty.
//
// All calls to Swap for a given Value must use values of the same concrete
// type. Swap of an inconsistent type panics, as does Swap(nil).
func (v *Value[T]) Swap(value *T) (old *T) {
	oldI := v.std.Swap(value)
	tmp, ok := oldI.(*T)

	if !ok {
		return old
	}

	return tmp
}

// CompareAndSwap executes the compare-and-swap operation for the [Value].
//
// All calls to CompareAndSwap for a given Value must use values of the same
// concrete type. CompareAndSwap of an inconsistent type panics, as does
// CompareAndSwap(old, nil).
func (v *Value[T]) CompareAndSwap(valueOld, valueNew *T) (swapped bool) {
	return v.std.CompareAndSwap(valueOld, valueNew)
}
