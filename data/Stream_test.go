package data

import (
	"testing"
)

func Test_EmptyStream(t *testing.T) {
	stream := EmptyStream[int]()
	i := stream.Iterator()

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
