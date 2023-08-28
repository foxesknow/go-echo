package linq

import "github.com/foxesknow/go-echo/data"

// Splits the elements of a sequence into chunks of size at most size.
// Every chunk accept the last will have "size" items in it.
// The last chunk will have the remaining elements, but will never be empty
// This method is implemented by using deferred execution.
func Chunk[T any](stream data.Streamable[T], size int) data.Streamable[[]T] {
	if size < 1 {
		slice := ToSlice(stream)
		return data.FromValues(slice)
	}

	return &data.FunctionStreamable[[]T]{
		OnGetStream: func() data.Stream[[]T] {
			var chunk []T
			done := false
			it := stream.GetStream()

			return &data.FunctionStream[[]T]{
				OnMoveNext: func() bool {
					if done {
						chunk = nil
						return false
					}

					chunk = make([]T, 0, size)
					for i := 0; i < size && it.MoveNext(); i++ {
						chunk = append(chunk, it.Current())
					}

					if len(chunk) == 0 {
						done = true
						chunk = nil
						return false
					}

					return true
				},
				OnCurrent: func() []T {
					return chunk
				},
			}
		},
	}
}
