package generic

// Returns the zero value for any type
func Zero[T any]() T {
	var zero T
	return zero
}
