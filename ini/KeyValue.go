package ini

import "fmt"

type KeyValue struct {
	Key   string
	Value string
}

func (self *KeyValue) String() string {
	return fmt.Sprintf("%s = %s", self.Key, self.Value)
}
