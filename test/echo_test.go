package test

import (
	"bytes"
	"testing"

	"github.com/flaque/boop/runtime"
	"github.com/stretchr/testify/assert"
)

func run(code string) string {
	var buf bytes.Buffer
	runtime.Run(code, &buf)
	return buf.String()
}

func TestEcho(t *testing.T) {
	code := `echo hello`

	output := run(code)

	assert.Equal(t, "hello\n", output)
}

func TestAssign(t *testing.T) {
	code := `$a = 2
	echo $a`

	assert.Equal(t, "2\n", run(code))
}

func TestMathAssign(t *testing.T) {
	code := `$a = 2 + 2
	echo $a`

	assert.Equal(t, "4\n", run(code))
}
