package runtime

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Flaque/boop/parser"
	"github.com/Flaque/boop/util"
)

func getVisitor() BeepBoopVisitor {
	var buf bytes.Buffer
	tree := NewTree(nil)
	logger := NewLogger(&buf)

	return BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}, &tree, &logger}
}

func getInterpreters(code string) (BeepBoopVisitor, *parser.BeepBoopParser) {
	return getVisitor(), getParser(code)
}

func TestNum(t *testing.T) {
	code := "2"

	v, p := getInterpreters(code)

	total := v.Visit(p.Num())

	val, ok := total.(int)
	assert.True(t, ok)
	assert.Equal(t, 2, val)
}

func TestBoolExprTrue(t *testing.T) {
	code := "true"

	v, p := getInterpreters(code)

	val := v.Visit(p.Boolexpr())

	b, ok := val.(bool)
	assert.True(t, ok)
	assert.True(t, b)
}

func TestBoolExprFalse(t *testing.T) {
	code := "false"

	v, p := getInterpreters(code)

	val := v.Visit(p.Boolexpr())

	b, ok := val.(bool)
	assert.True(t, ok)
	assert.True(t, !b)
}

func TestQuotedLiteral(t *testing.T) {
	samples := []string{`"hello there"`, "\"hi \n there\"", `""`, `"hey"`}

	for _, code := range samples {
		v, p := getInterpreters(code)
		val := v.Visit(p.Literal())

		str, ok := val.(string)
		assert.True(t, ok)
		assert.Equal(t, util.RemoveQuotes(code), str)
	}
}
