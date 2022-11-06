package linq

import "github.com/foxesknow/go-echo/data"

// Applies an accumulator over a sequence.
// The seed is used as the inital value and the accumulator function is called once for each item
// to combine the current item in the sequence with the current aggregate value
func ElementAt[T any](stream data.Stream[T], index int) (item T, found bool) {
	if index < 0 {
		return
	}

	count := 0
	for i := stream.Iterator(); i.MoveNext(); count++ {
		if count == index {
			return i.Current(), true
		}
	}

	return
}

func ElementAtOrDefault[T any](stream data.Stream[T], index int, defaultValue T) T {
	if item, found := ElementAt(stream, index); found {
		return item
	}

	return defaultValue
}
