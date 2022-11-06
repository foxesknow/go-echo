package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Prepend(t *testing.T) {
	initial := []int{1, 2, 3, 4, 5}
	initialEnum := data.FromSlice(initial)
	appendedEnum := Prepend(initialEnum, 0)
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
