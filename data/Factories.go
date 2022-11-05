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
