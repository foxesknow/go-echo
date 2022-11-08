package data

import "fmt"

// Groups a key and a value together
type KeyValuePair[K any, V any] struct {
	Key   K
	Value V
}

func (self *KeyValuePair[K, V]) String() string {
	return fmt.Sprintf("%v = %v", self.Key, self.Value)
}
