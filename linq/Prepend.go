package linq

import "github.com/foxesknow/go-echo/data"

// Adds an item to the fromt of the sequence
func Prepend[T any](stream data.Stream[T], item T) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			var current T
			i := stream.Iterator()
			state := 0

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						current = item
						state = 1
						return true

					case 1:
						if i.MoveNext() {
							current = i.Current()
							return true
						}

						state = -1

					default:
						break
					}

					return false
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}
