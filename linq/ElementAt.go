package linq

import (
	"github.com/foxesknow/go-echo/data"
)

// Returns the item at the specified index within a sequence
// Returns (item, true) if found.
// If the item is not found, or index is less than zero then (zero, false) is returned
func ElementAt[T any](stream data.Stream[T], index int) (item T, err error) {
	if index < 0 {
		err = makeInvalidIndex(index)
		return
	}

	if indexable, ok := stream.(data.IndexableCollection[T]); ok {
		item, err = indexable.Get(index)
		return
	}

	count := 0
	for i := stream.Iterator(); i.MoveNext(); count++ {
		if count == index {
			return i.Current(), nil
		}
	}

	err = makeInvalidIndex(index)
	return
}

// Returns the item at the specified index within a sequence.
// If the index does not exist, or is invalid, then the default value is returned.
func ElementAtOrDefault[T any](stream data.Stream[T], index int, defaultValue T) T {
	if item, err := ElementAt(stream, index); err == nil {
		return item
	}

	return defaultValue
}
