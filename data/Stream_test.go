package data

import (
	"testing"
)

func Test_EmptyStream(t *testing.T) {
	stream := EmptyStream[int]()
	i := stream.GetStream()

	if i.MoveNext() {
		t.Error("shouldn't be able to move in an empty stream")
	}

	if i.Current() != 0 {
		t.Error("we should get the zero value")
	}

	collection, ok := stream.(Collection)
	if !ok {
		t.Error("empty stream must implement Collection")
	}

	if collection.Count() != 0 {
		t.Error("empty stream should have count of 0")
	}
}

func Test_Streamable(t *testing.T) {
	numbers := FromSlice[int]([]int{1, 2, 3})
	sum := sumStream(numbers)

	if sum != 6 {
		t.Errorf("Expected 6 got %d", sum)
	}
}

func sumStream(numbers Streamable[int]) int {
	total := 0

	for i := numbers.GetStream(); i.MoveNext(); {
		total += i.Current()
	}

	return total
}
