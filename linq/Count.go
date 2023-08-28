package linq

import "github.com/foxesknow/go-echo/data"

// Returns the number of items in the sequence
func Count[T any](stream data.Streamable[T]) int {
	// Nice and easy!
	if collection, ok := stream.(data.Collection); ok {
		return collection.Count()
	}

	count := 0
	for i := stream.GetStream(); i.MoveNext(); {
		count++
	}

	return count
}

// Returns the number of items in the sequence that match the given predicate
func CountWhere[T any](stream data.Streamable[T], predicate func(T) bool) int {
	count := 0
	for i := stream.GetStream(); i.MoveNext(); {
		if predicate(i.Current()) {
			count++
		}
	}

	return count
}
