package linq

import (
	"github.com/foxesknow/go-echo/data"
	"github.com/foxesknow/go-echo/generic"
)

// Returns the first item in the sequence if there is one
// otherwise return (ZeroValue, false)
func First[T any](stream data.Stream[T]) (item T, found bool) {
	if i := stream.Iterator(); i.MoveNext() {
		return i.Current(), true
	}

	return generic.Zero[T](), false
}

// Returns the first item in a sequence that satisfies a given predicate
// otherwise return (ZeroValue, false)
func FirstWhere[T any](stream data.Stream[T], predicate func(T) bool) (item T, found bool) {

	for i := stream.Iterator(); i.MoveNext(); {
		current := i.Current()
		if predicate(current) {
			return current, true
		}
	}

	return generic.Zero[T](), false
}

// Returns the first item in the sequence, or a default value is the sequence is empty
func FirstOrDefault[T any](stream data.Stream[T], defaultValue T) T {
	if item, found := First(stream); found {
		return item
	}

	return defaultValue
}

// Returns the first item in the sequence that matches a predicate,
// or a default value is the sequence is empty
func FirstOrDefaultWhere[T any](stream data.Stream[T], defaultValue T, predicate func(T) bool) T {
	if item, found := FirstWhere(stream, predicate); found {
		return item
	}

	return defaultValue
}
