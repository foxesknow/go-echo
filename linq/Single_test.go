package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_Single(t *testing.T) {
	value, err := Single(data.FromValues(8))

	if err != nil {
		t.Error("should have one item")
	}

	if value != 8 {
		t.Error("should have got 8")
	}
}

func Test_Single_Empty(t *testing.T) {
	_, err := Single(data.EmptyStream[int]())

	if err == nil {
		t.Error("should have errored as the stream is empty")
	}
}

func Test_Single_ManyItems_Collection(t *testing.T) {
	_, err := Single(data.FromValues(8, 9, 10))

	if err == nil {
		t.Error("should have errored as the stream has more than one item in it")
	}
}

func Test_Single_ManyItems_Channel(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 8
	ch <- 9
	ch <- 10

	_, err := Single(data.FromChannel(ch))

	if err == nil {
		t.Error("should have errored as the stream has more than one item in it")
	}
}

func Test_Single_NoItems_Channel(t *testing.T) {
	ch := make(chan int, 10)
	close(ch)

	_, err := Single(data.FromChannel(ch))

	if err == nil {
		t.Error("should have errored as the stream has more than one item in it")
	}
}

func Test_SingleWhere(t *testing.T) {
	value, err := SingleWhere(data.FromValues(8, 9, 10, 11), func(x int) bool { return x == 10 })

	if err != nil {
		t.Error("should have one item")
	}

	if value != 10 {
		t.Error("value should be 10")
	}
}

func Test_SingleWhere_MultipleValues(t *testing.T) {
	_, err := SingleWhere(data.FromValues(8, 9, 10, 11, 3, 10, 1), func(x int) bool { return x == 10 })

	if err == nil {
		t.Error("we should have an error")
	}
}

func Test_SingleWhere_NoMatch(t *testing.T) {
	_, err := SingleWhere(data.FromValues(8, 9, 10, 11, 3, 10, 1), func(x int) bool { return x == 99 })

	if err == nil {
		t.Error("we should have an error as the predicate didn't match anything")
	}
}
