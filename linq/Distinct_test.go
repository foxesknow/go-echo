package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Distinct_Empty(t *testing.T) {
	distinct := Distinct(data.EmptyStream[int]())
	count := Count(distinct)
	if count != 0 {
		t.Error("expected zero items")
	}
}

func Test_Distinct(t *testing.T) {
	numbers := data.FromSlice([]int{1, 2, 3, 2, 3, 5, 7, 9, 1})
	distinct := Distinct(numbers)

	count := Count(distinct)
	if count != 6 {
		t.Error("expected 6 items")
	}

	for _, value := range []int{1, 2, 3, 5, 7, 9} {
		if ok, _ := Contains(distinct, value); !ok {
			t.Errorf("could not find %d", value)
		}
	}
}
