package linq

import "github.com/foxesknow/go-echo/data"

// Skips the specified number of items.
// If the count is less that 1 then no items are skipped.
// This method is implemented by using deferred execution
func Skip[T any](stream data.Streamable[T], count int) data.Streamable[T] {
	// If we're not skipping anything then just return the stream
	if count < 1 {
		return stream
	}

	// If we're skipping everything then we just need the empty stream
	if collection, ok := stream.(data.Collection); ok && count > collection.Count() {
		return data.EmptyStream[T]()
	}

	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			i := stream.GetStream()
			state := 0

			return &data.FunctionStream[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						// Skip the requested number of items
						for j := 0; j < count && i.MoveNext(); j++ {
							// Do nothing
						}

						state = 1
						fallthrough

					case 1:
						moved := i.MoveNext()

						if !moved {
							state = -1
						}

						return moved

					default:
						return false
					}
				},
				OnCurrent: func() T {
					return i.Current()
				},
			}
		},
	}
}

// Skips items while the predicate is true and then returns the rest of the items in the stream.
// This method is implemented by using deferred execution.
func SkipWhile[T any](stream data.Streamable[T], predicate func(T, int) bool) data.Streamable[T] {
	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			i := stream.GetStream()
			state := 0
			index := 0

			return &data.FunctionStream[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						hasNext := false
						for true {
							hasNext = i.MoveNext()
							if !hasNext {
								break
							}

							if !predicate(i.Current(), index) {
								break
							}

							index++
						}

						if hasNext {
							state = 1
							return true
						} else {
							state = -1
							return false
						}

					case 1:
						moved := i.MoveNext()

						if !moved {
							state = -1
						}

						return moved

					default:
						return false
					}
				},
				OnCurrent: func() T {
					return i.Current()
				},
			}
		},
	}
}
