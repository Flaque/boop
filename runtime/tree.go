package runtime

import (
	"fmt"
)

type Tree struct {
	frame  *Frame
	parent *Tree
}

func NewTree(parent *Tree) Tree {
	f := NewFrame()

	return Tree{
		&f,
		parent,
	}
}

func (t *Tree) Get(label string) (interface{}, error) {

	val := t.frame.Get(label)

	if val != nil {
		return val, nil
	}

	return val, fmt.Errorf("Variable named %s is not defined", label)
}

func (t *Tree) Set(label string, value interface{}) {
	// Do something
}
