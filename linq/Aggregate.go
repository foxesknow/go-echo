package linq

import "github.com/foxesknow/go-echo/collections"

// Applies an accumulator over a sequence.
// The seed is used as the inital value and the accumulator function is called once for each item
// to combine the current item in the sequence with the current aggregate value
func Aggregate[T any, ACC any](enumerable collections.Enumerable[T], seed ACC, accumulator func(ACC, T) ACC) ACC {
	for e := enumerable.GetEnumerator(); e.MoveNext(); {
		seed = accumulator(seed, e.Current())
	}

	return seed
}
