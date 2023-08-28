package linq

import "github.com/foxesknow/go-echo/data"

// Adds an item to the fromt of the sequence
// This method is implemented by using deferred execution.
func Prepend[T any](stream data.Streamable[T], item T) data.Streamable[T] {
	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			var current T
			i := stream.GetStream()
			state := 0

			return &data.FunctionStream[T]{
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
