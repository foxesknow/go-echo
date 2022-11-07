package linq

import "github.com/foxesknow/go-echo/data"

// Reverses the order of a sequence in reverse.
// Once you start iterating over the sequence is converted to a slice and iterated over.
// This method is implemented by using deferred execution.
func Reverse[T any](stream data.Stream[T]) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			next := 0
			state := 0
			var slice []T

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						slice = ToSlice(stream)
						next = len(slice)
						state = 1
						fallthrough

					case 1:
						next--

						if next < 0 {
							state = -1
							return false
						}

						return true

					default:
						break
					}

					return false
				},
				OnCurrent: func() T {
					return slice[next]
				},
			}
		},
	}
}
