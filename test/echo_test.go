package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	code := `:echo hello`

	output := run(code)

	assert.Equal(t, "hello\n", output)
}

func TestAssign(t *testing.T) {
	code := `:a = 2
	:echo :a`

	assert.Equal(t, "2\n", run(code))
}

func TestMathAssign(t *testing.T) {
	code := `:a = 2 add 2
	echo :a`

	assert.Equal(t, "4\n", run(code))
}
