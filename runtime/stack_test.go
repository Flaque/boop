package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeFrame() Frame {
	f := NewFrame()
	f.Set("cats", 1)
	f.Set("dogs", 2)
	return f
}

func TestStackPushPop(t *testing.T) {
	stack := NewStack()
	stack.Push(fakeFrame())

	// one thing pushed
	assert.Equal(t, stack.Len(), 1)

	// Data still entact after popping
	frame := stack.Pop()
	assert.Equal(t, frame.Get("dogs"), 2)
}
