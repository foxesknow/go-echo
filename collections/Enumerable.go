package collections

type Enumerator[T any] interface {
	// Moves to the next item, returning true if there is a next item, false if not
	MoveNext() bool

	// Returns the current item. If there is no current item the results are undefined
	Current() T
}

// Indicates that a type is enumerable
type Enumerable[T any] interface {
	// Create an enumerator
	GetEnumerator() Enumerator[T]
}

// Allows an enumerable to be created by calling a factory function
type FunctionEnumerable[T any] struct {
	OnGetEnumerator func() Enumerator[T]
}

func (self *FunctionEnumerable[T]) GetEnumerator() Enumerator[T] {
	return self.OnGetEnumerator()
}

// Defers enumerator calls to functions
type FunctionEnumerator[T any] struct {
	OnMoveNext func() bool
	OnCurrent  func() T
}

func (self *FunctionEnumerator[T]) MoveNext() bool {
	return self.OnMoveNext()
}

func (self *FunctionEnumerator[T]) Current() T {
	return self.OnCurrent()
}

func EmptyEnumerable[T any]() Enumerable[T] {
	return &FunctionEnumerable[T]{
		OnGetEnumerator: func() Enumerator[T] {
			return &FunctionEnumerator[T]{
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
