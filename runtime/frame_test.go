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
