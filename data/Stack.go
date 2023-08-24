package data

type Stack[T any] interface {
	// Adds an item to the stack
	Push(value T) Stack[T]

	// Attempts to pop an item from the top of the stack
	Pop() (value T, popped bool)

	// Attempts to return the top of the stack without removing it
	Peek() (value T, peeked bool)

	// Removes all items from the stack
	Clear()

	// Returns the number of items in the stack
	Count() int

	// Returns true if the stack is empty, otherwise false
	IsEmpty() bool

	// Returns a stream to the items in the stack, starting with the top item
	Stream() Stream[T]
}
