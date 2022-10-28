package ini

import (
	"testing"
)

func TestFromString(t *testing.T) {
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
