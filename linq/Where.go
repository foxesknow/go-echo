package linq

import "github.com/foxesknow/go-echo/data"

// Filters an enumerable based on a predicate.
// Any items that match the predicate will be returned.
// This method is implemented by using deferred execution
func Where[T any](stream data.Streamable[T], predicate func(T) bool) data.Streamable[T] {
	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			done := false
			i := stream.GetStream()

			return &data.FunctionStream[T]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					for i.MoveNext() {
						if predicate(i.Current()) {
							return true
						}
					}

					done = true
					return false
				},
				OnCurrent: func() T {
					return i.Current()
				},
			}
		},
	}
}

// Filters an enumerable based on a predicate.
// Any items that match the predicate will be returned.
// The predicate receives the value to test as well as the index of the item in the source data.
// This method is implemented by using deferred execution
func WhereIndex[T any](stream data.Streamable[T], predicate func(T, int) bool) data.Streamable[T] {
	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			done := false
			i := stream.GetStream()
			index := -1

			return &data.FunctionStream[T]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					for i.MoveNext() {
						index++
						if predicate(i.Current(), index) {
							return true
						}

					}

					done = true
					return false
				},
				OnCurrent: func() T {
					return i.Current()
				},
			}
		},
	}
}
