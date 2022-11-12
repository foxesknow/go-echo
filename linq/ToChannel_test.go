package linq

import (
	"testing"

	"github.com/foxesknow/go-echo/data"
)

func Test_ToChannel(t *testing.T) {
	var numbers = data.FromValues(5, 10, 15, 20, -1)

	adder := make(chan int, 10)
	result := make(chan int, 10)

	go func() {
		sum := 0
		for true {
			i := <-adder

			if i == -1 {
				break
			}

			sum += i
		}

		result <- sum
	}()

	ToChannel(numbers, adder)

	sum := <-result

	if sum != 50 {
		t.Error("expected 50")
	}
}
