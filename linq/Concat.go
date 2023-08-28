package linq

import "github.com/foxesknow/go-echo/data"

// Concatenates two sequences.
// This method is implemented by using deferred execution.
func Concat[T any](lhs, rhs data.Streamable[T]) data.Streamable[T] {
	// If the left side is empty then it's just the right side...
	if collection, ok := lhs.(data.Collection); ok && collection.Count() == 0 {
		return rhs
	}

	// ...whereas if the right side is empty it's the left side
	if collection, ok := rhs.(data.Collection); ok && collection.Count() == 0 {
		return lhs
	}

	return &data.FunctionStreamable[T]{
		OnGetStream: func() data.Stream[T] {
			state := 0
			l := lhs.GetStream()
			r := rhs.GetStream()

			return &data.FunctionStream[T]{
				OnMoveNext: func() bool {
					switch state {
					case 0:
						if l.MoveNext() {
							return true
						}
						state = 1
						fallthrough

					case 1:
						if r.MoveNext() {
							return true
						}
						state = -1

					default:
						// Do nothing
					}

					return false
				},
				OnCurrent: func() T {
					if state == 0 {
						return l.Current()
					} else if state == 1 {
						return r.Current()
					} else {
						// It's undefined to call this function if there's nothing available
						var value T
						return value
					}
				},
			}
		},
	}
}
