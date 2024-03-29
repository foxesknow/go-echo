package linq

import (
	"fmt"

	"github.com/foxesknow/go-echo/data"
	"github.com/foxesknow/go-echo/generic"
)

// Returns the first item in the sequence if there is one
// otherwise return (ZeroValue, false)
func First[T any](stream data.Streamable[T]) (item T, err error) {
	if indexable, ok := stream.(data.IndexableCollection[T]); ok {
		if indexable.Count() == 0 {
			err = makeNoItemsInStream()
			return
		}

		item, _ = indexable.Get(0)
		return
	}

	if collection, ok := stream.(data.Collection); ok {
		if collection.Count() == 0 {
			err = makeNoItemsInStream()
			return
		}
	}

	if i := stream.GetStream(); i.MoveNext() {
		return i.Current(), nil
	}

	return generic.Zero[T](), makeNoItemsInStream()
}

// Returns the first item in a sequence that satisfies a given predicate
// otherwise return (ZeroValue, false)
func FirstWhere[T any](stream data.Streamable[T], predicate func(T) bool) (item T, err error) {

	for i := stream.GetStream(); i.MoveNext(); {
		current := i.Current()
		if predicate(current) {
			return current, nil
		}
	}

	return generic.Zero[T](), fmt.Errorf("predicate did not match any items")
}

// Returns the first item in the sequence, or a default value is the sequence is empty
func FirstOrDefault[T any](stream data.Streamable[T], defaultValue T) T {
	if item, err := First(stream); err == nil {
		return item
	}

	return defaultValue
}

// Returns the first item in the sequence that matches a predicate,
// or a default value is the sequence is empty
func FirstOrDefaultWhere[T any](stream data.Streamable[T], defaultValue T, predicate func(T) bool) T {
	if item, err := FirstWhere(stream, predicate); err == nil {
		return item
	}

	return defaultValue
}
