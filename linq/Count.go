package linq

import "github.com/foxesknow/go-echo/data"

// Returns the number of items in the sequence
func Count[T any](stream data.Stream[T]) int {
	// Nice and easy!
	if collection, ok := stream.(data.Collection); ok {
		return collection.Count()
	}

	count := 0
	for i := stream.Iterator(); i.MoveNext(); {
		count++
	}

	return count
}

// Returns the number of items in the sequence that match the given predicate
func CountWhere[T any](stream data.Stream[T], predicate func(T) bool) int {
	count := 0
	for i := stream.Iterator(); i.MoveNext(); {
		if predicate(i.Current()) {
			count++
		}
	}

	return count
}
