package linq

import "github.com/foxesknow/go-echo/collections"

func Append[T any](enumerable collections.Enumerable[T], item T) collections.Enumerable[T] {
	return &collections.FunctionEnumerable[T]{
		OnGetEnumerator: func() collections.Enumerator[T] {
			var current T
			done := false
			e := enumerable.GetEnumerator()

			return &collections.FunctionEnumerator[T]{
				OnMoveNext: func() bool {
					if done {
						return false
					}

					if e.MoveNext() {
						current = e.Current()
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
