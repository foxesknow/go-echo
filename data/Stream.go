package data

import "fmt"

// Indicates that a type can stream the values it contains
type Streamable[T any] interface {
	// Create an iterator to the values in the stream
	GetStream() Stream[T]
}

type Stream[T any] interface {
	// Moves to the next item, returning true if there is a next item, false if not
	MoveNext() bool

	// Returns the current item. If there is no current item the results are undefined
	Current() T
}

// Allows an stream to be created by calling a factory function
type FunctionStreamable[T any] struct {
	OnGetStream func() Stream[T]
}

func (self *FunctionStreamable[T]) GetStream() Stream[T] {
	return self.OnGetStream()
}

// Defers iterator calls to functions
type FunctionStream[T any] struct {
	OnMoveNext func() bool
	OnCurrent  func() T
}

func (self *FunctionStream[T]) MoveNext() bool {
	return self.OnMoveNext()
}

func (self *FunctionStream[T]) Current() T {
	return self.OnCurrent()
}

type emptyStream[T any] struct {
}

func (self *emptyStream[T]) GetStream() Stream[T] {
	return &emptyIterator[T]{}
}

func (self *emptyStream[T]) Count() int {
	return 0
}

func (self *emptyStream[T]) Get(index int) (item T, err error) {
	err = fmt.Errorf("invalid index: %d", index)
	return
}

type emptyIterator[T any] struct {
}

func (self *emptyIterator[T]) MoveNext() bool {
	return false
}

func (self *emptyIterator[T]) Current() T {
	var zero T
	return zero
}

// Returns a stream that contains nothing
func EmptyStream[T any]() Streamable[T] {
	return &emptyStream[T]{}
}
