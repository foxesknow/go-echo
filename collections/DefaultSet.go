package collections

import "github.com/foxesknow/go-echo/data"

type defaultSet[T comparable] struct {
	set map[T]bool
}

func NewSet[T comparable]() data.Set[T] {
	return &defaultSet[T]{
		set: make(map[T]bool),
	}
}

func (self *defaultSet[T]) Contains(value T) bool {
	_, found := self.set[value]
	return found
}

func (self *defaultSet[T]) Remove(value T) bool {
	if _, found := self.set[value]; !found {
		return false
	}

	delete(self.set, value)
	return true
}

func (self *defaultSet[T]) Add(value T) bool {
	if _, found := self.set[value]; found {
		return false
	}

	self.set[value] = true
	return true
}

func (self *defaultSet[T]) Count() int {
	return len(self.set)
}

func (self *defaultSet[T]) IsEmpty() bool {
	return len(self.set) == 0
}

func (self *defaultSet[T]) Clear() {
	self.set = make(map[T]bool)
}

func (self *defaultSet[T]) Stream() data.Streamable[T] {
	if len(self.set) == 0 {
		return data.EmptyStream[T]()
	}

	return data.FromMapKeys(self.set)
}

func (self *defaultSet[T]) Union(other data.Streamable[T]) {
	for i := other.GetStream(); i.MoveNext(); {
		self.set[i.Current()] = true
	}
}

func (self *defaultSet[T]) Except(other data.Streamable[T]) {
	for i := other.GetStream(); i.MoveNext(); {
		delete(self.set, i.Current())
	}
}
