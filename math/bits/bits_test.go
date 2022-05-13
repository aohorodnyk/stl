package bits_test

import (
	"math"
	"testing"

	"github.com/aohorodnyk/stl/math/bits"
)

func TestSetBit(t *testing.T) {
	t.Parallel()

	got8 := bits.SetBit(int8(0), 0)
	if got8 != 1 {
		t.Errorf("SetBit(int8(0), 0) = %d, want 1", got8)
	}

	got8 = bits.SetBit(int8(1), 0)
	if got8 != 1 {
		t.Errorf("SetBit(int8(1), 0) = %d, want 2", got8)
	}

	got8 = bits.SetBit(int8(-1), 7)
	if got8 != -1 {
		t.Errorf("SetBit(int8(-1), 7) = %d, want -1", got8)
	}

	got8 = bits.SetBit(int8(-128), 7)
	if got8 != -128 {
		t.Errorf("SetBit(int8(-128), 7) = %d, want -128", got8)
	}

	got8 = bits.SetBit(int8(8), 7)
	if got8 != -120 {
		t.Errorf("SetBit(int8(8), 7) = %d, want -120", got8)
	}

	got8 = bits.SetBit(int8(8), 8)
	if got8 != 8 {
		t.Errorf("SetBit(int8(8), 8) = %d, want 8", got8)
	}

	got8 = bits.SetBit(int8(8), 9)
	if got8 != 8 {
		t.Errorf("SetBit(int8(8), 9) = %d, want 8", got8)
	}

	got8u := bits.SetBit(uint8(0), 0)
	if got8u != 1 {
		t.Errorf("SetBit(uint8(0), 0) = %d, want 1", got8u)
	}

	got8u = bits.SetBit(uint8(1), 0)
	if got8u != 1 {
		t.Errorf("SetBit(uint8(1), 0) = %d, want 2", got8u)
	}

	got8u = bits.SetBit(uint8(0), 7)
	if got8u != 128 {
		t.Errorf("SetBit(uint8(0), 7) = %d, want 128", got8u)
	}

	got8u = bits.SetBit(uint8(0), 8)
	if got8u != 0 {
		t.Errorf("SetBit(uint8(0), 8) = %d, want 0", got8u)
	}

	got8u = bits.SetBit(uint8(255), 7)
	if got8u != 255 {
		t.Errorf("SetBit(uint8(255), 7) = %d, want 255", got8u)
	}

	got64 := bits.SetBit(int64(0), 0)
	if got64 != 1 {
		t.Errorf("SetBit(int64(0), 0) = %d, want 1", got64)
	}

	got64 = bits.SetBit(int64(math.MaxInt64), 62)
	if got64 != math.MaxInt64 {
		t.Errorf("SetBit(int64(math.MaxInt64), 62) = %d, want math.MaxInt64", got64)
	}

	got64 = bits.SetBit(int64(math.MaxInt64), 63)
	if got64 != -1 {
		t.Errorf("SetBit(int64(math.MaxInt64), 63) = %d, want -1", got64)
	}

	got64 = bits.SetBit(int64(5), 63)
	if got64 != -9223372036854775803 {
		t.Errorf("SetBit(int64(math.MaxInt64), 63) = %d, want -9223372036854775803", got64)
	}
}

func TestUnsetBit(t *testing.T) {
	t.Parallel()

	got8 := bits.UnsetBit(int8(0), 0)
	if got8 != 0 {
		t.Errorf("UnsetBit(int8(0), 0) = %d, want 0", got8)
	}

	got8 = bits.UnsetBit(int8(1), 0)
	if got8 != 0 {
		t.Errorf("UnsetBit(int8(1), 0) = %d, want 0", got8)
	}

	got8 = bits.UnsetBit(int8(-1), 7)
	if got8 != 127 {
		t.Errorf("UnsetBit(int8(-1), 7) = %d, want -128", got8)
	}

	got8 = bits.UnsetBit(int8(-121), 4)
	if got8 != -121 {
		t.Errorf("UnsetBit(int8(-121), 4) = %d, want -121", got8)
	}

	got8 = bits.UnsetBit(int8(-125), 7)
	if got8 != 3 {
		t.Errorf("UnsetBit(int8(-121), 7) = %d, want 3", got8)
	}
}

func TestIsSetBit(t *testing.T) {
	t.Parallel()

	got8 := bits.IsSetBit(int8(0), 0)
	if got8 {
		t.Errorf("IsSetBit(int8(0), 0) = %v, want false", got8)
	}

	got8 = bits.IsSetBit(int8(1), 0)
	if !got8 {
		t.Errorf("IsSetBit(int8(1), 0) = %v, want true", got8)
	}

	got8 = bits.IsSetBit(int8(-128), 7)
	if !got8 {
		t.Errorf("IsSetBit(int8(-128), 7) = %v, want true", got8)
	}

	got8 = bits.IsSetBit(int8(0b0100000), 5)
	if !got8 {
		t.Errorf("IsSetBit(int8(0b0100000), 5) = %v, want true", got8)
	}

	got8 = bits.IsSetBit(int8(0b0100100), 2)
	if !got8 {
		t.Errorf("IsSetBit(int8(0b0100100), 2) = %v, want true", got8)
	}

	got8 = bits.IsSetBit(int8(-1), 6)
	if !got8 {
		t.Errorf("IsSetBit(int8(-1), 7) = %v, want true", got8)
	}

	got8 = bits.IsSetBit(int8(-1), 7)
	if !got8 {
		t.Errorf("IsSetBit(int8(-1), 7) = %v, want true", got8)
	}

	got8 = bits.IsSetBit(int8(-1), 8)
	if got8 {
		t.Errorf("IsSetBit(int8(-1), 8) = %v, want false", got8)
	}

	got8u := bits.IsSetBit(uint8(255), 7)
	if !got8u {
		t.Errorf("IsSetBit(uint8(255), 7) = %v, want true", got8u)
	}

	got8u = bits.IsSetBit(uint8(0), 0)
	if got8u {
		t.Errorf("IsSetBit(uint8(0), 0) = %v, want false", got8u)
	}
}
