package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/collections"
	"github.com/foxesknow/go-echo/data"
)

func Test_Count_Empty(t *testing.T) {
	count := Count(data.EmptyStream[int]())
	if count != 0 {
		t.Error("expected zero items")
	}
}

func Test_Count_NonEmpty(t *testing.T) {
	count := Count(data.FromSlice([]int{1, 2, 3}))
	if count != 3 {
		t.Error("expected 3 items")
	}
}

func Test_CountWhere_Empty(t *testing.T) {
	count := CountWhere(data.EmptyStream[int](), func(x int) bool { return x == 0 })
	if count != 0 {
		t.Error("expected zero items")
	}
}

func Test_CountWhere_NonEmpty(t *testing.T) {
	// Count the odd numbers
	count := CountWhere(data.FromSlice([]int{1, 2, 3}), func(x int) bool { return (x & 1) == 1 })
	if count != 2 {
		t.Error("expected 2 items")
	}
}

func Test_CountStack(t *testing.T) {
	// Count the odd numbers
	stack := collections.NewStack[int]().Push(10).Push(20)
	count := Count(stack.Stream())
	if count != 2 {
		t.Error("expected 2 items")
	}
}
