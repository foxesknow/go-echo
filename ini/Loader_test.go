package ini

import (
	"bytes"
	"testing"
)

func Test_Loader_FromString(t *testing.T) {
	source := `
	[Jack]
	age = 39
	location = island
	`

	var config = Config{}
	ini, err := FromString(&config, source)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if ini == nil {
		t.Error("ini should be something")
		return
	}

	if _, ok := ini.Section("Jack"); !ok {
		t.Error("couldn't find Jack when case matched")
	}

	if _, ok := ini.Section("jack"); !ok {
		t.Error("couldn't find Jack when case didn't match")
	}
}

func Test_Loader_FromReader(t *testing.T) {
	source := `
	[Jack]
	age = 39
	location = island
	`

	b := []byte(source)
	reader := bytes.NewReader(b)

	var config = Config{}
	ini, err := FromReader(&config, reader)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if ini == nil {
		t.Error("ini should be something")
		return
	}

	if _, ok := ini.Section("Jack"); !ok {
		t.Error("couldn't find Jack when case matched")
	}

	if _, ok := ini.Section("jack"); !ok {
		t.Error("couldn't find Jack when case didn't match")
	}
}
