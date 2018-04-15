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
