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

func (t *Tree) pointToParent(label string) {
	t.frame.Set(label, t.parent.frame)
}

func (t *Tree) Get(label string) (interface{}, error) {

	// Search current
	val := t.frame.Get(label)

	// If we're looking at a frame pointer, we're looking in the wrong place
	if isFramePointer(val) {

		// If the parent doesnt exist, but we've got a pointer to a parent, then
		// some serious fuckery has gone ary
		if t.parent == nil {
			return nil, UNABLE_TO_RESOLVE_VARIABLE
		}

		owner := t.frame.Resolve(label)
		return owner.Get(label), nil
	}

	// The value exists in normal form inside this node
	if val != nil {
		return val, nil
	}

	// Check if parent exists
	if t.parent == nil {

		// We've reached the end with nothing to show for it :(
		return nil, fmt.Errorf("Variable named %s is not defined", label)
	}

	// Search parent
	return t.parent.Get(label)
}

func (t *Tree) Set(label string, value interface{}) {

	// Attempt to find it's original owner
	owner := t.frame.Resolve(label)

	// There is no owner, we can claim ownership
	if owner == nil {
		t.frame.Set(label, value)
		return
	}

	// There is an owner, let's set it's value
}
