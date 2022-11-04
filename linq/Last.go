package linq

import (
	"github.com/foxesknow/go-echo/collections"
)

// Returns the last item in a sequence, or (zero, false) if not found
func Last[T any](enumerable collections.Enumerable[T]) (item T, found bool) {
	var last T
	gotSomething := false
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		last = e.Current()
		gotSomething = true
	}

	if gotSomething {
		return last, true
	}

	var zero T
	return zero, false
}

func LastWhere[T any](enumerable collections.Enumerable[T], predicate func(T) bool) (item T, found bool) {
	var last T
	gotSomething := false

	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		next := e.Current()
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

func LastOrDefault[T any](enumerable collections.Enumerable[T], defaultValue T) T {
	if item, found := Last(enumerable); found {
		return item
	}

	return defaultValue
}

func LastOrDefaultWhere[T any](enumerable collections.Enumerable[T], defaultValue T, predicate func(T) bool) T {
	if item, found := LastWhere(enumerable, predicate); found {
		return item
	}

	return defaultValue
}
