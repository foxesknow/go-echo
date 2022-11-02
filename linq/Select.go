package linq

import "github.com/foxesknow/go-echo/collections"

func Select[T any, V any](enumerable collections.Enumerable[T], projection func(T) V) collections.Enumerable[V] {
	return &collections.FunctionEnumerable[V]{
		OnGetEnumerator: func() collections.Enumerator[V] {
			done := false
			e := enumerable.GetEnumerator()

			return &collections.FunctionEnumerator[V]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					if e.MoveNext() {
						return true
					}

					done = true
					return false
				},
				OnCurrent: func() V {
					return projection(e.Current())
				},
			}
		},
	}
}

func SelectIndex[T any, V any](enumerable collections.Enumerable[T], projection func(T, int) V) collections.Enumerable[V] {
	return &collections.FunctionEnumerable[V]{
		OnGetEnumerator: func() collections.Enumerator[V] {
			done := false
			e := enumerable.GetEnumerator()
			index := -1

			return &collections.FunctionEnumerator[V]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					if e.MoveNext() {
						index++
						return true
					}

					done = true
					return false
				},
				OnCurrent: func() V {
					return projection(e.Current(), index)
				},
			}
		},
	}
}
