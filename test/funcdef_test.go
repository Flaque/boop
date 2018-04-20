package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncDef(t *testing.T) {
	output := run(`
	func :jazz do 
		:echo hello 
	end

	:jazz`)

	assert.Equal(t, "hello\n", output)
}

func TestFuncDefWithArgs(t *testing.T) {
	output := run(`
	func :jazz :a :b do
		:val = :a + :b
		:echo :val
	end

	:jazz 1 2
	`)

	assert.Equal(t, "3\n", output)
}

func TestFuncDefWithStringArgs(t *testing.T) {
	output := run(`
	func :jazz :a :b do
		echo :a :b
	end

	:jazz hello there
	`)

	assert.Equal(t, "hello there\n", output)
}

func TestFuncDefWithVariableArgs(t *testing.T) {
	output := run(`
	func :jazz :str :num do
		echo :num
		echo :str
	end

	:jazz cats 3`)

	assert.Equal(t, "3\ncats\n", output)
}
