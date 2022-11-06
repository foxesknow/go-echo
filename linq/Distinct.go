package linq

import (
	"github.com/foxesknow/go-echo/data"
)

// Returns distinct items from a sequence.
// Do not make any assumptions about the order of the returned sequence.
func Distinct[T comparable](stream data.Stream[T]) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			// Filter any duplicates
			distinct := make(map[T]bool)
			items := make([]T, 0, 16)

			for i := stream.Iterator(); i.MoveNext(); {
				current := i.Current()
				if _, found := distinct[current]; !found {
					items = append(items, current)
					distinct[current] = true
				}
			}

			return data.FromSlice(items).Iterator()
		},
	}
}
