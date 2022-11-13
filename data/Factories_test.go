package data

import (
	"context"
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

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		channel <- 4
		close(channel)
	}()

	sum := 0
	stream := FromChannel(channel)
	for i := stream.Iterator(); i.MoveNext(); {
		current := i.Current()
		sum += current

		if current == 4 {
			break
		}
	}

	if sum != 10 {
		t.Error("not all data made it through")
	}
}

func Test_FromChannel_CloseBeforeAllRead(t *testing.T) {
	channel := make(chan int)

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		channel <- 4
		close(channel)
	}()

	sum := 0
	stream := FromChannel(channel)
	for i := stream.Iterator(); i.MoveNext(); {
		current := i.Current()
		sum += current

		if current == 5 {
			break
		}
	}

	if sum != 10 {
		t.Error("not all data made it through")
	}
}

func Test_FromChannelWithContext(t *testing.T) {
	channel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		channel <- 4
		cancel()
	}()

	sum := 0
	stream := FromChannelWithContext(channel, ctx)
	for i := stream.Iterator(); i.MoveNext(); {
		current := i.Current()
		sum += current
	}

	if sum != 10 {
		t.Error("not all data made it through")
	}
}

func Test_FromChannelWhile(t *testing.T) {
	channel := make(chan int, 5)
	result := make(chan int)

	go func() {
		stream := FromChannelWhile(channel, func(i int) bool { return i != -1 })
		sum := 0

		for i := stream.Iterator(); i.MoveNext(); {
			sum += i.Current()
		}

		result <- sum
	}()

	channel <- 1
	channel <- 5
	channel <- 7
	channel <- 9
	channel <- -1

	sum := <-result

	if sum != 22 {
		t.Error("expected 22")
	}
}

func Test_Generate(t *testing.T) {
	next := 0
	stream := Generate(func() (value int, keepGoing bool) {
		next++
		value = next
		keepGoing = (next != 5)

		return
	})

	sum := 0
	for i := stream.Iterator(); i.MoveNext(); {
		current := i.Current()
		sum += current
	}

	if sum != 10 {
		t.Error("not all data made it through")
	}
}
