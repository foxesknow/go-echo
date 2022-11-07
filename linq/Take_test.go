package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Take_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()
	taken := Take(empty, 10)

	if Count(taken) != 0 {
		t.Error("should be no items")
	}
}

func Test_Take(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := Take(data, 2)
	flattened := ToSlice(skipped)

	if len(flattened) != 2 {
		t.Error("expected 2 items")
	}

	if flattened[0] != 1 {
		t.Error("flattened[0] should be 1")
	}

	if flattened[1] != 2 {
		t.Error("flattened[1] should be 2")
	}
}

func Test_Take_More_Than_Available_1(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := Take(data, 10)
	flattened := ToSlice(skipped)

	if len(flattened) != 4 {
		t.Error("expected 4 items")
	}
}

func Test_Take_Size_Of_Stream(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := Take(data, 4)
	flattened := ToSlice(skipped)

	if len(flattened) != 4 {
		t.Error("expected 4 items")
	}
}

func Test_Take_Negative_Count(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := Take(data, -2)
	flattened := ToSlice(skipped)

	if len(flattened) != 0 {
		t.Error("expected 0 items")
	}
}

func Test_TakeWhileIndex_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()
	taken := TakeWhileIndex(empty, func(value int, index int) bool { return index < 10 })

	if Count(taken) != 0 {
		t.Error("should be no items")
	}
}

func Test_TakeWhileIndex(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := TakeWhileIndex(data, func(value int, index int) bool { return index < 2 })
	flattened := ToSlice(skipped)

	if len(flattened) != 2 {
		t.Error("expected 2 items")
	}

	if flattened[0] != 1 {
		t.Error("flattened[0] should be 1")
	}

	if flattened[1] != 2 {
		t.Error("flattened[1] should be 2")
	}
}

func Test_TakeWhileIndex_2(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := TakeWhileIndex(data, func(value int, index int) bool { return index < 20 })
	flattened := ToSlice(skipped)

	if len(flattened) != 4 {
		t.Error("expected 4 items")
	}
}
