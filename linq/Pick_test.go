package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Pick_Max_Empty(t *testing.T) {
	numbers := data.FromValues[int]()

	max, err := Pick(numbers, func(candidate, current int) bool { return candidate > current })

	if err == nil {
		t.Error("shouldn't have found anything as the sequence is empty")
	}

	if max != 0 {
		t.Error("max should be the default int value")
	}
}

func Test_Pick_Max(t *testing.T) {
	numbers := data.FromValues(8, 1, 78, 31, 2, 7, 11)

	max, err := Pick(numbers, func(candidate, current int) bool { return candidate > current })

	if err != nil {
		t.Error("expected to find something")
	}

	if max != 78 {
		t.Error("should have found 78 as it's the max value")
	}
}

func Test_Pick_Max_One_Item(t *testing.T) {
	numbers := data.FromValues(8)

	max, err := Pick(numbers, func(candidate, current int) bool { return candidate > current })

	if err != nil {
		t.Error("expected to find something")
	}

	if max != 8 {
		t.Error("should have found 8 as it's the max value")
	}
}

func Test_Pick_Min(t *testing.T) {
	numbers := data.FromValues(8, 1, 78, 31, 2, 7, 11)

	max, err := Pick(numbers, func(candidate, current int) bool { return candidate < current })

	if err != nil {
		t.Error("expected to find something")
	}

	if max != 1 {
		t.Error("should have found 1 as it's the min value")
	}
}
