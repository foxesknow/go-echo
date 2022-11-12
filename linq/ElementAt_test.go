package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_ElementAt_Empty(t *testing.T) {
	for i := -5; i != 5; i++ {
		_, err := ElementAt(data.EmptyStream[int](), i)
		if err == nil {
			t.Errorf("nothing should be found at %d", i)
		}
	}
}

func Test_ElementAt_First_And_Last(t *testing.T) {
	numbers := data.FromValues(20, 21, 22, 23)

	if item, _ := ElementAt(numbers, 0); item != 20 {
		t.Error("expected 20")
	}

	if item, _ := ElementAt(numbers, 3); item != 23 {
		t.Error("expected 20")
	}
}

func Test_ElementAt_Middle(t *testing.T) {
	numbers := data.FromValues(20, 21, 22, 23)

	if item, _ := ElementAt(numbers, 2); item != 22 {
		t.Error("expected 22")
	}
}

func Test_ElementAtOrDefault_Empty(t *testing.T) {
	for i := -5; i != 5; i++ {
		item := ElementAtOrDefault(data.EmptyStream[int](), i, 99)
		if item != 99 {
			t.Error("should have got the default value")
		}
	}
}

func Test_ElementAtOrDefault_First_And_Last(t *testing.T) {
	numbers := data.FromValues(20, 21, 22, 23)

	if item := ElementAtOrDefault(numbers, 0, 99); item != 20 {
		t.Error("expected 20")
	}

	if item := ElementAtOrDefault(numbers, 3, 99); item != 23 {
		t.Error("expected 20")
	}
}

func Test_ElementAtOrDefault_Middle(t *testing.T) {
	numbers := data.FromValues(20, 21, 22, 23)

	if item := ElementAtOrDefault(numbers, 2, 22); item != 22 {
		t.Error("expected 22")
	}
}
