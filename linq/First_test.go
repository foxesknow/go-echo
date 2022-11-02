package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/collections"
)

func Test_First_Empty(t *testing.T) {
	value, found := First(collections.EmptyEnumerable[int]())

	if found {
		t.Error("nothing should have been found")
	}

	// value should be the zero value
	if value != 0 {
		t.Error("value should be zero")
	}
}

func Test_First(t *testing.T) {
	numbers := collections.EnumerateSlice([]int{5, 7, 9})
	value, found := First(numbers)

	if !found {
		t.Error("should have found something")
	}

	if value != 5 {
		t.Error("value should be 5")
	}
}
