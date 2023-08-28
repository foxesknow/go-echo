package linq

import "github.com/foxesknow/go-echo/data"

// Appends a value to the end of a sequence.
// This method is implemented by using deferred execution.
func Append[T any](stream data.Streamable[T], item T) data.Streamable[T] {
	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			var current T
			done := false
			i := stream.GetStream()

			return &data.FunctionStream[T]{
				OnMoveNext: func() bool {
					if done {
						return false
					}

					if i.MoveNext() {
						current = i.Current()
						return true
					}

					done = true
					current = item
					return true
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}
