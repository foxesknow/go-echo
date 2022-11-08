package ini

import (
	"fmt"

	"github.com/foxesknow/go-echo/data"
)

type node struct {
	name string
	tree *Tree
}

type Tree struct {
	values         map[string]data.KeyValuePair[string, string]
	config         *Config
	children       map[string]*node
	name           string
	normalizedName string
	parent         *Tree
}

func newTree(config *Config, name string) *Tree {
	return &Tree{
		values:         make(map[string]data.KeyValuePair[string, string]),
		config:         config,
		children:       make(map[string]*node),
		name:           name,
		normalizedName: config.caseNormalize(name),
	}
}

// Returns the name of the tree in the case represented in the source.
// The root node has an empty name
func (self *Tree) Name() string {
	return self.name
}

// Returns the parent node, or nil if the tree is the parent
func (self *Tree) Parent() *Tree {
	return self.parent
}

// Moves up the tree until it reaches the root.
// If this tree is the root then it returns itself
func (self *Tree) Root() *Tree {
	tree := self

	for tree.parent != nil {
		tree = tree.parent
	}

	return tree
}

// Returns true if the tree has any children
func (self *Tree) HasChildren() bool {
	return len(self.children) != 0
}

// Returns true if the tree has any values
func (self *Tree) HasValues() bool {
	return len(self.values) != 0
}

// Adds or replaces an existing value in the tree
func (self *Tree) AddValue(name string, value string) {
	normalizedName := self.config.caseNormalize(name)
	self.values[normalizedName] = data.KeyValuePair[string, string]{Key: name, Value: value}
}

// Returns the specified child, or (nil, false) if not found
func (self *Tree) Child(name string) (*Tree, bool) {
	normalizedName := self.config.caseNormalize(name)

	node, ok := self.children[normalizedName]
	if !ok {
		return nil, false
	}

	return node.tree, true
}

func (self *Tree) Value(name string) (string, bool) {
	normalizedName := self.config.caseNormalize(name)
	pair, ok := self.values[normalizedName]
	if !ok {
		return "", false
	}

	return pair.Value, true
}

func (self *Tree) ValueOrDefault(name string, defaultValue string) string {
	value, ok := self.Value(name)
	if ok {
		return value
	}

	return defaultValue
}

// Walks the path down the tree returning the final sub tree.
// If not found the the retured error indicates which part of the path was missing
func (self *Tree) NavigateTo(path ...string) (*Tree, error) {
	tree := self

	for _, part := range path {
		child, ok := tree.Child(part)
		if !ok {
			return nil, fmt.Errorf("could not find %s", part)
		}

		tree = child
	}

	return tree, nil
}

func (self *Tree) GetOrCreateChild(name string) *Tree {
	normalizedName := self.config.caseNormalize(name)

	child, ok := self.children[normalizedName]
	if ok {
		return child.tree
	}

	tree := newTree(self.config, name)
	tree.parent = self
	node := node{name: name, tree: tree}
	self.children[normalizedName] = &node

	return tree
}

// Returns the keys and their values in the current tree.
// The case of the key will be the same as the source
func (self *Tree) Pairs() []data.KeyValuePair[string, string] {
	pairs := make([]data.KeyValuePair[string, string], len(self.values))

	i := 0
	for _, pair := range self.values {
		pairs[i] = pair
		i++
	}

	return pairs

}
