package linq

import "github.com/foxesknow/go-echo/data"

// Projects the items in a sequence into a new form
// This method is implemented by using deferred execution.
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

// Projects the items in a sequence into a new form, supplying the index of the item in the sequence
// This method is implemented by using deferred execution.
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

// Projects each item of a Stream[T] to a  Stream[V] and flattens the result into one sequence
// This method is implemented by using deferred execution.
func SelectMany[T any, V any](stream data.Stream[T], selector func(T) data.Stream[V]) data.Stream[V] {
	return &data.FunctionStream[V]{
		OnIterator: func() data.Iterator[V] {
			state := 0
			var current V

			i := stream.Iterator()
			var v data.Iterator[V]

			return &data.FunctionIterator[V]{
				OnMoveNext: func() bool {
					for state != -1 {
						switch state {
						case 0:
							if !i.MoveNext() {
								state = -1
								return false
							}

							state = 1
							v = selector(i.Current()).Iterator()
							fallthrough

						case 1:
							if v.MoveNext() {
								current = v.Current()
								return true
							}

							// We're at the end of the inner sequence so loop around
							// and get the next item from the outer sequence
							state = 0

						default:
							// Do nothing
						}
					}

					return false
				},
				OnCurrent: func() V {
					return current
				},
			}
		},
	}
}
