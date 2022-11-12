package linq

import (
	"fmt"

	"github.com/foxesknow/go-echo/data"
)

// Picks a value from a sequence.
// The first item in the sequence is taken as the initially picked value and then the accept function
// is called for each successive value and the current pick. The function should return which of the
// two values it wishes to use for the current pick
// If the sequence is empty then returns (zero, false)
// You can use this method to implement min or max.
func Pick[T any](stream data.Stream[T], accept func(candidate, current T) bool) (value T, err error) {
	if i := stream.Iterator(); i.MoveNext() {
		picked := i.Current()

		for i.MoveNext() {
			candidate := i.Current()
			if accept(candidate, picked) {
				picked = candidate
			}
		}

		return picked, nil
	}

	var zero T
	return zero, fmt.Errorf("stream was empty")
}
