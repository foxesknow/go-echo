package linq

import "github.com/foxesknow/go-echo/data"

// Sends a sequence to a channel
func ToChannel[T any](stream data.Stream[T], channel chan<- T) {

	for i := stream.Iterator(); i.MoveNext(); {
		channel <- i.Current()
	}

}
