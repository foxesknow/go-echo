package linq

import "github.com/foxesknow/go-echo/data"

func Select[T any, V any](stream data.Stream[T], projection func(T) V) data.Stream[V] {
	return &data.FunctionStream[V]{
		OnIterator: func() data.Iterator[V] {
			done := false
			i := stream.Iterator()

			return &data.FunctionIterator[V]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					if i.MoveNext() {
						return true
					}

					done = true
					return false
				},
				OnCurrent: func() V {
					return projection(i.Current())
				},
			}
		},
	}
}

func SelectIndex[T any, V any](stream data.Stream[T], projection func(T, int) V) data.Stream[V] {
	return &data.FunctionStream[V]{
		OnIterator: func() data.Iterator[V] {
			done := false
			i := stream.Iterator()
			index := -1

			return &data.FunctionIterator[V]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					if i.MoveNext() {
						index++
						return true
					}

					done = true
					return false
				},
				OnCurrent: func() V {
					return projection(i.Current(), index)
				},
			}
		},
	}
}
