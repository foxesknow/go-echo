package linq

import "fmt"

func makeInvalidIndex(index int) error {
	return fmt.Errorf("invalid index: %d", index)
}

func makeNoItemsInStream() error {
	return fmt.Errorf("no items in stream")
}
