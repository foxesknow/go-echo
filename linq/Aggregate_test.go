package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/collections"
)

func Test_Aggregate_Empty(t *testing.T) {
	empty := collections.EmptyEnumerable[int]()
	agg := Aggregate(empty, 0, func(acc, x int) int { return acc + x })
	if agg != 0 {
		t.Error("expected zero")
	}
}

func Test_Aggregate_NonEmpty(t *testing.T) {
	data := collections.EnumerateSlice([]int{1, 2, 3, 4})
	agg := Aggregate(data, 0, func(acc, x int) int { return acc + x })
	if agg != 10 {
		t.Error("expected 10")
	}
}

func Test_Aggregate_String(t *testing.T) {
	data := collections.EnumerateSlice([]string{"h", "e", "l", "l", "o"})
	agg := Aggregate(data, "", func(acc, x string) string { return acc + x })
	if agg != "hello" {
		t.Error("expected hello")
	}
}
