package linq

import "github.com/foxesknow/go-echo/data"

// Returns true if there are any items in the sequence
func Any[T any](stream data.Stream[T]) bool {
	if collection, ok := stream.(data.Collection); ok {
		return collection.Count() != 0
	}

	i := stream.Iterator()
	return i.MoveNext()
}

// Returns true if the sequence is not empty and at least one item
// matches the predicate, otherwise false.
// This method is implemented by using deferred execution.
func AnyWhere[T any](stream data.Stream[T], predicate func(T) bool) bool {
	for i := stream.Iterator(); i.MoveNext(); {
		if predicate(i.Current()) {
			return true
		}
	}

	return false
}
