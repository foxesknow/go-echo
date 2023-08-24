package collections

import (
	"sync"

	"github.com/foxesknow/go-echo/data"
)

type syncStack[T any] struct {
	lock  sync.RWMutex
	head  *syncStackNode[T]
	count int
}

type syncStackNode[T any] struct {
	Value T
	Next  *syncStackNode[T]
}

// Creates a new stack, which is thread safe
func NewSyncStack[T any]() data.Stack[T] {
	return &syncStack[T]{
		lock:  sync.RWMutex{},
		head:  nil,
		count: 0,
	}
}

// Removes all items from the stack
func (self *syncStack[T]) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.head = nil
	self.count = 0
}

// Pushes a new item onto the top of the stack
func (self *syncStack[T]) Push(value T) data.Stack[T] {
	self.lock.Lock()
	defer self.lock.Unlock()

	node := &syncStackNode[T]{
		Value: value,
		Next:  self.head,
	}

	self.head = node
	self.count++

	return self
}

// Attempts to remove on item from the top of the stack
func (self *syncStack[T]) Pop() (value T, popped bool) {
	self.lock.Lock()
	defer self.lock.Unlock()

	if self.count != 0 {
		oldHead := self.head

		self.head = oldHead.Next
		self.count--

		return oldHead.Value, true
	}

	return value, false
}

// Attempts to return the top of the stack without removing it
func (self *syncStack[T]) Peek() (value T, peeked bool) {
	self.lock.RLock()
	defer self.lock.RUnlock()

	if self.count != 0 {
		return self.head.Value, true
	}

	return value, false
}

// Returns the number of items in the ctack
func (self *syncStack[T]) Count() int {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return self.count
}

// Returns true if the stack is empty, otherwise false
func (self *syncStack[T]) IsEmpty() bool {
	return self.Count() == 0
}

// Returns a stream containing the then items on the stack.
// The item on the top of the stack is first in the stream.
// As this is a sync stack the returned items represent a snapshot of the items in the
// stack, and by the time you have iterated over the stack it may have changed.
func (self *syncStack[T]) Stream() data.Stream[T] {
	if self.IsEmpty() {
		return data.EmptyStream[T]()
	}

	return &data.FunctionStream[T]{
		OnIterator: func() data.Iterator[T] {
			head := self.grabHead()
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

func (self *syncStack[T]) grabHead() *syncStackNode[T] {
	self.lock.RLock()
	defer self.lock.RUnlock()

	return self.head
}
