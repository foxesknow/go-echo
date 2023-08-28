package data

// Defines a read/write list
type List[T any] interface {
	// Adds an item to the end of the list
	Add(item T)

	// Attempts to get an item an the specified index
	Get(index int) (item T, err error)

	// Sets an item an an existing index
	Set(index int, item T) error

	// Inserts at item at the specified index
	Insert(index int, item T) error

	// Removes all items from the list
	Clear()

	// Returns the number of items in the stack
	Count() int

	// Returns true if the stack is empty, otherwise false
	IsEmpty() bool

	// Returns a stream to the items in the stack, starting with the top item
	Stream() Streamable[T]

	// Removes the item at the specified index
	RemoveAt(index int) error
}
