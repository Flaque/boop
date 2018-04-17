package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFromSelfSet(t *testing.T) {
	root := NewFrame()
	root.Set("$a", "hello")

	assert.Equal(t, "hello", root.Get("$a"))
}

func TestGetNil(t *testing.T) {
	root := NewFrame()
	root.Set("$a", "hello")

	assert.Nil(t, root.Get("$notA"))
}

func TestResolve(t *testing.T) {
	root := NewFrame()
	root.Set("$a", "hello")

	parent := NewFrame()
	parent.Set("$a", &root)

	child := NewFrame()
	child.Set("$a", &parent)

	f := child.Resolve("$a")
	assert.NotNil(t, f)

	assert.Equal(t, f.Get("$a"), "hello")
}

func TestResolveNil(t *testing.T) {
	child := NewFrame()

	f := child.Resolve("$a")
	assert.Nil(t, f)
}

func TestGetFromResolved(t *testing.T) {
	root := NewFrame()
	root.Set("$a", "hello")

	child := NewFrame()
	child.Set("$a", &root)

	assert.Equal(t, "hello", child.Get("$a"))
}

func TestResolveSelf(t *testing.T) {
	root := NewFrame()
	root.Set("$a", "hello")

	f := root.Resolve("$a")
	assert.Equal(t, &root, f)
}
