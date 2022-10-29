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

func Test_Loader_MapKey(t *testing.T) {
	source := `
	[Jack]
	age = 39
	location = island
	`

	b := []byte(source)
	reader := bytes.NewReader(b)

	var config = Config{
		KeyMapper: func(value string) string {
			return "bar-" + value
		},
	}
	ini, err := FromReader(&config, reader)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if ini == nil {
		t.Error("ini should be something")
		return
	}

	jack, ok := ini.Section("Jack")
	if !ok {
		t.Error("couldn't find Jack when case matched")
	}

	if jack.HasKey("age") {
		t.Error("shouldn't find age")
	}

	if jack.HasKey("location") {
		t.Error("shouldn't find age")
	}

	if jack.ValueOrDefault("bar-age", "") != "39" {
		t.Error("couldn't find mapped key")
	}

	if jack.ValueOrDefault("bar-location", "") != "island" {
		t.Error("couldn't find mapped key")
	}
}

func Test_Loader_MapValue(t *testing.T) {
	source := `
	[Jack]
	age = 39
	location = island
	`

	b := []byte(source)
	reader := bytes.NewReader(b)

	var config = Config{
		ValueMapper: func(value string) string {
			return "foo-" + value
		},
	}
	ini, err := FromReader(&config, reader)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if ini == nil {
		t.Error("ini should be something")
		return
	}

	jack, ok := ini.Section("Jack")
	if !ok {
		t.Error("couldn't find Jack when case matched")
	}

	if jack.ValueOrDefault("age", "") != "foo-39" {
		t.Error("mapped value is wrong")
	}

	if jack.ValueOrDefault("location", "") != "foo-island" {
		t.Error("mapped value is wrong")
	}
}

func Test_Loader_InvalidSectionName(t *testing.T) {
	source := `
	[]
	age = 39
	location = island
	`

	b := []byte(source)
	reader := bytes.NewReader(b)

	var config = Config{
		ValueMapper: func(value string) string {
			return "foo-" + value
		},
	}
	ini, err := FromReader(&config, reader)

	if err == nil {
		t.Error("parsing should have failed")
		return
	}

	if ini != nil {
		t.Error("ini should be nil as parsing failed")
		return
	}
}

func Test_Loader_KeyNotInSection(t *testing.T) {
	source := `
	age = 39
	location = island
	`

	b := []byte(source)
	reader := bytes.NewReader(b)

	var config = Config{
		ValueMapper: func(value string) string {
			return "foo-" + value
		},
	}
	ini, err := FromReader(&config, reader)

	if err == nil {
		t.Error("parsing should have failed")
		return
	}

	if ini != nil {
		t.Error("ini should be nil as parsing failed")
		return
	}
}

func Test_Loader_BadLineInSection(t *testing.T) {
	source := `
	[Jack]
	age 39
	location = island
	`

	b := []byte(source)
	reader := bytes.NewReader(b)

	var config = Config{
		ValueMapper: func(value string) string {
			return "foo-" + value
		},
	}
	ini, err := FromReader(&config, reader)

	if err == nil {
		t.Error("parsing should have failed")
		return
	}

	if ini != nil {
		t.Error("ini should be nil as parsing failed")
		return
	}
}
