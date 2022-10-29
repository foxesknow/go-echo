package properties

import (
	"fmt"

	"github.com/foxesknow/go-echo/collections"
)

type node struct {
	name string
	tree *Tree
}

type Tree struct {
	values   map[string]collections.KeyValuePair[string, string]
	config   *Config
	children map[string]*node
}

func newTree(config *Config) *Tree {
	return &Tree{
		values:   make(map[string]collections.KeyValuePair[string, string]),
		config:   config,
		children: make(map[string]*node),
	}
}

func (self *Tree) Child(name string) (*Tree, bool) {
	normalizedName := self.config.caseNormalize(name)

	node, ok := self.children[normalizedName]
	if !ok {
		return nil, false
	}

	return node.tree, true
}

func (self *Tree) NavigateTo(path ...string) (*Tree, error) {
	root := self

	for _, part := range path {
		normalizedPart := self.config.caseNormalize(part)

		child, ok := root.children[normalizedPart]
		if !ok {
			return nil, fmt.Errorf("could not find %s", part)
		}

		root = child.tree
	}

	return root, nil
}

func (self *Tree) getOrCreateChild(name string) *Tree {
	normalizedName := self.config.caseNormalize(name)

	child, ok := self.children[normalizedName]
	if ok {
		return child.tree
	}

	tree := newTree(self.config)
	node := node{name: name, tree: tree}
	self.children[normalizedName] = &node

	return tree
}
