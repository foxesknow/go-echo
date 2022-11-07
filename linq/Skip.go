package linq

import "github.com/foxesknow/go-echo/data"

// Skips the specified number of items.
// If the count is less that 1 then no items are skipped
func Skip[T any](stream data.Stream[T], count int) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			i := stream.Iterator()
			state := 0

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						// Skip the requested number of items
						for j := 0; j < count && i.MoveNext(); j++ {
							// Do nothing
						}

						state = 1
						fallthrough

					case 1:
						moved := i.MoveNext()

						if !moved {
							state = -1
						}

						return moved

					default:
						return false
					}
				},
				OnCurrent: func() T {
					return i.Current()
				},
			}
		},
	}
}
