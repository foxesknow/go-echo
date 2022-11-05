package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Last_Empty(t *testing.T) {
	value, found := Last(data.EmptyStream[int]())

	if found {
		t.Error("nothing should have been found")
	}

	// value should be the zero value
	if value != 0 {
		t.Error("value should be zero")
	}
}

func Test_Last(t *testing.T) {
	numbers := data.StreamSlice([]int{5, 7, 9})
	value, found := Last(numbers)

	if !found {
		t.Error("should have found something")
	}

	if value != 9 {
		t.Error("value should be 9")
	}
}

func Test_LastOrDefault_Empty(t *testing.T) {
	value := LastOrDefault(data.EmptyStream[int](), 99)

	if value != 99 {
		t.Error("should have 99")
	}
}

func Test_LastOrDefault(t *testing.T) {
	value := LastOrDefault(data.StreamValues(5, 7, 9), 99)

	if value != 9 {
		t.Error("should be 9")
	}
}

func Test_LastWhere_Empty(t *testing.T) {
	value, found := LastWhere(data.EmptyStream[int](), func(x int) bool { return x > 1 })

	// value will bet set to the "zero value"
	if value != 0 || found {
		t.Error("should not have found anything")
	}
}

func Test_LastWhere(t *testing.T) {
	value, found := LastWhere(data.StreamValues(5, 7, 9, 11), func(x int) bool { return x > 8 })

	if value != 11 || !found {
		t.Error("should have found something")
	}
}

func Test_LastOrDefaultWhere(t *testing.T) {
	value := LastOrDefaultWhere(data.StreamValues(5, 7, 9), 20, func(x int) bool { return x > 9 })

	if value != 20 {
		t.Error("should have found 20")
	}
}

func Test_LastOrDefaultWhere_Found(t *testing.T) {
	value := LastOrDefaultWhere(data.StreamValues(5, 7, 9, 11), 20, func(x int) bool { return x > 8 })

	if value != 11 {
		t.Error("should have found 11")
	}
}
