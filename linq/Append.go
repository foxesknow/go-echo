package linq

import "github.com/foxesknow/go-echo/data"

// Appends a value to the end of a sequence.
// This method is implemented by using deferred execution.
func Append[T any](stream data.Stream[T], item T) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			var current T
			done := false
			i := stream.Iterator()

			return &data.FunctionIterator[T]{
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
