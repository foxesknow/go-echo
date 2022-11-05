package linq

import "github.com/foxesknow/go-echo/data"

func Concat[T any](lhs, rhs data.Stream[T]) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			state := 0
			l := lhs.Iterator()
			r := rhs.Iterator()

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						if l.MoveNext() {
							return true
						}
						state = 1
						fallthrough

					case 1:
						if r.MoveNext() {
							return true
						}
						state = -1

					default:
						// Do nothing
					}

					return false
				},
				OnCurrent: func() T {
					if state == 0 {
						return l.Current()
					} else if state == 1 {
						return r.Current()
					} else {
						// It's undefined to call this function if there's nothing available
						var value T
						return value
					}
				},
			}
		},
	}
}
