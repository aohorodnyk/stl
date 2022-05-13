package binflags

import (
	"unsafe"

	"github.com/aohorodnyk/stl/math/bits"
	"github.com/aohorodnyk/stl/types"
)

// Static is a static size implementation of the binflags.
// This is a slice that is allocated to collect flags in bits.
// The max amount of flags can be stored per slice can be calculated by the formula: `slice length` * `bits in byte`.
//
// This implementation can be used whenever you want to have a predictable size and performance.
// Also it can be saved in a database as an array type.
// To initialize the Dynamic type, use make function or usual instatination:
//   flags := make(Static[uint], 3)
// or
//   flags := Static[uint]{0, 0, 0}
type Static[T types.Integer] []T

// IsSet returns true if the flag is set.
// If the flag is not set, then it returns false.
// This function has O(1) complexity.
func (s Static[T]) IsSet(bitNumber uint) bool {
	if len(s) == 0 {
		return false
	}

	size := uint(unsafe.Sizeof(s[0])) * bits.BitsInByte

	byteIdx := int(bitNumber / size)
	if byteIdx >= len(s) {
		// Bit number is out of range.
		return false
	}

	bitIdx := uint8(bitNumber % size)

	return bits.IsSetBit(s[byteIdx], bitIdx)
}

// Set sets the flag.
// This function is idempotent.
// This function has O(1) complexity.
// It returns false if the flag is out of range.
func (s Static[T]) Set(bitNumber uint) bool {
	if len(s) == 0 {
		return false
	}

	size := uint(unsafe.Sizeof(s[0])) * bits.BitsInByte

	byteIdx := int(bitNumber / size)
	if byteIdx >= len(s) {
		// Bit number is out of range.
		return false
	}

	bitIdx := uint8(bitNumber % size)
	s[byteIdx] = bits.SetBit(s[byteIdx], bitIdx)

	return true
}

// Unset unsets the flag.
// This function is idempotent.
// This function has O(1) complexity.
// It returns false if the flag is out of range.
func (s Static[T]) Unset(bitNumber uint) bool {
	if len(s) == 0 {
		return false
	}

	size := uint(unsafe.Sizeof(s[0])) * bits.BitsInByte

	byteIdx := int(bitNumber / size)
	if byteIdx >= len(s) {
		// Bit number is out of range.
		return false
	}

	bitIdx := uint8(bitNumber % size)
	s[byteIdx] = bits.UnsetBit(s[byteIdx], bitIdx)

	return true
}
