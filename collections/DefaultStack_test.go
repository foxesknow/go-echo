package collections

import (
	"testing"
)

func Test_Stack_Empty(t *testing.T) {
	stack := NewStack[int]()

	if stack.Count() != 0 {
		t.Error("stack should have zero count")
	}

	if !stack.IsEmpty() {
		t.Error("stack should be empty")
	}

	if _, popped := stack.Pop(); popped {
		t.Error("shouldn't be able to pop anything!")
	}

	if _, peeked := stack.Peek(); peeked {
		t.Error("shouldn't be able to peek anything!")
	}
}

func Test_Stack_WithData(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(10)

	if stack.Count() != 1 {
		t.Error("stack should have 1 item")
	}

	if stack.IsEmpty() {
		t.Error("stack isn't empty empty")
	}

	if value, peeked := stack.Peek(); !peeked || value != 10 {
		t.Error("should have peeked 10")
	}

	if value, popped := stack.Pop(); !popped || value != 10 {
		t.Error("should have popped 10")
	}

	if stack.Count() != 0 {
		t.Error("stack should be empty")
	}
}

func Test_Stack_Multiple(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(10).Push(20).Push(30)

	if stack.Count() != 3 {
		t.Error("stack should have 3 items")
	}

	if stack.IsEmpty() {
		t.Error("stack isn't empty empty")
	}

	if value, popped := stack.Pop(); !popped || value != 30 {
		t.Error("should have popped 30")
	}

	if value, popped := stack.Pop(); !popped || value != 20 {
		t.Error("should have popped 20")
	}

	if value, popped := stack.Pop(); !popped || value != 10 {
		t.Error("should have popped 10")
	}

	if stack.Count() != 0 {
		t.Error("stack should be empty")
	}
}

func Test_Stack_Clear(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(10).Push(20).Push(30)

	if stack.Count() != 3 {
		t.Error("stack should have 3 items")
	}

	stack.Clear()

	if stack.Count() != 0 {
		t.Error("stack should be empty")
	}

	if _, popped := stack.Pop(); popped {
		t.Error("shouldn't have popped anything")
	}
}

func Test_Stack_Stream(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(10).Push(20).Push(30)
	i := stack.Stream().GetStream()

	for _, value := range []int{30, 20, 10} {
		if !i.MoveNext() {
			t.Error("should have moved")
			return
		}

		if i.Current() != value {
			t.Errorf("should have %d", value)
			return
		}
	}

	if i.MoveNext() {
		t.Error("shouldn't be able to move")
		return
	}
}
