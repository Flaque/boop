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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 43, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 6, 3, 23, 10, 3, 13, 3, 14,
	3, 24, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 32, 10, 5, 13, 5, 14, 5, 33,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 2, 2, 10, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2, 5, 4, 2, 12, 12, 15, 15, 3,
	2, 50, 59, 4, 2, 67, 92, 99, 124, 2, 44, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 22,
	3, 2, 2, 2, 7, 26, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13,
	37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2, 19, 20, 7, 38,
	2, 2, 20, 4, 3, 2, 2, 2, 21, 23, 9, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 24,
	3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 6, 3, 2, 2, 2,
	26, 27, 7, 34, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29, 8, 4, 2, 2, 29, 8, 3,
	2, 2, 2, 30, 32, 9, 3, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33,
	31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 45,
	2, 2, 36, 12, 3, 2, 2, 2, 37, 38, 7, 47, 2, 2, 38, 14, 3, 2, 2, 2, 39,
	40, 7, 63, 2, 2, 40, 16, 3, 2, 2, 2, 41, 42, 9, 4, 2, 2, 42, 18, 3, 2,
	2, 2, 5, 2, 24, 33, 3, 8, 2, 2,
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
	"", "'$'", "", "' '", "", "'+'", "'-'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "ASSIGN", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "ASSIGN", "STRING",
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
	BeepBoopLexerT__0       = 1
	BeepBoopLexerNEWLINE    = 2
	BeepBoopLexerWHITESPACE = 3
	BeepBoopLexerINT        = 4
	BeepBoopLexerPLUS       = 5
	BeepBoopLexerMINUS      = 6
	BeepBoopLexerASSIGN     = 7
	BeepBoopLexerSTRING     = 8
)
