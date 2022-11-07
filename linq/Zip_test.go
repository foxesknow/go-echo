package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Zip(t *testing.T) {
	lhs := data.FromValues(1, 2, 3)
	rhs := data.FromValues(100, 200, 300)
	zipped := Zip(lhs, rhs, func(l, r int) int { return l + r })
	var slice = ToSlice(zipped)

	if len(slice) != 3 {
		t.Error("should be 3 items")
	}

	if slice[0] != 101 {
		t.Error("should be 101")
	}

	if slice[1] != 202 {
		t.Error("should be 202")
	}

	if slice[2] != 303 {
		t.Error("should be 303")
	}
}

func Test_Zip_Different_Lengths(t *testing.T) {
	lhs := data.FromValues(1, 2, 3)
	rhs := data.FromValues(100, 200, 300, 400, 500)
	zipped := Zip(lhs, rhs, func(l, r int) int { return l + r })
	var slice = ToSlice(zipped)

	if len(slice) != 3 {
		t.Error("should be 3 items")
	}

	if slice[0] != 101 {
		t.Error("should be 101")
	}

	if slice[1] != 202 {
		t.Error("should be 202")
	}

	if slice[2] != 303 {
		t.Error("should be 303")
	}
}
