package linq

import "github.com/foxesknow/go-echo/collections"

// Determines if a value exists in a sequence.
// If found then (true, index) is returned.
// If not found then (false, -1) is returned.
// Note that depending on the underlying sequence the index may not make sense.
// For example, for a slice it is a valid index, but for the keys of a map it is meaningless
func Contains[T comparable](enumerable collections.Enumerable[T], value T) (found bool, index int) {
	index = 0
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		if e.Current() == value {
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
func ContainsWhere[T any](enumerable collections.Enumerable[T], value T, equal func(T, T) bool) (found bool, index int) {
	index = 0
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		if equal(e.Current(), value) {
			return true, index
		}

		index++
	}

	return false, -1
}
