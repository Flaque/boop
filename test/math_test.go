package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddValues(t *testing.T) {
	output := run(`:val = 1 + 2
		:echo :val`)

	assert.Equal(t, "3\n", output)
}

func TestAssignValuesAndAdd(t *testing.T) {
	output := run(`
	:a = 1
	:b = 2
	:val = :a + :b
	echo :val
	`)

	assert.Equal(t, "3\n", output)
}
