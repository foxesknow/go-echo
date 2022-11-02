package linq

import "github.com/foxesknow/go-echo/collections"

func Concat[T any](lhs, rhs collections.Enumerable[T]) collections.Enumerable[T] {
	return &collections.FunctionEnumerable[T]{
		OnGetEnumerator: func() collections.Enumerator[T] {
			state := 0
			l := lhs.GetEnumerator()
			r := rhs.GetEnumerator()

			return &collections.FunctionEnumerator[T]{
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
