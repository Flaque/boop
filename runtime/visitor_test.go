package runtime

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Flaque/boop/parser"
)

func getVisitor() BeepBoopVisitor {
	var buf bytes.Buffer
	tree := NewTree(nil)
	logger := NewLogger(&buf)

	return BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}, &tree, &logger}
}

func TestNum(t *testing.T) {
	code := "2"

	p := getParser(code)
	v := getVisitor()

	total := v.Visit(p.Num())

	val, ok := total.(int)
	assert.True(t, ok)
	assert.Equal(t, 2, val)
}

func TestBoolExprTrue(t *testing.T) {
	code := "true"

	p := getParser(code)
	v := getVisitor()

	val := v.Visit(p.Boolexpr())

	b, ok := val.(bool)
	assert.True(t, ok)
	assert.True(t, b)
}

func TestBoolExprFalse(t *testing.T) {
	code := "false"

	p := getParser(code)
	v := getVisitor()

	val := v.Visit(p.Boolexpr())

	b, ok := val.(bool)
	assert.True(t, ok)
	assert.True(t, !b)
}
