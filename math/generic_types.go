package math

// Comparable is a golang type list that are related to math comparison operations.
// All types have ~ prefix to work with custom types that created from the listed types.
// This interface is public to be used in other packages for custom impleentations.
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}
