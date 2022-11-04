package linq

import (
	"github.com/foxesknow/go-echo/collections"
)

// Picks a value from a sequence.
// If the sequence is empty then returns (zero, false)
// You can use this method to implement min or max
func Pick[T any](enumerable collections.Enumerable[T], accept func(candidate, current T) bool) (value T, found bool) {
	e := enumerable.GetEnumerator()

	if e.MoveNext() {
		picked := e.Current()

		for e.MoveNext() {
			candidate := e.Current()
			if accept(candidate, picked) {
				picked = candidate
			}
		}

		return picked, true
	}

	var zero T
	return zero, false
}
