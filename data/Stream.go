package data

// Indicates that a type can stream the values it contains
type Stream[T any] interface {
	// Create an iterator to the values in the stream
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	// Moves to the next item, returning true if there is a next item, false if not
	MoveNext() bool

	// Returns the current item. If there is no current item the results are undefined
	Current() T
}

// Allows an stream to be created by calling a factory function
type FunctionStream[T any] struct {
	OnIterator func() Iterator[T]
}

func (self *FunctionStream[T]) Iterator() Iterator[T] {
	return self.OnIterator()
}

// Defers iterator calls to functions
type FunctionIterator[T any] struct {
	OnMoveNext func() bool
	OnCurrent  func() T
}

func (self *FunctionIterator[T]) MoveNext() bool {
	return self.OnMoveNext()
}

func (self *FunctionIterator[T]) Current() T {
	return self.OnCurrent()
}

type emptyStream[T any] struct {
}

func (self *emptyStream[T]) Iterator() Iterator[T] {
	return &emptyIterator[T]{}
}

func (self *emptyStream[T]) Count() int {
	return 0
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
func EmptyStream[T any]() Stream[T] {
	return &emptyStream[T]{}
}
