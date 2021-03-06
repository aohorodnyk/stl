package math

import "github.com/aohorodnyk/stl/types"

// MaxMulti is a function that returns the maximum value for the more than 2 values in a row.
// By using this functions we can compare from 0 to any number of values.
// This function returns 0 if there is no values in the input.
// If you surely know that you need to compare only two values, use Max instead.
// MaxMulti function is a little bit slower than Max function.
func MaxMulti[T types.Ordered](input ...T) T {
	cmp := func(first, second T) bool {
		return first > second
	}

	return MinMaxMulti(cmp, input...)
}

// MaxMulti is a function that returns the minimum value for the more than 2 values in a row.
// By using this functions we can compare from 0 to any number of values.
// This function returns 0 if there is no values in the input.
// If you surely know that you need to compare only two values, use Min instead.
// MinMulti function is a little bit slower than Min function.
func MinMulti[T types.Ordered](input ...T) T {
	cmp := func(first, second T) bool {
		return first < second
	}

	return MinMaxMulti(cmp, input...)
}

// Min is a function that returns the minimum value for the two values.
// This function is blazing fast.
func Min[T types.Ordered](first, second T) T {
	if first < second {
		return first
	}

	return second
}

// Max is a function that returns the minimum value for the two values.
// This function is blazing fast.
func Max[T types.Ordered](first, second T) T {
	if first > second {
		return first
	}

	return second
}

// MinMaxMulti is a function that returns the minimum or maximum value for the 0..any number of values.
// To use this function you must pass a function that returns true if the first value
// is greater or lower than the second value.
// This function can be used to compare custom structures or interfaces.
func MinMaxMulti[T any](cmp func(T, T) bool, input ...T) T {
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

// Compare is a function that returns the compartison result of two values.
// It's a spaceship operator '<=>' implementation.
// If the source value is equal to the target value, the function returns 0.
// If the source value is greater than the target value, the function returns 1.
// If the source value is less than the target value, the function returns -1.
func Compare[T types.Ordered](source, target T) int {
	if source < target {
		return -1
	}

	if source > target {
		return 1
	}

	return 0
}
