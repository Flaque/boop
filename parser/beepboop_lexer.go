// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 6,
	2, 15, 10, 2, 13, 2, 14, 2, 16, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 6, 4, 24,
	10, 4, 13, 4, 14, 4, 25, 3, 5, 3, 5, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 3, 2, 4, 4, 2, 12, 12, 15, 15, 3, 2, 50, 59, 2, 32, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 3, 14, 3, 2, 2, 2, 5, 18, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2,
	9, 27, 3, 2, 2, 2, 11, 29, 3, 2, 2, 2, 13, 15, 9, 2, 2, 2, 14, 13, 3, 2,
	2, 2, 15, 16, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 4,
	3, 2, 2, 2, 18, 19, 7, 34, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 8, 3, 2, 2,
	21, 6, 3, 2, 2, 2, 22, 24, 9, 3, 2, 2, 23, 22, 3, 2, 2, 2, 24, 25, 3, 2,
	2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 8, 3, 2, 2, 2, 27, 28,
	7, 45, 2, 2, 28, 10, 3, 2, 2, 2, 29, 30, 7, 47, 2, 2, 30, 12, 3, 2, 2,
	2, 5, 2, 16, 25, 3, 8, 2, 2,
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
	"", "", "' '", "", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS",
}

var lexerRuleNames = []string{
	"NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS",
}

type BeepBoopLexer struct {
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

func NewBeepBoopLexer(input antlr.CharStream) *BeepBoopLexer {

	l := new(BeepBoopLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "BeepBoop.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BeepBoopLexer tokens.
const (
	BeepBoopLexerNEWLINE    = 1
	BeepBoopLexerWHITESPACE = 2
	BeepBoopLexerINT        = 3
	BeepBoopLexerPLUS       = 4
	BeepBoopLexerMINUS      = 5
)
