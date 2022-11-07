package linq

type emptyIterator[T any] struct {
}

func (self *emptyIterator[T]) MoveNext() bool {
	return false
}

func (self *emptyIterator[T]) Current() T {
	var zero T
	return zero
}
