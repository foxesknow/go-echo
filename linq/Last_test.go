package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Last_Empty(t *testing.T) {
	value, err := Last(data.EmptyStream[int]())

	if err == nil {
		t.Error("nothing should have been found")
	}

	// value should be the zero value
	if value != 0 {
		t.Error("value should be zero")
	}
}

func Test_Last(t *testing.T) {
	numbers := data.FromSlice([]int{5, 7, 9})
	value, err := Last(numbers)

	if err != nil {
		t.Error("should have found something")
	}

	if value != 9 {
		t.Error("value should be 9")
	}
}

func Test_Last_Generator(t *testing.T) {
	next := 0
	numbers := data.Generate(func() (int, bool) { next++; return next, next != 10 })
	value, err := Last(numbers)

	if err != nil {
		t.Error("should have found something")
	}

	if value != 9 {
		t.Error("value should be 9")
	}
}

func Test_Last_Generator_Nothing(t *testing.T) {
	next := 0
	numbers := data.Generate(func() (int, bool) { next++; return next, next != 1 })
	_, err := Last(numbers)

	if err == nil {
		t.Error("should have found nothing")
	}
}

func Test_LastOrDefault_Empty(t *testing.T) {
	value := LastOrDefault(data.EmptyStream[int](), 99)

	if value != 99 {
		t.Error("should have 99")
	}
}

func Test_LastOrDefault(t *testing.T) {
	value := LastOrDefault(data.FromValues(5, 7, 9), 99)

	if value != 9 {
		t.Error("should be 9")
	}
}

func Test_LastWhere_Empty(t *testing.T) {
	value, err := LastWhere(data.EmptyStream[int](), func(x int) bool { return x > 1 })

	// value will bet set to the "zero value"
	if value != 0 || err == nil {
		t.Error("should not have found anything")
	}
}

func Test_LastWhere(t *testing.T) {
	value, err := LastWhere(data.FromValues(5, 7, 9, 11), func(x int) bool { return x > 8 })

	if value != 11 || err != nil {
		t.Error("should have found something")
	}
}

func Test_LastOrDefaultWhere(t *testing.T) {
	value := LastOrDefaultWhere(data.FromValues(5, 7, 9), 20, func(x int) bool { return x > 9 })

	if value != 20 {
		t.Error("should have found 20")
	}
}

func Test_LastOrDefaultWhere_Found(t *testing.T) {
	value := LastOrDefaultWhere(data.FromValues(5, 7, 9, 11), 20, func(x int) bool { return x > 8 })

	if value != 11 {
		t.Error("should have found 11")
	}
}
