package linq

import "github.com/foxesknow/go-echo/collections"

// Determines if all items in the sequence satisfy a predicate.
// If the sequence is empty the true is returned
func All[T any](enumerable collections.Enumerable[T], predicate func(T) bool) bool {
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		if !predicate(e.Current()) {
			return false
		}
	}

	return true
}
