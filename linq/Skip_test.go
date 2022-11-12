package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Skip_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()
	skipped := Skip(empty, 10)

	if Count(skipped) != 0 {
		t.Error("should be no items")
	}
}

func Test_Skip_NonEmpty(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := Skip(data, 2)
	flattened := ToSlice(skipped)

	if len(flattened) != 2 {
		t.Error("expected 2 items")
	}

	if flattened[0] != 3 {
		t.Error("flattened[0] should be 3")
	}

	if flattened[1] != 4 {
		t.Error("flattened[1] should be 4")
	}
}

func Test_Skip_Negative_Count(t *testing.T) {
	data := data.FromValues(1, 2, 3, 4)
	skipped := Skip(data, -2)
	flattened := ToSlice(skipped)

	if len(flattened) != 4 {
		t.Error("expected 4 items")
	}
}

func Test_SkipWhile_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()
	skipped := SkipWhile(empty, func(item, _ int) bool { return item < 5 })

	if Count(skipped) != 0 {
		t.Error("should be no items")
	}
}

func Test_SkipWhile(t *testing.T) {
	empty := data.FromValues(2, 4, 6, 8, 10)
	skipped := SkipWhile(empty, func(item, _ int) bool { return item < 5 })

	if len(ToSlice(skipped)) != 3 {
		t.Error("should be 3 items")
	}
}
