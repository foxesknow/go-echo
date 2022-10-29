package ini

import "fmt"

// A key value pair representing an item in a section
type KeyValue struct {
	// The key
	Key string

	// The value
	Value string
}

func (self *KeyValue) String() string {
	return fmt.Sprintf("%s = %s", self.Key, self.Value)
}
