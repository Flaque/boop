package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncDef(t *testing.T) {
	output := run(`
	func jazz do 
		echo hello 
	end

	jazz`)

	assert.Equal(t, "hello\n", output)
}

func TestFuncDefWithArgs(t *testing.T) {
	output := run(`
	func jazz $a $b do
		$val = $a + $b
		echo $val
	end

	jazz 1 2
	`)

	assert.Equal(t, "3\n", output)
}
