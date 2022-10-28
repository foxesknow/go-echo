package ini

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

var iniSource = `
	[Jack]
	age = 39
	location = island

	[Server]
	host = www.foo.com
	port = 8080
	path = /some/where
	`

func Test_Section_Len(t *testing.T) {
	var config = Config{}
	ini, err := FromString(&config, iniSource)

	if err != nil || ini == nil {
		t.Error("failed to parse")
		return
	}

	jack, ok := ini.Section("jack")
	if !ok {
		t.Errorf("could not find jack")
		return
	}

	if jack.Len() != 2 {
		t.Errorf("expected 2 items in Jack section")
	}

	server, ok := ini.Section("Server")
	if !ok {
		t.Errorf("could not find server")
		return
	}

	if server.Len() != 3 {
		t.Errorf("expected 3 items in server section")
	}
}

func Test_Section_Name(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	jack, _ := ini.Section("jack")
	if jack.Name() != "Jack" {
		t.Errorf("could not find jack")
		return
	}
	server, _ := ini.Section("Server")
	if server.Name() != "Server" {
		t.Errorf("expected 3 items in server section")
	}
}

func Test_Section_Value(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	jack, _ := ini.Section("jack")

	age, ok := jack.Value("age")
	if !ok {
		t.Errorf("could not find age")
		return
	}

	if age != "39" {
		t.Errorf("age should be 39")
		return
	}

	// case is ignored
	age, ok = jack.Value("AGE")
	if !ok {
		t.Errorf("could not find age")
		return
	}

	if age != "39" {
		t.Errorf("age should be 39")
		return
	}

	dob, ok := jack.Value("dob")
	if ok {
		t.Errorf("should not have found dob")
	}

	// When we don't find something the value should be the zero value
	if dob != "" {
		t.Errorf("dob should be empty")
	}
}

func Test_Section_ValueOrDefault(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	jack, _ := ini.Section("jack")

	dob := jack.ValueOrDefault("dob", "October")
	if dob != "October" {
		t.Error("dob should have defaulted to October")
	}

	age := jack.ValueOrDefault("age", "10")
	if age != "39" {
		t.Error("shouldn't get back the default value")
	}
}

func Test_Section_HasKey(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	jack, _ := ini.Section("jack")

	if !jack.HasKey("age") {
		t.Error("age should exist")
	}

	if jack.HasKey("dob") {
		t.Error("dob shouldn't exist")
	}
}

func Test_Section_Keys(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	server, _ := ini.Section("Server")

	keys := server.Keys()
	if len(keys) != 3 {
		t.Error("expected 3 keys")
	}

	if !slices.Contains(keys, "host") {
		t.Error("expected host")
	}

	if !slices.Contains(keys, "port") {
		t.Error("expected port")
	}

	if !slices.Contains(keys, "path") {
		t.Error("expected path")
	}
}

func Test_Section_Pairs(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	jack, _ := ini.Section("jack")
	pairs := jack.Pairs()

	if len(pairs) != 2 {
		t.Error("expected 2 items")
	}

	for _, pair := range pairs {
		if pair.Key == "age" && pair.Value == "39" {
			continue
		} else if pair.Key == "location" && pair.Value == "island" {
			continue
		} else {
			t.Errorf("unexpeced pair: %s", pair)
		}
	}
}

func Test_Section_AsString(t *testing.T) {
	var config = Config{}
	ini, _ := FromString(&config, iniSource)

	server, _ := ini.Section("Server")
	asString := fmt.Sprintf("%s", server)
	if asString != "Server" {
		t.Error("expected Server")
	}
}

func Test_Section_CaseSensitive(t *testing.T) {
	var config = Config{CaseSensitive: true}
	ini, _ := FromString(&config, iniSource)

	server, _ := ini.Section("Server")
	if server.HasKey("PORT") {
		t.Error("shouldn't have found PORT")
	}

	if !server.HasKey("port") {
		t.Error("should have found port")
	}
}
