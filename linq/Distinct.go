package linq

import (
	"github.com/foxesknow/go-echo/collections"
)

// Returns distinct items from a sequence.
// Do not make any assumptions about the order of the returned sequence.
func Distinct[T comparable](enumerable collections.Enumerable[T]) collections.Enumerable[T] {
	return &collections.FunctionEnumerable[T]{
		OnGetEnumerator: func() collections.Enumerator[T] {
			// Filter any duplicates
			distinct := make(map[T]bool)
			items := make([]T, 0, 16)

			for e := enumerable.GetEnumerator(); e.MoveNext(); {
				current := e.Current()
				if _, found := distinct[current]; !found {
					items = append(items, current)
					distinct[current] = true
				}
			}

			return collections.EnumerateSlice(items).GetEnumerator()
		},
	}
}
