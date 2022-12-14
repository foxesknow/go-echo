package collections

import (
	"github.com/foxesknow/go-echo/data"
)

type defaultStack[T any] struct {
	head  *stackNode[T]
	count int
}

type stackNode[T any] struct {
	Value T
	Next  *stackNode[T]
}

// Creates a new stack that is not thread safe
func NewStack[T any]() Stack[T] {
	return &defaultStack[T]{
		head:  nil,
		count: 0,
	}
}

// Removes all items from the stack
func (self *defaultStack[T]) Clear() {
	self.head = nil
	self.count = 0
}

// Pushes a new item onto the top of the stack
func (self *defaultStack[T]) Push(value T) Stack[T] {
	node := &stackNode[T]{
		Value: value,
		Next:  self.head,
	}

	self.head = node
	self.count++

	return self
}

// Attempts to remove on item from the top of the stack
func (self *defaultStack[T]) Pop() (value T, popped bool) {
	if self.count != 0 {
		oldHead := self.head

		self.head = oldHead.Next
		self.count--

		return oldHead.Value, true
	}

	return value, false
}

// Attempts to return the top of the stack without removing it
func (self *defaultStack[T]) Peek() (value T, peeked bool) {
	if self.count != 0 {
		return self.head.Value, true
	}

	return value, false
}

// Returns the number of items in the ctack
func (self *defaultStack[T]) Count() int {
	return self.count
}

// Returns true if the stack is empty, otherwise false
func (self *defaultStack[T]) IsEmpty() bool {
	return self.count == 0
}

func (self *defaultStack[T]) Stream() data.Stream[T] {
	if self.count == 0 {
		return data.EmptyStream[T]()
	}

	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			head := self.head
			var current T

			return &data.FunctionIterator[T]{
				OnMoveNext: func() bool {
					if head != nil {
						current = head.Value
						head = head.Next

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
