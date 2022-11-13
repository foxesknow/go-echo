package data

// A collection that allows you to access its data via a zero-based index
type IndexableCollection[T any] interface {
	// Attempts to get an item an the specified index
	Get(index int) (item T, err error)

	// Returns the number of items in the collection
	Count() int
}
