package linq

import "github.com/foxesknow/go-echo/data"

// Determines if all items in the sequence satisfy a predicate.
// If the sequence is empty the true is returned.
// This method is implemented by using deferred execution
func All[T any](stream data.Stream[T], predicate func(T) bool) bool {
	for i := stream.Iterator(); i.MoveNext(); {
		if !predicate(i.Current()) {
			return false
		}
	}

	return true
}
