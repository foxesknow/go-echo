package linq

import (
	"github.com/foxesknow/go-echo/data"
)

// Picks a value from a sequence.
// If the sequence is empty then returns (zero, false)
// You can use this method to implement min or max
func Pick[T any](stream data.Stream[T], accept func(candidate, current T) bool) (value T, found bool) {
	if i := stream.Iterator(); i.MoveNext() {
		picked := i.Current()

		for i.MoveNext() {
			candidate := i.Current()
			if accept(candidate, picked) {
				picked = candidate
			}
		}

		return picked, true
	}

	var zero T
	return zero, false
}
