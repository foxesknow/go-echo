package linq

import (
	"github.com/foxesknow/go-echo/collections"
	"github.com/foxesknow/go-echo/generic"
)

// Returns the first item in the sequence if there is one
// otherwise return (ZeroValue, false)
func First[T any](enumerable collections.Enumerable[T]) (item T, found bool) {
	e := enumerable.GetEnumerator()

	if e.MoveNext() {
		return e.Current(), true
	}

	return generic.Zero[T](), false
}

// Returns the first item in a sequence that satisfies a given predicate
// otherwise return (ZeroValue, false)
func FirstWhere[T any](enumerable collections.Enumerable[T], predicate func(T) bool) (item T, found bool) {

	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		current := e.Current()
		if predicate(current) {
			return current, true
		}
	}

	return generic.Zero[T](), false
}

// Returns the first item in the sequence, or a default value is the sequence is empty
func FirstOrDefault[T any](enumerable collections.Enumerable[T], defaultValue T) T {
	e := enumerable.GetEnumerator()

	if e.MoveNext() {
		return e.Current()
	}

	return defaultValue
}

// Returns the first item in the sequence that matches a predicate,
// or a default value is the sequence is empty
func FirstOrDefaultWhere[T any](enumerable collections.Enumerable[T], defaultValue T, predicate func(T) bool) T {
	e := enumerable.GetEnumerator()

	if e.MoveNext() {
		current := e.Current()
		if predicate(current) {
			return current
		}
	}

	return defaultValue
}
