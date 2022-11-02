package linq

import "github.com/foxesknow/go-echo/collections"

// Converts a sequence to a slice
func ToSlice[T any](enumerable collections.Enumerable[T]) []T {
	slice := make([]T, 0, 8)

	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		slice = append(slice, e.Current())
	}

	return slice
}
