package data

type Stack[T any] struct {
	head  *stackNode[T]
	count int
}

type stackNode[T any] struct {
	Value T
	Next  *stackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head:  nil,
		count: 0,
	}
}

func (self *Stack[T]) Clear() {
	self.head = nil
	self.count = 0
}

func (self *Stack[T]) Push(value T) *Stack[T] {
	node := &stackNode[T]{
		Value: value,
		Next:  self.head,
	}

	self.head = node
	self.count++

	return self
}

func (self *Stack[T]) Pop() (value T, popped bool) {
	if self.count != 0 {
		oldHead := self.head

		self.head = oldHead.Next
		self.count--

		return oldHead.Value, true
	}

	return value, false
}

func (self *Stack[T]) Peek() (value T, peeked bool) {
	if self.count != 0 {
		return self.head.Value, true
	}

	return value, false
}

func (self *Stack[T]) Count() int {
	return self.count
}

func (self *Stack[T]) IsEmpty() bool {
	return self.count == 0
}

func (self *Stack[T]) Values() Stream[T] {
	if self.count == 0 {
		return EmptyStream[T]()
	}

	return &FunctionStream[T]{
		OnIterator: func() Iterator[T] {
			head := self.head
			var current T

			return &FunctionIterator[T]{
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
