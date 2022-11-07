package linq

import "github.com/foxesknow/go-echo/data"

// Converts a sequence to a slice
func ToSlice[T any](stream data.Stream[T]) []T {
	capacity := 8

	// We can pre-allocate all the spce up front
	if collection, ok := stream.(data.Collection); ok {
		capacity = collection.Count()
	}

	slice := make([]T, 0, capacity)
	return ToExistingSlice(stream, slice)
}

// Converts a sequence to a slice and adds it to an exsiting slice
func ToExistingSlice[T any](stream data.Stream[T], slice []T) []T {
	for i := stream.Iterator(); i.MoveNext(); {
		slice = append(slice, i.Current())
	}

	return slice
}
