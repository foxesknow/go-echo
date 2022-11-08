package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_OrderBy(t *testing.T) {
	original := data.FromValues(5, 2, 9, 1, 3, 8, 10)
	sorted := ToSlice(OrderBy(original, func(lhs, rhs *int) bool { return *lhs < *rhs }))

	if len(sorted) != 7 {
		t.Error("should have 7 items")
	}

	for i, expected := range []int{1, 2, 3, 5, 8, 9, 10} {
		if sorted[i] != expected {
			t.Errorf("Got %d, expected %d", sorted[i], expected)
		}
	}
}
