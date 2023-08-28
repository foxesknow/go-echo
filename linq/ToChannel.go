package linq

import "github.com/foxesknow/go-echo/data"

// Sends a sequence to a channel
func ToChannel[T any](stream data.Streamable[T], channel chan<- T) {

	for i := stream.GetStream(); i.MoveNext(); {
		channel <- i.Current()
	}
}
