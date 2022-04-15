package math

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func Max[T Comparable](input ...T) T {
	cmp := func(a, b T) bool {
		return a > b
	}

	return minMax(cmp, input...)
}

func Min[T Comparable](input ...T) T {
	cmp := func(a, b T) bool {
		return a < b
	}

	return minMax(cmp, input...)
}

func minMax[T Comparable](cmp func(a, b T) bool, input ...T) T {
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
