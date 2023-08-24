package collections

type ReadOnlyList[T any] interface {
	ReadOnlyCollection[T]

	// Attempts to get an item an the specified index
	Get(index int) (item T, err error)
}
