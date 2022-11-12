package linq

import (
	"fmt"

	"github.com/foxesknow/go-echo/data"
	"github.com/foxesknow/go-echo/generic"
)

// Returns the only item in a sequence, or an error if there is not exactly one item in the sequence.
// This method is implemented by using deferred execution.
func Single[T any](stream data.Stream[T]) (item T, err error) {
	if collection, ok := stream.(data.Collection); ok {
		if collection.Count() == 0 {
			return generic.Zero[T](), fmt.Errorf("stream is empty")
		} else if collection.Count() > 1 {
			return generic.Zero[T](), fmt.Errorf("stream has more than one item in it")
		}
	}

	i := stream.Iterator()
	if i.MoveNext() {
		item = i.Current()

		if i.MoveNext() {
			return generic.Zero[T](), fmt.Errorf("stream has more than one item in it")
		}

		return item, nil
	} else {
		return generic.Zero[T](), fmt.Errorf("stream is empty")
	}
}

// Returns the only item in a sequece that matches the predicate.
// If the stream is empty, or more than one item matches the predicate than an error is returned.
// This method is implemented by using deferred execution.
func SingleWhere[T any](stream data.Stream[T], predicate func(T) bool) (item T, err error) {
	itemsMatched := 0
	var matchedItem T

	for i := stream.Iterator(); i.MoveNext(); {
		if predicate(i.Current()) {
			if itemsMatched == 1 {
				err = fmt.Errorf("more than one item matched the predicate")
				return
			}

			matchedItem = i.Current()

			itemsMatched++
		}
	}

	if itemsMatched == 0 {
		err = fmt.Errorf("no items matched the predicate")
	} else {
		item = matchedItem
	}

	return
}
