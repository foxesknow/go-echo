package properties

import (
	"testing"
)

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
