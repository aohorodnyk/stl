package bits

import (
	"github.com/aohorodnyk/stl/types"
)

// BitsInByte collects number of bits in byte.
const BitsInByte = 8

// IsSetBit returns true if bit is set in a container value.
func IsSetBit[T types.Integer](container T, bitNumber uint8) bool {
	var val T = 1 << bitNumber

	return val != 0 && container&val == val
}

// SetBit sets bit in a container value and returns the new value.
func SetBit[T types.Integer](container T, bitNumber uint8) T {
	return container | (1 << bitNumber)
}

// UnsetBit unsets bit in a container value and returns the new value.
func UnsetBit[T types.Integer](container T, bitNumber uint8) T {
	return container &^ (1 << bitNumber)
}
