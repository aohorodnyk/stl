package slices

// Make2D is a functions that returns a 2D slice of type T.
// This is a factory function.
// In case of negative values, it will pass the values to make function and panic.
func Make2D[T any](rows, cols int) [][]T {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}

	return matrix
}

// Make3D is a functions that returns a 3D slice of type T.
// This is a factory function.
// In case of negative values, it will pass the values to make function and panic.
func Make3D[T any](rows, cols, tubes int) [][][]T {
	matrix := make([][][]T, rows)
	for i := range matrix {
		matrix[i] = Make2D[T](cols, tubes)
	}

	return matrix
}

// Make4D is a functions that returns a 4D slice of type T.
// This is a factory function.
// In case of negative values, it will pass the values to make function and panic.
func Make4D[T any](rows, cols, tubes, cells int) [][][][]T {
	matrix := make([][][][]T, rows)
	for i := range matrix {
		matrix[i] = Make3D[T](cols, tubes, cells)
	}

	return matrix
}
