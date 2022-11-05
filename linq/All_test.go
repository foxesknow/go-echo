package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_All_Empty(t *testing.T) {
	called := false
	numbers := data.EmptyStream[int]()
	all := All(numbers, func(x int) bool { called = true; return true })

	if !all {
		t.Error("All should be true for an empty sequence")
	}

	if called {
		t.Error("the callback should not be called for an empty sequence")
	}
}

func Test_All(t *testing.T) {
	callCount := 0
	numbers := data.StreamSlice([]int{1, 2, 3, 4})
	all := All(numbers, func(x int) bool { callCount++; return x < 10 })

	if !all {
		t.Error("All should be true")
	}

	if callCount != 4 {
		t.Error("the callback should have been called 4 times")
	}
}

func Test_All_Odd(t *testing.T) {
	callCount := 0
	numbers := data.StreamSlice([]int{1, 2, 3, 4})
	all := All(numbers, func(x int) bool { callCount++; return (x & 1) == 1 })

	if all {
		t.Error("All should be false")
	}

	// Once the predicate fails (on the 2nd item) we'll stop
	if callCount != 2 {
		t.Error("the callback should have been called 2 times")
	}
}
