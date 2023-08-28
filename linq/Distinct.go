package linq

import (
	"github.com/foxesknow/go-echo/data"
)

// Returns distinct items from a sequence.
// Do not make any assumptions about the order of the returned sequence.
// This method is implemented by using deferred execution.
func Distinct[T comparable](stream data.Streamable[T]) data.Streamable[T] {
	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			// Filter any duplicates
			distinct := make(map[T]bool)
			items := make([]T, 0, 16)

			for i := stream.GetStream(); i.MoveNext(); {
				current := i.Current()
				if _, found := distinct[current]; !found {
					items = append(items, current)
					distinct[current] = true
				}
			}

			return data.FromSlice(items).GetStream()
		},
	}
}
