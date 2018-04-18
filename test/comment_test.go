package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComment(t *testing.T) {
	output := run(`# echo goodbye 
    echo hello`)

	assert.Equal(t, "hello\n", output)
}

func TestCommentAfter(t *testing.T) {
	output := run(`echo hello # $a = goodbye`)
	assert.Equal(t, "hello\n", output)
}
