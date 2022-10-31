package ini

import (
	"bytes"
	"testing"
)

func Test_Loader_FromReader(t *testing.T) {
	source := []byte(`
	[People.Jack]
	Age = 39
	Location = Island

	[People.Ben]
	Age = 58
	Location = London
	`)

	reader := bytes.NewReader(source)
	var config = Config{}
	tree, err := FromReader(&config, reader)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if tree == nil {
		t.Error("tree should be something")
		return
	}

	jack, err := tree.NavigateTo("People", "Jack")
	if err != nil {
		t.Error("couldn't find Jack")
		return
	}

	if jack == nil {
		t.Error("Jack is nil")
	}
}

func Test_Loader_FromString(t *testing.T) {
	source := `
	[People.Jack]
	Age = 39
	Location = Island

	[People.Ben]
	Age = 58
	Location = London
	`

	var config = Config{}
	tree, err := FromString(&config, source)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if tree == nil {
		t.Error("tree should be something")
		return
	}

	jack, err := tree.NavigateTo("People", "Jack")
	if err != nil {
		t.Error("couldn't find Jack")
		return
	}

	if jack == nil {
		t.Error("Jack is nil")
	}
}

func Test_Loader_QuoteBlock(t *testing.T) {
	source := `
[People.Jack]
Age = 39
Location = Island

[People.Ben]
Age = 58
Location = London

[Greeting.Ben]
message = """
Hello Ben.
How are you?
"""
	`

	var config = Config{}
	tree, err := FromString(&config, source)

	if err != nil {
		t.Error("failed to parse")
		return
	}

	if tree == nil {
		t.Error("tree should be something")
		return
	}

	greetingBen, err := tree.NavigateTo("Greeting", "Ben")
	if err != nil {
		t.Error("couldn't find Greeting.Ben")
		return
	}

	if greetingBen == nil {
		t.Error("Jack is nil")
	}

	if message, ok := greetingBen.Value("Message"); ok {
		if message != "Hello Ben.\nHow are you?" {
			t.Error("Invalid message value")
			return
		}
	} else {
		t.Error("could not find message")
		return
	}
}

func Test_Loader_QuoteBlockNotClosed(t *testing.T) {
	source := `
[People.Jack]
Age = 39
Location = Island

[People.Ben]
Age = 58
Location = London

[Greeting.Ben]
message = """
Hello Ben.
How are you?
	`

	var config = Config{}
	tree, err := FromString(&config, source)

	if err == nil {
		t.Error("parsing should have failed")
	}

	if tree != nil {
		t.Error("tree should be nil")
	}
}

func Test_Loader_BadKeyValue(t *testing.T) {
	source := `
	[People.Jack]
	Age  39
	`

	var config = Config{}
	tree, err := FromString(&config, source)

	if err == nil {
		t.Error("parsing should have failed")
	}

	if tree != nil {
		t.Error("tree should be nil")
	}
}

func Test_Loader_BadSectionName(t *testing.T) {
	source := `
	[People..Jack]
	Age  39
	`

	var config = Config{}
	tree, err := FromString(&config, source)

	if err == nil {
		t.Error("parsing should have failed")
	}

	if tree != nil {
		t.Error("tree should be nil")
	}
}

func Test_Loader_EmptySectionName(t *testing.T) {
	source := `
	[]
	Age  39
	`

	var config = Config{}
	tree, err := FromString(&config, source)

	if err == nil {
		t.Error("parsing should have failed")
	}

	if tree != nil {
		t.Error("tree should be nil")
	}
}

func Test_Loader_MapKey(t *testing.T) {
	source := `
	Age = 39
	`

	var config = Config{
		KeyMapper: func(key string) string { return "foo-" + key },
	}
	tree, _ := FromString(&config, source)

	age := tree.ValueOrDefault("foo-age", "")
	if age != "39" {
		t.Error("could not find foo-age")
	}
}

func Test_Loader_MapValue(t *testing.T) {
	source := `
	Age = 39
	`

	var config = Config{
		ValueMapper: func(key, value string) string { return key + "-" + value },
	}
	tree, _ := FromString(&config, source)

	age := tree.ValueOrDefault("age", "")
	if age != "Age-39" {
		t.Error("value wasn't mapped")
	}
}
