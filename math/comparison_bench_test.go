package math_test

import (
	mathstd "math"
	"testing"

	"github.com/aohorodnyk/stl/math"
)

func BenchmarkMinByteStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathstd.Min(float64(byte(i)), float64(byte(i+1)))
	}
}

func BenchmarkMinByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Min(byte(i), byte(i+1))
	}
}

func BenchmarkMinMultiByte2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.MinMulti(byte(i), byte(i+1))
	}
}

func BenchmarkMinByte4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.MinMulti(byte(i), byte(i+1), byte(i+15), byte(i+6))
	}
}

func BenchmarkMinIntStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathstd.Min(float64(i), float64(i+1))
	}
}

func BenchmarkMinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Min(i, i+1)
	}
}

func BenchmarkMinMultiInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.MinMulti(i, i+1)
	}
}

func BenchmarkMinInt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.MinMulti(i, i+1, i+15, i+6)
	}
}

func BenchmarkMinFloat64Standard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathstd.Min(mathstd.MaxFloat64-345, mathstd.SmallestNonzeroFloat64+15)
	}
}

func BenchmarkMinFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Min(mathstd.MaxFloat64-345, mathstd.SmallestNonzeroFloat64+15)
	}
}

func BenchmarkMinMultiFloat642(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.MinMulti(mathstd.MaxFloat64-345, mathstd.SmallestNonzeroFloat64+15)
	}
}

func BenchmarkMinFloat644(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.MinMulti(mathstd.MaxFloat64, mathstd.MaxFloat64-345, mathstd.SmallestNonzeroFloat64+15, mathstd.SmallestNonzeroFloat64+15)
	}
}
