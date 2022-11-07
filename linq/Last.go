package linq

import (
	"github.com/foxesknow/go-echo/data"
)

// Returns the last item in a sequence, or (zero, false) if not found
func Last[T any](stream data.Stream[T]) (item T, found bool) {
	if collection, ok := stream.(data.Collection); ok && collection.Count() == 0 {
		var zero T
		return zero, false
	}

	var last T
	gotSomething := false
	for i := stream.Iterator(); i.MoveNext(); {
		last = i.Current()
		gotSomething = true
	}

	if gotSomething {
		return last, true
	}

	var zero T
	return zero, false
}

// Returns the last item in a sequence that matches the predicate
// or (zero, false) if not found
func LastWhere[T any](stream data.Stream[T], predicate func(T) bool) (item T, found bool) {
	var last T
	gotSomething := false

	for i := stream.Iterator(); i.MoveNext(); {
		next := i.Current()
		if predicate(next) {
			last = next
			gotSomething = true
		}
	}

	if gotSomething {
		return last, true
	}

	var zero T
	return zero, false
}

// Returns the last item in a sequence, or the specified
// default if the sequence is empty
func LastOrDefault[T any](stream data.Stream[T], defaultValue T) T {
	if item, found := Last(stream); found {
		return item
	}

	return defaultValue
}

// Returns the last item in the sequence which matches the specified predicate,
// or the specified default if none is found
func LastOrDefaultWhere[T any](stream data.Stream[T], defaultValue T, predicate func(T) bool) T {
	if item, found := LastWhere(stream, predicate); found {
		return item
	}

	return defaultValue
}
