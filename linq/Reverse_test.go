package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Reverse(t *testing.T) {
	initial := data.FromValues(1, 2, 3, 4, 5)
	reversed := Reverse(initial)
	slice := ToSlice(reversed)

	if slice == nil {
		t.Error("no slice returned")
		return
	}

	if len(slice) != 5 {
		t.Error("should have 5 items")
		return
	}

	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		if slice[i] != j {
			t.Errorf("unexpected value in at index %d", i)
		}
	}
}

func Test_ReverseEmpty(t *testing.T) {
	reversed := Reverse(data.EmptyStream[int]())
	slice := ToSlice(reversed)

	if slice == nil {
		t.Error("no slice returned")
		return
	}

	if len(slice) != 0 {
		t.Error("slice should be empty")
		return
	}
}
