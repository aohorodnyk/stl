package math

// Comparable is a golang type list that are related to math comparison operations.
// All types have ~ prefix to work with custom types that created from the listed types.
// This interface is public to be used in other packages for custom impleentations.
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// MaxMulti is a function that returns the maximum value for the more than 2 values in a row.
// By using this functions we can compare from 0 to any number of values.
// This function returns 0 if there is no values in the input.
// If you surely know that you need to compare only two values, use Max instead.
// MaxMulti function is a little bit slower than Max function.
func MaxMulti[T Comparable](input ...T) T {
	cmp := func(first, second T) bool {
		return first > second
	}

	return CompareMulti(cmp, input...)
}

// MaxMulti is a function that returns the minimum value for the more than 2 values in a row.
// By using this functions we can compare from 0 to any number of values.
// This function returns 0 if there is no values in the input.
// If you surely know that you need to compare only two values, use Min instead.
// MinMulti function is a little bit slower than Min function.
func MinMulti[T Comparable](input ...T) T {
	cmp := func(first, second T) bool {
		return first < second
	}

	return CompareMulti(cmp, input...)
}

// Min is a function that returns the minimum value for the two values.
// This function is blazing fast.
func Min[T Comparable](first, second T) T {
	if first < second {
		return first
	}

	return second
}

// Max is a function that returns the minimum value for the two values.
// This function is blazing fast.
func Max[T Comparable](first, second T) T {
	if first > second {
		return first
	}

	return second
}

// CompareMulti is a function that returns the minimum or maximum value for the 0..any number of values.
// To use this function you must pass a function that returns true if the first value
// is greater or lower than the second value.
// This function can be used to compare custom structures or interfaces.
func CompareMulti[T any](cmp func(T, T) bool, input ...T) T {
	var result T

	if len(input) == 0 {
		return result
	}

	result = input[0]
	for i := 1; i < len(input); i++ {
		if cmp(input[i], result) {
			result = input[i]
		}
	}

	return result
}
