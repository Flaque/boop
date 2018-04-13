// Code generated from parser/Do.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 31, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 6, 4, 19, 10, 4, 13, 4, 14, 4, 20, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 6, 6, 28, 10, 6, 13, 6, 14, 6, 29, 2, 2, 7, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 3, 2, 4, 4, 2, 12, 12, 15, 15, 3, 2, 50, 59, 2, 32,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15, 3, 2, 2, 2, 7, 18, 3, 2, 2,
	2, 9, 22, 3, 2, 2, 2, 11, 27, 3, 2, 2, 2, 13, 14, 7, 45, 2, 2, 14, 4, 3,
	2, 2, 2, 15, 16, 7, 47, 2, 2, 16, 6, 3, 2, 2, 2, 17, 19, 9, 2, 2, 2, 18,
	17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2,
	2, 21, 8, 3, 2, 2, 2, 22, 23, 7, 34, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25,
	8, 5, 2, 2, 25, 10, 3, 2, 2, 2, 26, 28, 9, 3, 2, 2, 27, 26, 3, 2, 2, 2,
	28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 12, 3,
	2, 2, 2, 5, 2, 20, 29, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "", "' '",
}

var lexerSymbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "INT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "NEWLINE", "WHITESPACE", "INT",
}

type DoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDoLexer(input antlr.CharStream) *DoLexer {

	l := new(DoLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Do.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DoLexer tokens.
const (
	DoLexerT__0       = 1
	DoLexerT__1       = 2
	DoLexerNEWLINE    = 3
	DoLexerWHITESPACE = 4
	DoLexerINT        = 5
)
