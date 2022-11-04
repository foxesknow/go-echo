package collections

type sliceEnumerable[T any] struct {
	slice []T
}

type sliceEnumerator[T any] struct {
	slice []T
	next  int
}

func (self *sliceEnumerator[T]) MoveNext() bool {
	if self.next+1 < len(self.slice) {
		self.next++
		return true
	}

	return false
}

func (self *sliceEnumerator[T]) Current() T {
	return self.slice[self.next]
}

func (self *sliceEnumerable[T]) GetEnumerator() Enumerator[T] {
	return &sliceEnumerator[T]{slice: self.slice, next: -1}

}

func EnumerateSlice[T any](slice []T) Enumerable[T] {
	return &sliceEnumerable[T]{slice: slice}
}

func EnumerateValues[T any](values ...T) Enumerable[T] {
	return &sliceEnumerable[T]{slice: values}
}
