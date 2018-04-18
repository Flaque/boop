package runtime

import (
	"errors"
	"fmt"
)

type Tree struct {
	frame  *Frame
	parent *Tree
}

var UNABLE_TO_RESOLVE_VARIABLE = errors.New("Something terribly wrong has happened in frame tree." +
	" If you see this error, please contact the Beep maintainer.")

func NewTree(parent *Tree) Tree {
	f := NewFrame()

	return Tree{
		&f,
		parent,
	}
}

func isFramePointer(val interface{}) bool {
	_, ok := val.(*Frame)
	return ok
}

func notDefinedError(label string) error {
	return fmt.Errorf("Label '%s' is not defined.", label)
}

func (t *Tree) pointToParent(label string) {
	t.frame.Set(label, t.parent.frame)
}

func (t *Tree) getFromFrame(label string) (interface{}, error) {
	val := t.frame.Get(label)
	if val == nil {
		return nil, notDefinedError(label)
	}
	return val, nil
}

func (t *Tree) Get(label string) (interface{}, error) {

	// if there's no parent, check self
	if t.parent == nil {
		return t.getFromFrame(label)
	}

	// Check parent
	owner := t.parent.frame.Resolve(label)

	// If there's an owner, let's check it
	if owner != nil {
		return owner.Get(label), nil
	}

	// If there's no owner, then it doesn't exist
	return nil, notDefinedError(label)
}

func (t *Tree) Set(label string, value interface{}) {

	// If there's no parent, we can just set this frame
	if t.parent == nil {
		t.frame.Set(label, value)
		return
	}

	// Attempt to find it's original owner
	owner := t.parent.frame.Resolve(label)

	// There is no owner, we can claim ownership
	if owner == nil {
		t.frame.Set(label, value)
		return
	}

	// There is an owner, let's set it's value and leave a trail for children
	owner.Set(label, value)
	t.pointToParent(label)
}
