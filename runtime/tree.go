package runtime

import (
	"errors"
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

func isFramePointer(val interface{}) bool {
	_, ok := val.(*Frame)
	return ok
}

func (t *Tree) Get(label string) (interface{}, error) {

	// Search current
	val := t.frame.Get(label)

	// If we're looking at a frame pointer, we're looking in the wrong place
	if isFramePointer(val) {

		// If the parent doesnt exist, but we've got a pointer to a parent, then
		// some serious fuckery has gone ary
		if t.parent == nil {
			return nil, errors.New("Something terribly wrong has happened in frame tree." +
				" If you see this error, please contact the Beep maintainer.")
		}

		return t.parent.Get(label)
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

	// If there's no parent, then we don't need to check it
	if t.parent == nil {
		t.frame.Set(label, value)
		return
	}

	// If there is a parent, and it has the item, we should
	// point our new value to the parent's memory location
	if val, err := t.parent.Get(label); err == nil {

		// If the value is a pointer to a hash table, then it's not the
		// owner of the label
		// This parent is the owner of the original value, so we can reset it here
		if !isFramePointer(val) {
			t.parent.frame.Set(label, value)

			// Let future generations know where to find the label owner
			t.frame.Set(label, t.parent.frame)

			return
		}

		//This parent is NOT the owner, so we need to keep searching
		t.parent.Set(label, value)
		t.frame.Set(label, t.parent.frame)
	}

	// There is a parent, but it doesn't know about the label so
	// we can just set it here.
	t.frame.Set(label, value)
}
