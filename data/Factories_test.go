package data

import (
	"testing"
)

func Test_Range(t *testing.T) {
	i := Range(10, 3).Iterator()

	if i.MoveNext() {
		if i.Current() != 10 {
			t.Error("expected 10")
		}
	} else {
		t.Error("MoveNext should have succeeded")
	}

	if i.MoveNext() {
		if i.Current() != 11 {
			t.Error("expected 11")
		}
	} else {
		t.Error("MoveNext should have succeeded")
	}

	if i.MoveNext() {
		if i.Current() != 12 {
			t.Error("expected 12")
		}
	} else {
		t.Error("MoveNext should have succeeded")
	}

	if i.MoveNext() {
		t.Error("There shouldn't be any more data")
	}
}

func Test_Range_Zero_Count(t *testing.T) {
	i := Range(10, 0).Iterator()

	if i.MoveNext() {
		t.Error("There shouldn't be any more data")
	}
}

func Test_Range_Negative_Count(t *testing.T) {
	i := Range(10, -3).Iterator()

	if i.MoveNext() {
		t.Error("There shouldn't be any more data")
	}
}

func Test_Repeat(t *testing.T) {
	i := Repeat("hello", 3).Iterator()

	if i.MoveNext() {
		if i.Current() != "hello" {
			t.Error("expected hello")
		}
	} else {
		t.Error("MoveNext should have succeeded")
	}

	if i.MoveNext() {
		if i.Current() != "hello" {
			t.Error("expected hello")
		}
	} else {
		t.Error("MoveNext should have succeeded")
	}

	if i.MoveNext() {
		if i.Current() != "hello" {
			t.Error("expected hello")
		}
	} else {
		t.Error("MoveNext should have succeeded")
	}

	if i.MoveNext() {
		t.Error("There shouldn't be any more data")
	}
}

func Test_FromChannel(t *testing.T) {
	channel := make(chan int)
	stop := make(chan bool)

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		channel <- 4

		<-stop
	}()

	sum := 0
	stream := FromChannel(channel)
	for i := stream.Iterator(); i.MoveNext(); {
		current := i.Current()
		sum += current

		if current == 4 {
			stop <- true
			close(channel)
		}
	}

	if sum != 10 {
		t.Error("not all data made it through")
	}
}
