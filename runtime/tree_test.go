package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToOneNode(t *testing.T) {
	tree := NewTree(nil)

	tree.Set("$a", "hello")
	val, err := tree.Get("$a")

	assert.Nil(t, err)
	assert.Equal(t, "hello", val)
}

func TestNoNodeDefinedOnSingleNode(t *testing.T) {
	tree := NewTree(nil)

	_, err := tree.Get("$a")
	assert.NotNil(t, err)
}

func TestGetFromParent(t *testing.T) {
	root := NewTree(nil)
	root.Set("$a", "hello")

	child := NewTree(root)

	val, err := child.Get("$a")

	assert.Nil(t, err)
	assert.Equal(t, "hello", val)
}

func TestGetFromGrandParent(t *testing.T) {
	root := NewTree(nil)
	root.Set("$a", "hello")

	parent := NewTree(root)
	child := NewTree(parent)

	val, err := child.Get("$a")

	assert.Nil(t, err)
	assert.Equal(t, "hello", val)
}

func TestGetFromParentWithGrandparent(t *testing.T) {
	root := NewTree(nil)

	parent := NewTree(root)
	parent.Set("$a", "hello")

	child := NewTree(parent)

	val, err := child.Get("$a")
	assert.Nil(t, err)
	assert.Equal(t, "hello", val)
}

func TestSetFromChildWithPredefinedVariable(t *testing.T) {
	root := NewTree(nil)
	root.Set("$a", "hello")

	child := NewTree(root)
	child.Set("$a", "goodbye")

	// Test child has correct element
	val, err := child.Get("$a")
	assert.Nil(t, err)
	assert.Equal(t, "goodbye", val)

	// Test that the parent also has the correct element
	val, err = root.Get("$a")
	assert.NoError(t, err)
	assert.Equal(t, "goodbye", val)
}

func TestSetFromChildWithGrandparentVariable(t *testing.T) {
	root := NewTree(nil)
	root.Set("$a", "hello")

	parent := NewTree(root)
	child := NewTree(parent)
	child.Set("$a", "goodbye")

	// Test child has correct element
	val, err := child.Get("$a")
	assert.Nil(t, err)
	assert.Equal(t, "goodbye", val)

	// Test that the parent also has the correct element
	val, err = parent.Get("$a")
	assert.NoError(t, err)
	assert.Equal(t, "goodbye", val)

	// Test that the grandparent has the correct element
	val, err = root.Get("$a")
	assert.NoError(t, err)
	assert.Equal(t, "goodbye", val)
}

func TestSetFromChildWithNoPredefinedVariable(t *testing.T) {
	root := NewTree(nil)
	child := NewTree(root)

	child.Set("$a", "hello")

	// test it exists in the child
	val, err := child.Get("$a")
	assert.Nil(t, err)
	assert.Equal(t, "hello", val)

	// Test it does not exist in the parent
	_, err = root.Get("$a")
	assert.Error(t, err)
}
