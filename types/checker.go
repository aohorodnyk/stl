package types

// IsType checks if the given value is of the given type.
func IsType[Type any](value any) bool {
	switch value.(type) {
	case Type:
		return true
	}

	return false
}
