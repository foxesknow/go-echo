package linq

import "github.com/foxesknow/go-echo/collections"

// Returns true if there are any items in the sequence
func Any[T any](enumerable collections.Enumerable[T]) bool {
	e := enumerable.GetEnumerator()
	return e.MoveNext()
}

// Returns true if the sequence is not empty and at least one item
// matches the predicate, otherwise false
func AnyWhere[T any](enumerable collections.Enumerable[T], predicate func(T) bool) bool {
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		if predicate(e.Current()) {
			return true
		}
	}

	return false
}
