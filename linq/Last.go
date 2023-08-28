package linq

import (
	"fmt"

	"github.com/foxesknow/go-echo/data"
)

// Returns the last item in a sequence, or (zero, false) if not found
func Last[T any](stream data.Streamable[T]) (item T, err error) {
	if indexable, ok := stream.(data.IndexableCollection[T]); ok {
		if indexable.Count() == 0 {
			err = makeNoItemsInStream()
			return
		}

		item, _ = indexable.Get(indexable.Count() - 1)
		return
	}

	if collection, ok := stream.(data.Collection); ok && collection.Count() == 0 {
		var zero T
		return zero, nil
	}

	var last T
	gotSomething := false
	for i := stream.GetStream(); i.MoveNext(); {
		last = i.Current()
		gotSomething = true
	}

	if gotSomething {
		return last, nil
	}

	var zero T
	return zero, fmt.Errorf("predicate did not match any items")
}

// Returns the last item in a sequence that matches the predicate
// or (zero, false) if not found
func LastWhere[T any](stream data.Streamable[T], predicate func(T) bool) (item T, err error) {
	var last T
	gotSomething := false

	for i := stream.GetStream(); i.MoveNext(); {
		next := i.Current()
		if predicate(next) {
			last = next
			gotSomething = true
		}
	}

	if gotSomething {
		return last, nil
	}

	var zero T
	return zero, fmt.Errorf("predicate did not match any items")
}

// Returns the last item in a sequence, or the specified
// default if the sequence is empty
func LastOrDefault[T any](stream data.Streamable[T], defaultValue T) T {
	if item, err := Last(stream); err == nil {
		return item
	}

	return defaultValue
}

// Returns the last item in the sequence which matches the specified predicate,
// or the specified default if none is found
func LastOrDefaultWhere[T any](stream data.Streamable[T], defaultValue T, predicate func(T) bool) T {
	if item, err := LastWhere(stream, predicate); err == nil {
		return item
	}

	return defaultValue
}
