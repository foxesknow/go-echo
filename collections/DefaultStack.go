package collections

import (
	"github.com/foxesknow/go-echo/data"
)

type defaultStack[T any] struct {
	data []T
}

// Creates a new stack that is not thread safe
func NewStack[T any]() Stack[T] {
	return &defaultStack[T]{
		data: nil,
	}
}

// Removes all items from the stack
func (self *defaultStack[T]) Clear() {
	self.data = nil
}

// Pushes a new item onto the top of the stack
func (self *defaultStack[T]) Push(value T) Stack[T] {
	self.data = append(self.data, value)
	return self
}

// Attempts to remove on item from the top of the stack
func (self *defaultStack[T]) Pop() (value T, popped bool) {
	length := len(self.data)
	if length != 0 {
		item := self.data[length-1]
		self.data = self.data[0 : length-1]

		return item, true
	}

	return value, false
}

// Attempts to return the top of the stack without removing it
func (self *defaultStack[T]) Peek() (value T, peeked bool) {
	length := len(self.data)
	if length != 0 {
		return self.data[length-1], true
	}

	return value, false
}

// Returns the number of items in the ctack
func (self *defaultStack[T]) Count() int {
	return len(self.data)
}

// Returns true if the stack is empty, otherwise false
func (self *defaultStack[T]) IsEmpty() bool {
	return len(self.data) == 0
}

func (self *defaultStack[T]) Stream() data.Stream[T] {
	if len(self.data) == 0 {
		return data.EmptyStream[T]()
	}

	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			slice := self.data
			index := len(slice) - 1
			var current T

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					if index != -1 {
						current = self.data[index]
						index--

						return true
					}

					var zero T
					current = zero
					return false
				},
				OnCurrent: func() T {
					return current
				},
			}
		},
	}
}
