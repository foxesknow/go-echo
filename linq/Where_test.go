package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Where(t *testing.T) {
	var numbers = data.FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var oddNumbers = ToSlice(Where(numbers, func(n int) bool { return (n & 1) == 1 }))

	if len(oddNumbers) != 5 {
		t.Error("should have 5 items")
	}

	for i, v := range []int{1, 3, 5, 7, 9} {
		if oddNumbers[i] != v {
			t.Errorf("Expected %d got %d", v, oddNumbers[i])
		}
	}
}

func Test_WhereIndex(t *testing.T) {
	var numbers = data.FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var oddNumbers = ToSlice(WhereIndex(numbers, func(_, index int) bool { return index > 4 }))

	if len(oddNumbers) != 4 {
		t.Error("should have 4 items")
	}

	for i, v := range []int{6, 7, 8, 9} {
		if oddNumbers[i] != v {
			t.Errorf("Expected %d got %d", v, oddNumbers[i])
		}
	}
}
