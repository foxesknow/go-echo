package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_First_Empty(t *testing.T) {
	value, found := First(data.EmptyStream[int]())

	if found {
		t.Error("nothing should have been found")
	}

	// value should be the zero value
	if value != 0 {
		t.Error("value should be zero")
	}
}

func Test_First(t *testing.T) {
	numbers := data.FromSlice([]int{5, 7, 9})
	value, found := First(numbers)

	if !found {
		t.Error("should have found something")
	}

	if value != 5 {
		t.Error("value should be 5")
	}
}

func Test_FirstOrDefault_Empty(t *testing.T) {
	value := FirstOrDefault(data.EmptyStream[int](), 99)

	if value != 99 {
		t.Error("should have 99")
	}
}

func Test_FirstOrDefault(t *testing.T) {
	value := FirstOrDefault(data.FromValues(5, 7, 9), 99)

	if value != 5 {
		t.Error("should have 5")
	}
}

func Test_FirstWhere_Empty(t *testing.T) {
	value, found := FirstWhere(data.EmptyStream[int](), func(x int) bool { return x > 1 })

	// value will bet set to the "zero value"
	if value != 0 || found {
		t.Error("should not have found anything")
	}
}

func Test_FirstWhere(t *testing.T) {
	value, found := FirstWhere(data.FromValues(5, 7, 9), func(x int) bool { return x > 8 })

	if value != 9 || !found {
		t.Error("should have found something")
	}
}

func Test_FirstOrDefaultWhere(t *testing.T) {
	value := FirstOrDefaultWhere(data.FromValues(5, 7, 9), 20, func(x int) bool { return x > 9 })

	if value != 20 {
		t.Error("should have found 20")
	}
}

func Test_FirstOrDefaultWhere_Found(t *testing.T) {
	value := FirstOrDefaultWhere(data.FromValues(5, 7, 9, 11), 20, func(x int) bool { return x > 8 })

	if value != 9 {
		t.Error("should have found 9")
	}
}
