package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinesBeforeCode(t *testing.T) {
	output := run("\n\n\necho hello")

	assert.Equal(t, "hello\n", output)
}

func TestNoNewLinesBeforeCode(t *testing.T) {
	output := run("echo hello")

	assert.Equal(t, "hello\n", output)
}

func TestNewLinesAfterCode(t *testing.T) {
	output := run("echo hello\n\n\n\n")

	assert.Equal(t, "hello\n", output)
}

func TestTabsBeforeCode(t *testing.T) {
	output := run("\t\t\techo hello")

	assert.Equal(t, "hello\n", output)
}

func TestTabsInCode(t *testing.T) {
	output := run("echo \t\thello")

	assert.Equal(t, "hello\n", output)
}

func TestTabsAfterCode(t *testing.T) {
	output := run("echo hello\t\t")

	assert.Equal(t, "hello\n", output)
}
