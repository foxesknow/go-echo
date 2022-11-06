package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Concat(t *testing.T) {
	lhs := data.FromSlice([]int{0, 1, 2, 3, 4})
	rhs := data.FromSlice([]int{99, 100})
	concat := ToSlice(Concat(lhs, rhs))

	if len(concat) != 7 {
		t.Error("expected 7 items")
		return
	}

	expected := []int{0, 1, 2, 3, 4, 99, 100}
	for i, value := range expected {
		if concat[i] != value {
			t.Errorf("expected %d got %d", value, concat[i])
		}
	}
}
