package data

import "fmt"

type sliceStreamable[T any] struct {
	slice []T
}

type sliceStream[T any] struct {
	slice []T
	next  int
}

func (self *sliceStream[T]) MoveNext() bool {
	if self.next+1 < len(self.slice) {
		self.next++
		return true
	}

	return false
}

func (self *sliceStream[T]) Current() T {
	return self.slice[self.next]
}

func (self *sliceStreamable[T]) GetStream() Stream[T] {
	return &sliceStream[T]{slice: self.slice, next: -1}
}

func (self *sliceStreamable[T]) Count() int {
	return len(self.slice)
}

func (self *sliceStreamable[T]) Get(index int) (item T, err error) {
	if index >= 0 && index < len(self.slice) {
		return self.slice[index], nil
	}

	err = fmt.Errorf("invalid index: %d", index)
	return
}

// Returns a stream for a slice
func FromSlice[T any](slice []T) Streamable[T] {
	return &sliceStreamable[T]{slice: slice}
}

// Returns a stream for a range of values
func FromValues[T any](values ...T) Streamable[T] {
	return &sliceStreamable[T]{slice: values}
}
