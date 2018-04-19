package runtime

import (
	"fmt"
)

type Tree struct {
	frame  *Frame
	Parent *Tree
}

func NewTree(parent *Tree) Tree {
	f := NewFrame()

	return Tree{
		&f,
		parent,
	}
}

func notDefinedError(label string) error {
	return fmt.Errorf("label '%s' is not defined", label)
}

func (t *Tree) owner(label string) *Tree {
	val := t.frame.Get(label)

	if val != nil {
		return t
	}

	if t.Parent == nil {
		return nil // no owner
	}

	return t.Parent.owner(label)
}

func (t *Tree) Get(label string) (interface{}, error) {
	owner := t.owner(label)
	if owner == nil {
		return nil, notDefinedError(label)
	}

	return owner.frame.Get(label), nil
}

func (t *Tree) Set(label string, value interface{}) {
	owner := t.owner(label)

	// No owner, we'll claim it
	if owner == nil {
		t.frame.Set(label, value)
		return
	}

	owner.frame.Set(label, value)
}
