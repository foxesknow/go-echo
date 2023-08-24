package data

type Map[K comparable, V any] interface {
	// Checks to see if a key is in the map
	ContainsKey(key K) bool

	// Attempts to remove an item, returning true if the item was removed, otherwise false
	Remove(key K) bool

	// Attempts to get an item
	Get(key K) (item V, found bool)

	// Adds an item, but only if it is not already there
	Add(key K, value V) bool

	// Adds an item if it is not there, otherwise updates the value mapped to the key
	AddOrUpdate(key K, value V)

	// Removes all items from the map
	Clear()

	// Returns the number of items in the map
	Count() int

	// Returns true if the map is empty, otherwise false
	IsEmpty() bool

	// Returns a stream of all keys in the map, in an undefined order
	Keys() Stream[K]

	// Returns a stream of all values in the map, in an undefined order
	Values() Stream[V]

	// Returns a stream of all pairs in the map, in an undefined order
	KeyValuePairs() Stream[KeyValuePair[K, V]]
}
