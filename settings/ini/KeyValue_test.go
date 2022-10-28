package ini

import (
	"testing"
)

func TestAsString(t *testing.T) {
	kvp := KeyValue{Key: "Jack", Value: "Island"}
	asString := kvp.String()

	if len(asString) == 0 {
		t.Errorf("string is empty!")
	}
}
