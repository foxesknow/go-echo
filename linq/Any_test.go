package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Any_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()

	if Any(empty) {
		t.Error("Any should be false for an empty sequence")
	}
}

func Test_Any(t *testing.T) {
	numbers := data.StreamSlice([]int{1, 2, 3, 4})

	if !Any(numbers) {
		t.Error("Any should be true for a non-empty sequence")
	}
}

func Test_AnyWhere_Empty(t *testing.T) {
	called := false
	empty := data.EmptyStream[int]()
	any := AnyWhere(empty, func(x int) bool { called = true; return true })

	if any {
		t.Error("AnyWhere should be false for an empty sequence")
	}

	if called {
		t.Error("the callback should not be called for an empty sequence")
	}
}

func Test_AnyWhere(t *testing.T) {
	callCount := 0
	numbers := data.StreamSlice([]int{1, 2, 3, 4})
	any := AnyWhere(numbers, func(x int) bool { callCount++; return x < 10 })

	if !any {
		t.Error("Any should be true")
	}

	if callCount != 1 {
		t.Error("the callback should have been called once")
	}
}

func Test_AnyWhere_NoMatch(t *testing.T) {
	callCount := 0
	numbers := data.StreamSlice([]int{1, 2, 3, 4})
	any := AnyWhere(numbers, func(x int) bool { callCount++; return x > 20 })

	if any {
		t.Error("Any should be false")
	}

	// Once the predicate fails (on the 2nd item) we'll stop
	if callCount != 4 {
		t.Error("the callback should have been called 2 times")
	}
}
