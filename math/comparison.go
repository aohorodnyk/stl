package math

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func MaxMulti[T Comparable](input ...T) T {
	cmp := func(a, b T) bool {
		return a > b
	}

	return minMaxMulti(cmp, input...)
}

func MinMulti[T Comparable](input ...T) T {
	cmp := func(a, b T) bool {
		return a < b
	}

	return minMaxMulti(cmp, input...)
}

func Min[T Comparable](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Max[T Comparable](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func minMaxMulti[T Comparable](cmp func(a, b T) bool, input ...T) T {
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
