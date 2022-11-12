package linq

import (
	"sort"

	"github.com/foxesknow/go-echo/data"
)

// Does a stable ordering of the stream
func OrderBy[T any](stream data.Stream[T], less func(lhs, rhs T) bool) data.Stream[T] {
	if collection, ok := stream.(data.Collection); ok {
		if collection.Count() == 0 {
			return data.EmptyStream[T]()
		} else if collection.Count() == 1 {
			// Specical case, it won't need sorting
			return stream
		}
	}

	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			slice := ToSlice(stream)
			sort.SliceStable(slice, func(i, j int) bool {
				return less(slice[i], slice[j])
			})

			next := -1

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					if next+1 < len(slice) {
						next++
						return true
					}

					return false
				},
				OnCurrent: func() T {
					return slice[next]
				},
			}
		},
	}
}
