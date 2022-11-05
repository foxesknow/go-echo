package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Contains_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()
	if found, _ := Contains(empty, 10); found {
		t.Error("should not have found anything")
	}

	if _, index := Contains(empty, 10); index != -1 {
		t.Error("index should be -1")
	}
}

func Test_Contains(t *testing.T) {
	values := data.StreamValues(3, 5, 7, 9)
	if found, _ := Contains(values, 10); found {
		t.Error("should not have found anything")
	}

	if found, index := Contains(values, 7); !found && index != 2 {
		t.Error("should have found 7 at index 2")
	}
}

func Test_Contains_Not_Found(t *testing.T) {
	values := data.StreamValues(3, 5, 7, 9)

	if found, index := Contains(values, 20); found || index != -1 {
		t.Error("should not have found anything")
	}
}

func Test_ContainsWhere_Empty(t *testing.T) {
	empty := data.EmptyStream[int]()
	if found, _ := ContainsWhere(empty, func(x int) bool { return x == 1 }); found {
		t.Error("should not have found anything")
	}

	if _, index := ContainsWhere(empty, func(x int) bool { return x == 1 }); index != -1 {
		t.Error("index should be -1")
	}
}

func Test_ContainsWhere(t *testing.T) {
	values := data.StreamValues(3, 5, 7, 9)
	if found, _ := ContainsWhere(values, func(x int) bool { return x < 0 }); found {
		t.Error("should not have found anything")
	}

	if found, index := ContainsWhere(values, func(x int) bool { return x > 6 }); !found && index != 2 {
		t.Error("should have found 7 at index 2")
	}
}
