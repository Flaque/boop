package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	va := NewVariable(INT, 3)
	vb := NewVariable(INT, 5)

	assert.Equal(t, 8, va.Plus(vb).AsInt())
}

func TestMinus(t *testing.T) {
	va := NewVariable(INT, 3)
	vb := NewVariable(INT, 5)

	assert.Equal(t, -2, va.Minus(vb).AsInt())
}
