package data

import "context"

// Generates a stream of data by calling a generator function
func Generate[T any](generator func() (value T, keepGoing bool)) Streamable[T] {
	return &FunctionStreamable[T]{
		OnGetStream: func() Stream[T] {
			var current T
			keepGoing := true

			return &FunctionStream[T]{
				OnMoveNext: func() bool {
					if keepGoing {
						current, keepGoing = generator()

						if !keepGoing {
							var zero T
							current = zero
						}
					}

					return keepGoing
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}

// Returns "count" numbers, starting from "start" and increasing by 1
// if count is less than 1 then nothing is returned
func Range(start, count int) Streamable[int] {
	return &FunctionStreamable[int]{
		OnGetStream: func() Stream[int] {
			var current int
			next := start
			i := 0

			return &FunctionStream[int]{
				OnMoveNext: func() bool {
					if i < count {
						current = next
						next++
						i++

						return true
					}

					return false
				},
				OnCurrent: func() int {
					return current
				},
			}
		},
	}
}

// Returns a sequence that will yield "item" the specified number of times
func Repeat[T any](item T, count int) Streamable[T] {
	return &FunctionStreamable[T]{
		OnGetStream: func() Stream[T] {
			i := 0

			return &FunctionStream[T]{
				OnMoveNext: func() bool {
					if i < count {
						i++

						return true
					}

					return false
				},
				OnCurrent: func() T {
					return item
				},
			}
		},
	}
}

// Enumerates over data received from a channel until the channel is closed
func FromChannel[T any](channel <-chan T) Streamable[T] {
	return &FunctionStreamable[T]{
		OnGetStream: func() Stream[T] {
			var current T

			return &FunctionStream[T]{
				OnMoveNext: func() bool {
					item, ok := <-channel
					if ok {
						current = item
						return true
					}

					var zero T
					current = zero
					return false
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}

// Creates a stream from a channel and supports cancellation via a context
func FromChannelWithContext[T any](channel <-chan T, cancel context.Context) Streamable[T] {
	return &FunctionStreamable[T]{
		OnGetStream: func() Stream[T] {
			var current T
			var done bool

			return &FunctionStream[T]{
				OnMoveNext: func() bool {
					if done {
						return false
					}

					select {
					case <-cancel.Done():
						var zero T
						current = zero
						done = true
						return false

					case current = <-channel:
						return true
					}
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}

// Enumerates over data received from a channel until the predicate evaluates to false
// or the channel is closed.
func FromChannelWhile[T any](channel <-chan T, predicate func(T) bool) Streamable[T] {
	return &FunctionStreamable[T]{
		OnGetStream: func() Stream[T] {
			var current T

			return &FunctionStream[T]{
				OnMoveNext: func() bool {
					item, ok := <-channel
					if ok {
						if predicate(item) {
							current = item
							return true
						}
					}

					var zero T
					current = zero
					return false
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}
