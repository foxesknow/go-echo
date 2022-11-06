package data

// Generates a stream of data by calling a generator function
func Generate[T any](generator func() (value T, keepGoing bool)) Stream[T] {
	return &FunctionStream[T]{
		OnIterator: func() Iterator[T] {
			var current T

			return &FunctionIterator[T]{
				OnMoveNext: func() bool {
					var keepGoing bool
					current, keepGoing = generator()
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
func Range(start, count int) Stream[int] {
	return &FunctionStream[int]{
		OnIterator: func() Iterator[int] {
			var current int
			next := start
			i := 0

			return &FunctionIterator[int]{
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

func Repeat[T any](item T, count int) Stream[T] {
	return &FunctionStream[T]{
		OnIterator: func() Iterator[T] {
			i := 0

			return &FunctionIterator[T]{
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
func FromChannel[T any](channel <-chan T) Stream[T] {
	return &FunctionStream[T]{
		OnIterator: func() Iterator[T] {
			var current T

			return &FunctionIterator[T]{
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

// TODO: Add FromChannel overload that take a "stop" channel
