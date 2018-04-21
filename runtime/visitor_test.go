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

	val := GetVariable(v.Visit(p.Num()))

	assert.True(t, val.Is(INT))
	assert.True(t, val.Equals(2))
}

func TestBoolExprTrue(t *testing.T) {
	code := "true"

	v, p := getInterpreters(code)

	val := GetVariable(v.Visit(p.Boolexpr()))

	assert.True(t, val.Is(BOOL))
	assert.True(t, val.Equals(true))
}

func TestBoolExprFalse(t *testing.T) {
	code := "false"

	v, p := getInterpreters(code)

	val := GetVariable(v.Visit(p.Boolexpr()))

	assert.True(t, val.Is(BOOL))
	assert.True(t, val.Equals(false))
}

func TestQuotedLiteral(t *testing.T) {
	samples := []string{`"hello there"`, "\"hi \n there\"", `""`, `"hey"`}

	for _, code := range samples {
		v, p := getInterpreters(code)
		val := GetVariable(v.Visit(p.Literal()))

		assert.True(t, val.Is(STRING))
		assert.Equal(t, util.RemoveQuotes(code), val.AsString())
	}
}

func TestStringLiteral(t *testing.T) {
	samples := []string{"hello", "dogs", "cats"}

	for _, code := range samples {
		v, p := getInterpreters(code)

		val := GetVariable(v.Visit(p.Literal()))

		assert.True(t, val.Is(STRING))
		assert.Equal(t, code, val.AsString())
	}
}

func TestBoolLiteral(t *testing.T) {

	// true
	v, p := getInterpreters("true")
	val := GetVariable(v.Visit(p.Literal()))
	assert.True(t, val.Equals(true), val.Is(BOOL))

	// false
	v, p = getInterpreters("false")
	val = GetVariable(v.Visit(p.Literal()))
	assert.True(t, val.Equals(false), val.Is(BOOL))
}
