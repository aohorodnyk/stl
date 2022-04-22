// This file copied from golang.org/x/exp/constraints

// Package types defines a set of useful types to be used
// with type parameters.
package types_test

import (
	"github.com/aohorodnyk/stl/types"
)

type (
	testSigned[T types.Signed]     struct{ f T }
	testUnsigned[T types.Unsigned] struct{ f T }
	testInteger[T types.Integer]   struct{ f T }
	testFloat[T types.Float]       struct{ f T }
	testComplex[T types.Complex]   struct{ f T }
	testOrdered[T types.Ordered]   struct{ f T }
)

type (
	myInt    int
	myString string
)

// TestTypes passes if it compiles.
type TestTypes struct {
	_ testSigned[int]
	_ testSigned[myInt]
	_ testSigned[int64]
	_ testUnsigned[uint]
	_ testUnsigned[uintptr]
	_ testInteger[int8]
	_ testInteger[uint8]
	_ testInteger[uintptr]
	_ testFloat[float32]
	_ testComplex[complex64]
	_ testOrdered[int]
	_ testOrdered[float64]
	_ testOrdered[string]
	_ testOrdered[myString]
}
