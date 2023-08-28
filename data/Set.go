package data

type Set[T comparable] interface {
	// Checks to see if a valie is in the set
	Contains(value T) bool

	// Attempts to remove an item, returning true if the item was removed, otherwise false
	Remove(value T) bool

	// Adds an item, but only if it is not already there.
	// Returns true if added, false it not
	Add(value T) bool

	// Removes all items from the map
	Clear()

	// Returns the number of items in the map
	Count() int

	// Returns true if the map is empty, otherwise false
	IsEmpty() bool

	// Modifies the set to contain all items that are in the set and in the other set
	Union(other Streamable[T])

	// Modifies the set so it doesn't contain the items in the other set
	Except(other Streamable[T])

	// Returns a stream of all items in the set
	Stream() Streamable[T]
}
