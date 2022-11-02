package linq

import "github.com/foxesknow/go-echo/collections"

// Returns the number of items in the sequence
func Count[T comparable](enumerable collections.Enumerable[T]) int {
	count := 0
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		count++
	}

	return count
}

// Returns the number of items in the sequence that match the given predicate
func CountWhere[T any](enumerable collections.Enumerable[T], predicate func(T) bool) int {
	count := 0
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		if predicate(e.Current()) {
			count++
		}
	}

	return count
}
