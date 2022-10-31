package ini

import (
	"testing"
)

const basicSource = `
	[People.Jack]
	Age = 39
	Location = Island

	[People.Ben]
	Age = 58
	Location = London
	`

func Test_Tree_Root(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	if tree.Parent() != nil {
		t.Error("root should have no parent")
		return
	}

	if tree.Name() != "" {
		t.Error("root should have an empty name")
		return
	}

	if tree != tree.Root() {
		t.Error("the root of the root is the root!")
		return
	}

	if tree.HasValues() {
		t.Error("there should be no values")
		return
	}

	if !tree.HasChildren() {
		t.Error("there should be children")
		return
	}
}

func Test_Tree_Children(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	if _, err := tree.NavigateTo("People", "Jack"); err != nil {
		t.Error("could not find People.Jack")
		return
	}

	if _, err := tree.NavigateTo("People", "Sawyer"); err == nil {
		t.Error("found something that didn't exist")
		return
	}
}

func Test_Tree_Value(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	jack, _ := tree.NavigateTo("People", "Jack")
	if value, ok := jack.Value("Location"); !ok || value != "Island" {
		t.Error("invalid location value")
		return
	}

	if _, ok := jack.Value("dob"); ok {
		t.Error("found something that didn't exist")
		return
	}
}

func Test_Tree_ValueOrDefault(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	jack, _ := tree.NavigateTo("People", "Jack")

	if value := jack.ValueOrDefault("Location", "who knows"); value != "Island" {
		t.Error("should have got Island")
		return
	}

	if value := jack.ValueOrDefault("dob", "July"); value != "July" {
		t.Error("should have got July")
		return
	}
}

func Test_Tree_AddValue(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	jack, _ := tree.NavigateTo("People", "Jack")
	if value, ok := jack.Value("Location"); !ok || value != "Island" {
		t.Error("invalid location value")
		return
	}

	if _, ok := jack.Value("dob"); ok {
		t.Error("found something that didn't exist")
		return
	}

	jack.AddValue("dob", "July")

	if value, _ := jack.Value("dob"); value != "July" {
		t.Error("should have got July")
		return
	}
}

func Test_Tree_GotoRoot(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	jack, _ := tree.NavigateTo("People", "Jack")
	if jack == tree {
		t.Error("jack should be different to tree")
		return
	}

	root := jack.Root()
	if root != tree {
		t.Error("root should be same as tree")
		return
	}

}

func Test_Tree_Pairs(t *testing.T) {
	var config = Config{}
	tree, _ := FromString(&config, basicSource)

	jack, _ := tree.NavigateTo("People", "Jack")

	pairs := jack.Pairs()
	if len(pairs) != 2 {
		t.Error("expected 2 pairs")
		return
	}

	for _, pair := range pairs {
		key := pair.Key
		value := pair.Value

		if key == "Age" && value == "39" {
			continue
		} else if key == "Location" && value == "Island" {
			continue
		} else {
			t.Errorf("unexpcted key: %v", pair)
		}

	}

}

func Test_Tree_CaseSensitive(t *testing.T) {
	var config = Config{CaseSensitive: true}
	tree, _ := FromString(&config, basicSource)

	if tree.Parent() != nil {
		t.Error("root should have no parent")
		return
	}

	if jack, _ := tree.NavigateTo("People", "jack"); jack != nil {
		t.Error("shouldn't have found jack as case is different")
		return
	}

	jack, _ := tree.NavigateTo("People", "Jack")

	if _, ok := jack.Value("age"); ok {
		t.Error("shouldn't have found age")
	}

	if _, ok := jack.Value("AGE"); ok {
		t.Error("shouldn't have found age")
	}

	if _, ok := jack.Value("Age"); !ok {
		t.Error("should have found age")
	}
}
