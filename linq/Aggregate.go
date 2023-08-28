package linq

import "github.com/foxesknow/go-echo/data"

// Applies an accumulator over a sequence.
// The seed is used as the inital value and the accumulator function is called once for each item
// to combine the current item in the sequence with the current aggregate value.
// This method is implemented by using deferred execution
func Aggregate[T any, ACC any](stream data.Streamable[T], seed ACC, accumulator func(ACC, T) ACC) ACC {
	for i := stream.GetStream(); i.MoveNext(); {
		seed = accumulator(seed, i.Current())
	}

	return seed
}
