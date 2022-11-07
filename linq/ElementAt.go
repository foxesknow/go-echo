package linq

import "github.com/foxesknow/go-echo/data"

// Returns the item at the specified index within a sequence
// Returns (item, true) if found.
// If the item is not found, or index is less than zero then (zero, false) is returned
func ElementAt[T any](stream data.Stream[T], index int) (item T, found bool) {
	if index < 0 {
		return
	}

	count := 0
	for i := stream.Iterator(); i.MoveNext(); count++ {
		if count == index {
			return i.Current(), true
		}
	}

	return
}

// Returns the item at the specified index within a sequence.
// If the index does not exist, or is invalid, then the default value is returned.
func ElementAtOrDefault[T any](stream data.Stream[T], index int, defaultValue T) T {
	if item, found := ElementAt(stream, index); found {
		return item
	}

	return defaultValue
}
