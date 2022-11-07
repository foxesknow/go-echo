package linq

import "github.com/foxesknow/go-echo/data"

// Returns the specified number of items from a sequence.
// This method is implemented by using deferred execution.
func Take[T any](stream data.Stream[T], count int) data.Stream[T] {
	if collection, ok := stream.(data.Collection); ok && count >= collection.Count() {
		return stream
	}

	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			if count < 1 {
				return &emptyIterator[T]{}
			}

			i := stream.Iterator()
			soFar := 0

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					if soFar < count {
						soFar++
						return i.MoveNext()
					}

					return false
				},
				OnCurrent: func() T {
					return i.Current()
				},
			}
		},
	}
}

// Take data from a stream whilst the predicate evaluates to true
// Once the predicate returns false the remaining items are skipped.
// This method is implemented by using deferred execution.
func TakeWhileIndex[T any](stream data.Stream[T], predicate func(T, int) bool) data.Stream[T] {
	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			i := stream.Iterator()
			done := false
			index := 0

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					if done {
						return false
					}

					if i.MoveNext() {
						if predicate(i.Current(), index) {
							index++
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
