package linq

import "github.com/foxesknow/go-echo/collections"

// Filters an enumerable based on a predicate.
// Any items that match the predicate will be returned
func Where[T any](enumerable collections.Enumerable[T], predicate func(T) bool) collections.Enumerable[T] {
	return &collections.FunctionEnumerable[T]{
		OnGetEnumerator: func() collections.Enumerator[T] {
			done := false
			e := enumerable.GetEnumerator()

			return &collections.FunctionEnumerator[T]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					for e.MoveNext() {
						if predicate(e.Current()) {
							return true
						}
					}

					done = true
					return false
				},
				OnCurrent: func() T {
					return e.Current()
				},
			}
		},
	}
}

// Filters an enumerable based on a predicate.
// Any items that match the predicate will be returned.
// The predicate receives the value to test as well as the index of the item in the source data.
func WhereIndex[T any](enumerable collections.Enumerable[T], predicate func(T, int) bool) collections.Enumerable[T] {
	return &collections.FunctionEnumerable[T]{
		OnGetEnumerator: func() collections.Enumerator[T] {
			done := false
			e := enumerable.GetEnumerator()
			index := -1

			return &collections.FunctionEnumerator[T]{
				OnMoveNext: func() bool {
					if done {
						return done
					}

					for e.MoveNext() {
						index++
						if predicate(e.Current(), index) {
							return true
						}

					}

					done = true
					return false
				},
				OnCurrent: func() T {
					return e.Current()
				},
			}
		},
	}
}
