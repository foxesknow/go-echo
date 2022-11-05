package data

// Indicates that a type is enumerable
type Stream[T any] interface {
	// Create an enumerator
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	// Moves to the next item, returning true if there is a next item, false if not
	MoveNext() bool

	// Returns the current item. If there is no current item the results are undefined
	Current() T
}

// Allows an enumerable to be created by calling a factory function
type FunctionStream[T any] struct {
	OnIterator func() Iterator[T]
}

func (self *FunctionStream[T]) Iterator() Iterator[T] {
	return self.OnIterator()
}

// Defers enumerator calls to functions
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

func EmptyStream[T any]() Stream[T] {
	return &FunctionStream[T]{
		OnIterator: func() Iterator[T] {
			return &FunctionIterator[T]{
				OnMoveNext: func() bool {
					return false
				},
				OnCurrent: func() T {
					var value T
					return value
				},
			}
		},
	}
}
