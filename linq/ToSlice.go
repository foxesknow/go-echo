package linq

import "github.com/foxesknow/go-echo/data"

// Converts a sequence to a slice
func ToSlice[T any](stream data.Stream[T]) []T {
	slice := make([]T, 0, 8)

	for i := stream.Iterator(); i.MoveNext(); {
		slice = append(slice, i.Current())
	}

	return slice
}
