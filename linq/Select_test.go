package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Select(t *testing.T) {
	var sequence = Select(data.FromValues("Jack", "Ben", "Sawyer"), func(x string) int { return len(x) })
	flattened := ToSlice(sequence)

	if len(flattened) != 3 {
		t.Error("should have 3 items")
	}

	if flattened[0] != 4 {
		t.Error("expected 4")
	}

	if flattened[1] != 3 {
		t.Error("expected 3")
	}

	if flattened[2] != 6 {
		t.Error("expected 6")
	}
}

func Test_SelectIndex(t *testing.T) {
	var sequence = SelectIndex(data.FromValues("Jack", "Ben", "Sawyer"), func(x string, index int) int { return len(x) + (10 * index) })
	flattened := ToSlice(sequence)

	if len(flattened) != 3 {
		t.Error("should have 3 items")
	}

	if flattened[0] != 4 {
		t.Error("expected 4")
	}

	if flattened[1] != 13 {
		t.Error("expected 13")
	}

	if flattened[2] != 26 {
		t.Error("expected 26")
	}
}

func Test_SelectMany_Empty(t *testing.T) {
	var sequence = SelectMany(data.EmptyStream[[]int](), func(inner []int) data.Streamable[int] { return data.FromSlice(inner) })
	flattened := ToSlice(sequence)

	if len(flattened) != 0 {
		t.Error("flattened should be empty")
	}
}

func Test_SelectMany(t *testing.T) {
	numbers := make([][]int, 3)
	numbers[0] = []int{0, 1, 2}
	numbers[1] = []int{3, 4, 5}
	numbers[2] = []int{6, 7, 8}

	var sequence = SelectMany(data.FromSlice(numbers), func(inner []int) data.Streamable[int] { return data.FromSlice(inner) })
	flattened := ToSlice(sequence)

	for i := 0; i < 9; i++ {
		if flattened[i] != i {
			t.Errorf("expected %d for %d", flattened[i], i)
		}
	}
}
