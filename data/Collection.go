package data

type Collection interface {
	// The number of items in the underlying collection
	Count() int
}
