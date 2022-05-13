package math

import (
	"math"
)

// Mathematical constants.
const (
	E   = math.E
	Pi  = math.Pi
	Phi = math.Phi

	Sqrt2   = math.Sqrt2
	SqrtE   = math.SqrtE
	SqrtPi  = math.SqrtPi
	SqrtPhi = math.SqrtPhi

	Ln2    = math.Ln2
	Log2E  = math.Log2E
	Ln10   = math.Ln10
	Log10E = math.Log10E
)

// Floating-point limit values.
// Max is the largest finite value representable by the type.
// SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32             = math.MaxFloat32
	SmallestNonzeroFloat32 = math.SmallestNonzeroFloat32

	MaxFloat64             = math.MaxFloat64
	SmallestNonzeroFloat64 = math.SmallestNonzeroFloat64
)

// Integer limit values.
const (
	MaxInt    = math.MaxInt
	MinInt    = math.MinInt
	MaxInt8   = math.MaxInt8
	MinInt8   = math.MinInt8
	MaxInt16  = math.MaxInt16
	MinInt16  = math.MinInt16
	MaxInt32  = math.MaxInt32
	MinInt32  = math.MinInt32
	MaxInt64  = math.MaxInt64
	MinInt64  = math.MinInt64
	MaxUint   = math.MaxUint
	MaxUint8  = math.MaxUint8
	MaxUint16 = math.MaxUint16
	MaxUint32 = math.MaxUint32
	MaxUint64 = math.MaxUint64
)
