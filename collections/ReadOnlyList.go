package collections

import (
	"github.com/foxesknow/go-echo/data"
)

type ReadOnlyList[T any] interface {
	// Attempts to get an item an the specified index
	Get(index int) (item T, err error)

	// Returns the number of items in the stack
	Count() int

	// Returns true if the stack is empty, otherwise false
	IsEmpty() bool

	// Returns a stream to the items in the stack, starting with the top item
	Stream() data.Stream[T]
}
