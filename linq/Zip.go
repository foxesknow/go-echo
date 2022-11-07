package linq

import "github.com/foxesknow/go-echo/data"

// Applied a combination function to elements from 2 streams to produce a new value
// which is streamed back
// If the streams are of different lengths then iteration will stop at the end of the shortest stream.
// This method is implemented by using deferred execution
func Zip[T1 any, T2 any, R any](first data.Stream[T1], second data.Stream[T2], combine func(T1, T2) R) data.Stream[R] {
	return &data.FunctionStream[R]{
		OnIterator: func() data.Iterator[R] {
			i1 := first.Iterator()
			i2 := second.Iterator()
			done := false
			var current R

			return &data.FunctionIterator[R]{
				OnMoveNext: func() bool {
					if done {
						return false
					}

					if i1.MoveNext() && i2.MoveNext() {
						current = combine(i1.Current(), i2.Current())
						return true
					}

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
