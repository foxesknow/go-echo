package linq

import "github.com/foxesknow/go-echo/data"

// Determines if a value exists in a sequence.
// If found then (true, index) is returned.
// If not found then (false, -1) is returned.
// Note that depending on the underlying sequence the index may not make sense.
// For example, for a slice it is a valid index, but for the keys of a map it is meaningless
func Contains[T comparable](stream data.Stream[T], value T) (found bool, index int) {
	index = 0
	for i := stream.Iterator(); i.MoveNext(); {
		if i.Current() == value {
			return true, index
		}

		index++
	}

	return false, -1
}

// Determines if a value exists in a sequence.
// If found then (true, index) is returned.
// If not found then (false, -1) is returned.
// Note that depending on the underlying sequence the index may not make sense.
// For example, for a slice it is a valid index, but for the keys of a map it is meaningless
func ContainsWhere[T any](stream data.Stream[T], compare func(T) bool) (found bool, index int) {
	index = 0
	for i := stream.Iterator(); i.MoveNext(); {
		if compare(i.Current()) {
			return true, index
		}

		index++
	}

	return false, -1
}
