package collections

import "github.com/foxesknow/go-echo/data"

type ReadOnlyCollection[T any] interface {
	// Returns the number of items in the stack
	Count() int

	// Returns true if the stack is empty, otherwise false
	IsEmpty() bool

	// Returns a stream to the items in the stack, starting with the top item
	Stream() data.Stream[T]
}