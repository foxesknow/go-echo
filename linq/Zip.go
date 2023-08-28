package linq

import "github.com/foxesknow/go-echo/data"

// Applied a combination function to elements from 2 streams to produce a new value
// which is streamed back
// If the streams are of different lengths then iteration will stop at the end of the shortest stream.
// This method is implemented by using deferred execution
func Zip[T1 any, T2 any, R any](first data.Streamable[T1], second data.Streamable[T2], combine func(T1, T2) R) data.Streamable[R] {
	return &data.FunctionStreamable[R]{
		OnGetStream: func() data.Stream[R] {
			i1 := first.GetStream()
			i2 := second.GetStream()
			done := false
			var current R

			return &data.FunctionStream[R]{
				OnMoveNext: func() bool {
					if done {
						return false
					}

					if i1.MoveNext() && i2.MoveNext() {
						current = combine(i1.Current(), i2.Current())
						return true
					}

					var zero R
					current = zero
					done = true
					return false
				},
				OnCurrent: func() R {
					return current
				},
			}
		},
	}
}
