package data

import (
	"testing"
)

func Test_KeyValue_Pair_AsString(t *testing.T) {
	kvp := KeyValuePair[string, string]{Key: "Jack", Value: "Island"}
	asString := kvp.String()

	if len(asString) == 0 {
		t.Errorf("string is empty!")
	}
}
