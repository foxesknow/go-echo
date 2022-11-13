package collections

import (
	"fmt"

	"github.com/foxesknow/go-echo/data"
)

type defaultList[T any] struct {
	items []T
}

func NewList[T any]() List[T] {
	return &defaultList[T]{
		items: make([]T, 0, 8),
	}
}

func (self *defaultList[T]) Count() int {
	return len(self.items)
}

func (self *defaultList[T]) IsEmpty() bool {
	return len(self.items) == 0
}

func (self *defaultList[T]) Clear() {
	self.items = make([]T, 0, 8)
}

func (self *defaultList[T]) Add(item T) {
	self.items = append(self.items, item)
}

func (self *defaultList[T]) Get(index int) (item T, err error) {
	if index >= 0 && index < len(self.items) {
		item = self.items[index]
		err = nil
	} else {
		err = fmt.Errorf("invalid index: %d", index)
	}

	return
}

func (self *defaultList[T]) Set(index int, item T) error {
	if index >= 0 && index < len(self.items) {
		self.items[index] = item
		return nil
	} else {
		return fmt.Errorf("invalid index: %d", index)
	}
}

func (self *defaultList[T]) Insert(index int, item T) error {
	// Inserting at one pass the end if the same as adding
	if index == len(self.items) {
		self.items = append(self.items, item)
		return nil
	}

	if index < 0 || index > len(self.items) {
		return fmt.Errorf("invalid index: %d", index)
	}

	// We'll need to shuffle things around a bit
	var zero T
	self.items = append(self.items, zero)

	// Move everything up be one
	copy(self.items[index+1:], self.items[index:])

	// And now just write over the old item
	self.items[index] = item

	return nil
}

func (self *defaultList[T]) Stream() data.Stream[T] {
	if len(self.items) == 0 {
		return data.EmptyStream[T]()
	}

	return data.FromSlice(self.items)
}
