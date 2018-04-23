package runtime

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Flaque/boop/parser"
	"github.com/Flaque/boop/util"
)

func getVisitor() *BeepBoopVisitor {
	var buf bytes.Buffer
	return NewBeepBoopVisitor(&buf)
}

func getInterpreters(code string) (*BeepBoopVisitor, *parser.BeepBoopParser) {
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

func TestIntLiteral(t *testing.T) {

	samples := []string{"1", "23", "123905920309443523"}

	for _, code := range samples {
		v, p := getInterpreters(code)

		val := GetVariable(v.Visit(p.Literal()))
		expected, _ := strconv.Atoi(code)
		assert.True(t, val.Equals(expected), val.Is(INT))
	}
}

func TestFloatNum(t *testing.T) {
	samples := []string{"1.0", "0.432", "23123.30"}

	for _, code := range samples {
		v, p := getInterpreters(code)

		val := GetVariable(v.Visit(p.Literal()))
		expected, _ := strconv.ParseFloat(code, 64)
		assert.True(t, val.Equals(expected))
		assert.True(t, val.Is(FLOAT))
	}
}

func TestSoloMath(t *testing.T) {
	samples := []string{"1", "3", "234.0"}

	for _, code := range samples {
		v, p := getInterpreters(code)

		val := GetVariable(v.Visit(p.Math()))
		assert.True(t, (val.Is(FLOAT) || val.Is(INT)))
	}
}

func TestUnarySubMath(t *testing.T) {

	// ints
	v, p := getInterpreters("-1")

	val := GetVariable(v.Visit(p.Math()))
	assert.True(t, val.Is(INT))
	assert.Equal(t, -1, val.AsInt())

	// Floats
	v, p = getInterpreters("-2.34")

	val = GetVariable(v.Visit(p.Math()))
	assert.True(t, val.Is(FLOAT))
	assert.Equal(t, -2.34, val.AsFloat())
}
