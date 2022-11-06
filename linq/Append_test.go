package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Append(t *testing.T) {
	initial := []int{0, 1, 2, 3, 4}
	initialEnum := data.FromSlice(initial)
	appendedEnum := Append(initialEnum, 5)
	slice := ToSlice(appendedEnum)

	if slice == nil {
		t.Error("no slice returned")
		return
	}

	if len(slice) != 6 {
		t.Error("should have 6 items")
		return
	}

	for i := 0; i < 6; i++ {
		if slice[i] != i {
			t.Errorf("expected %d got %d", i, slice[i])
		}
	}
}

func Test_Append_BeyondCapacity(t *testing.T) {
	initial := []int{0, 1, 2, 3, 4, 5, 6, 7}
	initialEnum := data.FromSlice(initial)
	appendedEnum := Append(initialEnum, 8)
	slice := ToSlice(appendedEnum)

	if slice == nil {
		t.Error("no slice returned")
		return
	}

	if len(slice) != 9 {
		t.Error("should have 9 items")
		return
	}

	for i := 0; i < 9; i++ {
		if slice[i] != i {
			t.Errorf("expected %d got %d", i, slice[i])
		}
	}
}
