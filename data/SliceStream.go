package data

type sliceStream[T any] struct {
	slice []T
}

type sliceIterator[T any] struct {
	slice []T
	next  int
}

func (self *sliceIterator[T]) MoveNext() bool {
	if self.next+1 < len(self.slice) {
		self.next++
		return true
	}

	return false
}

func (self *sliceIterator[T]) Current() T {
	return self.slice[self.next]
}

func (self *sliceStream[T]) Iterator() Iterator[T] {
	return &sliceIterator[T]{slice: self.slice, next: -1}
}

func (self *sliceStream[T]) Count() int {
	return len(self.slice)
}

// Returns a stream for a slice
func FromSlice[T any](slice []T) Stream[T] {
	return &sliceStream[T]{slice: slice}
}

// Returns a stream for a range of values
func FromValues[T any](values ...T) Stream[T] {
	return &sliceStream[T]{slice: values}
}
