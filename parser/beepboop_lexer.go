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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 37, 229,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3,
	2, 3, 2, 3, 3, 3, 3, 7, 3, 82, 10, 3, 12, 3, 14, 3, 85, 11, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 6, 4, 91, 10, 4, 13, 4, 14, 4, 92, 5, 4, 95, 10, 4, 3, 5,
	6, 5, 98, 10, 5, 13, 5, 14, 5, 99, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 33, 6, 33, 195, 10, 33, 13, 33, 14, 33, 196, 3, 34, 3, 34,
	6, 34, 201, 10, 34, 13, 34, 14, 34, 202, 7, 34, 205, 10, 34, 12, 34, 14,
	34, 208, 11, 34, 3, 34, 3, 34, 3, 35, 3, 35, 5, 35, 214, 10, 35, 3, 36,
	6, 36, 217, 10, 36, 13, 36, 14, 36, 218, 3, 37, 3, 37, 5, 37, 223, 10,
	37, 3, 38, 3, 38, 3, 38, 5, 38, 228, 10, 38, 2, 2, 39, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 2, 71, 36, 73, 37, 75, 2, 3, 2, 8, 4, 2, 12, 12,
	15, 15, 4, 2, 11, 11, 34, 34, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 4, 2,
	36, 36, 94, 94, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112,
	112, 116, 116, 118, 118, 2, 237, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2,
	53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2,
	2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 3, 77, 3, 2, 2, 2, 5, 79, 3, 2,
	2, 2, 7, 94, 3, 2, 2, 2, 9, 97, 3, 2, 2, 2, 11, 103, 3, 2, 2, 2, 13, 106,
	3, 2, 2, 2, 15, 108, 3, 2, 2, 2, 17, 110, 3, 2, 2, 2, 19, 112, 3, 2, 2,
	2, 21, 114, 3, 2, 2, 2, 23, 117, 3, 2, 2, 2, 25, 120, 3, 2, 2, 2, 27, 122,
	3, 2, 2, 2, 29, 124, 3, 2, 2, 2, 31, 126, 3, 2, 2, 2, 33, 128, 3, 2, 2,
	2, 35, 130, 3, 2, 2, 2, 37, 132, 3, 2, 2, 2, 39, 135, 3, 2, 2, 2, 41, 138,
	3, 2, 2, 2, 43, 142, 3, 2, 2, 2, 45, 147, 3, 2, 2, 2, 47, 154, 3, 2, 2,
	2, 49, 158, 3, 2, 2, 2, 51, 162, 3, 2, 2, 2, 53, 166, 3, 2, 2, 2, 55, 170,
	3, 2, 2, 2, 57, 175, 3, 2, 2, 2, 59, 180, 3, 2, 2, 2, 61, 186, 3, 2, 2,
	2, 63, 189, 3, 2, 2, 2, 65, 194, 3, 2, 2, 2, 67, 198, 3, 2, 2, 2, 69, 213,
	3, 2, 2, 2, 71, 216, 3, 2, 2, 2, 73, 222, 3, 2, 2, 2, 75, 224, 3, 2, 2,
	2, 77, 78, 7, 60, 2, 2, 78, 4, 3, 2, 2, 2, 79, 83, 7, 37, 2, 2, 80, 82,
	10, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 8,
	3, 2, 2, 87, 6, 3, 2, 2, 2, 88, 95, 9, 2, 2, 2, 89, 91, 9, 2, 2, 2, 90,
	89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2,
	2, 93, 95, 3, 2, 2, 2, 94, 88, 3, 2, 2, 2, 94, 90, 3, 2, 2, 2, 95, 8, 3,
	2, 2, 2, 96, 98, 9, 3, 2, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99,
	97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 8,
	5, 3, 2, 102, 10, 3, 2, 2, 2, 103, 104, 7, 63, 2, 2, 104, 105, 7, 63, 2,
	2, 105, 12, 3, 2, 2, 2, 106, 107, 7, 63, 2, 2, 107, 14, 3, 2, 2, 2, 108,
	109, 7, 126, 2, 2, 109, 16, 3, 2, 2, 2, 110, 111, 7, 62, 2, 2, 111, 18,
	3, 2, 2, 2, 112, 113, 7, 64, 2, 2, 113, 20, 3, 2, 2, 2, 114, 115, 7, 62,
	2, 2, 115, 116, 7, 63, 2, 2, 116, 22, 3, 2, 2, 2, 117, 118, 7, 64, 2, 2,
	118, 119, 7, 63, 2, 2, 119, 24, 3, 2, 2, 2, 120, 121, 7, 42, 2, 2, 121,
	26, 3, 2, 2, 2, 122, 123, 7, 43, 2, 2, 123, 28, 3, 2, 2, 2, 124, 125, 7,
	125, 2, 2, 125, 30, 3, 2, 2, 2, 126, 127, 7, 127, 2, 2, 127, 32, 3, 2,
	2, 2, 128, 129, 7, 93, 2, 2, 129, 34, 3, 2, 2, 2, 130, 131, 7, 95, 2, 2,
	131, 36, 3, 2, 2, 2, 132, 133, 7, 107, 2, 2, 133, 134, 7, 104, 2, 2, 134,
	38, 3, 2, 2, 2, 135, 136, 7, 102, 2, 2, 136, 137, 7, 113, 2, 2, 137, 40,
	3, 2, 2, 2, 138, 139, 7, 103, 2, 2, 139, 140, 7, 112, 2, 2, 140, 141, 7,
	102, 2, 2, 141, 42, 3, 2, 2, 2, 142, 143, 7, 104, 2, 2, 143, 144, 7, 119,
	2, 2, 144, 145, 7, 112, 2, 2, 145, 146, 7, 101, 2, 2, 146, 44, 3, 2, 2,
	2, 147, 148, 7, 116, 2, 2, 148, 149, 7, 103, 2, 2, 149, 150, 7, 118, 2,
	2, 150, 151, 7, 119, 2, 2, 151, 152, 7, 116, 2, 2, 152, 153, 7, 112, 2,
	2, 153, 46, 3, 2, 2, 2, 154, 155, 7, 104, 2, 2, 155, 156, 7, 113, 2, 2,
	156, 157, 7, 116, 2, 2, 157, 48, 3, 2, 2, 2, 158, 159, 7, 99, 2, 2, 159,
	160, 7, 102, 2, 2, 160, 161, 7, 102, 2, 2, 161, 50, 3, 2, 2, 2, 162, 163,
	7, 117, 2, 2, 163, 164, 7, 119, 2, 2, 164, 165, 7, 100, 2, 2, 165, 52,
	3, 2, 2, 2, 166, 167, 7, 102, 2, 2, 167, 168, 7, 107, 2, 2, 168, 169, 7,
	120, 2, 2, 169, 54, 3, 2, 2, 2, 170, 171, 7, 111, 2, 2, 171, 172, 7, 119,
	2, 2, 172, 173, 7, 110, 2, 2, 173, 174, 7, 118, 2, 2, 174, 56, 3, 2, 2,
	2, 175, 176, 7, 118, 2, 2, 176, 177, 7, 116, 2, 2, 177, 178, 7, 119, 2,
	2, 178, 179, 7, 103, 2, 2, 179, 58, 3, 2, 2, 2, 180, 181, 7, 104, 2, 2,
	181, 182, 7, 99, 2, 2, 182, 183, 7, 110, 2, 2, 183, 184, 7, 117, 2, 2,
	184, 185, 7, 103, 2, 2, 185, 60, 3, 2, 2, 2, 186, 187, 7, 113, 2, 2, 187,
	188, 7, 116, 2, 2, 188, 62, 3, 2, 2, 2, 189, 190, 7, 99, 2, 2, 190, 191,
	7, 112, 2, 2, 191, 192, 7, 102, 2, 2, 192, 64, 3, 2, 2, 2, 193, 195, 9,
	4, 2, 2, 194, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 194, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 66, 3, 2, 2, 2, 198, 206, 7, 36, 2, 2, 199,
	201, 5, 69, 35, 2, 200, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 200,
	3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204, 200, 3, 2,
	2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2,
	207, 209, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 7, 36, 2, 2, 210,
	68, 3, 2, 2, 2, 211, 214, 5, 73, 37, 2, 212, 214, 5, 7, 4, 2, 213, 211,
	3, 2, 2, 2, 213, 212, 3, 2, 2, 2, 214, 70, 3, 2, 2, 2, 215, 217, 9, 5,
	2, 2, 216, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2,
	218, 219, 3, 2, 2, 2, 219, 72, 3, 2, 2, 2, 220, 223, 5, 75, 38, 2, 221,
	223, 10, 6, 2, 2, 222, 220, 3, 2, 2, 2, 222, 221, 3, 2, 2, 2, 223, 74,
	3, 2, 2, 2, 224, 227, 7, 94, 2, 2, 225, 228, 9, 7, 2, 2, 226, 228, 5, 71,
	36, 2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2, 2, 2, 228, 76, 3, 2, 2, 2,
	14, 2, 83, 92, 94, 99, 196, 202, 206, 213, 218, 222, 227, 4, 8, 2, 2, 2,
	3, 2,
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
	"", "':'", "", "", "", "'=='", "'='", "'|'", "'<'", "'>'", "'<='", "'>='",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "'if'", "'do'", "'end'", "'func'",
	"'return'", "'for'", "'add'", "'sub'", "'div'", "'mult'", "'true'", "'false'",
	"'or'", "'and'",
}

var lexerSymbolicNames = []string{
	"", "", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE",
	"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", "LSQUIG",
	"RSQUIG", "LBLOCK", "RBLOCK", "IF", "DO", "END", "FUNC", "RETURN", "FOR",
	"PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", "QUOTED",
	"LETTERS", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "EQUALS", "ASSIGN", "PIPE",
	"LTHAN", "GTHAN", "LTHAN_EQUALS", "GTHAN_EQUALS", "LPAREN", "RPAREN", "LSQUIG",
	"RSQUIG", "LBLOCK", "RBLOCK", "IF", "DO", "END", "FUNC", "RETURN", "FOR",
	"PLUS", "SUB", "DIV", "MULT", "TRUE", "FALSE", "OR", "AND", "INT", "QUOTED",
	"STRINGORNEW", "LETTERS", "STRING", "ESC",
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
	BeepBoopLexerT__0         = 1
	BeepBoopLexerCOMMENT      = 2
	BeepBoopLexerNEWLINE      = 3
	BeepBoopLexerWHITESPACE   = 4
	BeepBoopLexerEQUALS       = 5
	BeepBoopLexerASSIGN       = 6
	BeepBoopLexerPIPE         = 7
	BeepBoopLexerLTHAN        = 8
	BeepBoopLexerGTHAN        = 9
	BeepBoopLexerLTHAN_EQUALS = 10
	BeepBoopLexerGTHAN_EQUALS = 11
	BeepBoopLexerLPAREN       = 12
	BeepBoopLexerRPAREN       = 13
	BeepBoopLexerLSQUIG       = 14
	BeepBoopLexerRSQUIG       = 15
	BeepBoopLexerLBLOCK       = 16
	BeepBoopLexerRBLOCK       = 17
	BeepBoopLexerIF           = 18
	BeepBoopLexerDO           = 19
	BeepBoopLexerEND          = 20
	BeepBoopLexerFUNC         = 21
	BeepBoopLexerRETURN       = 22
	BeepBoopLexerFOR          = 23
	BeepBoopLexerPLUS         = 24
	BeepBoopLexerSUB          = 25
	BeepBoopLexerDIV          = 26
	BeepBoopLexerMULT         = 27
	BeepBoopLexerTRUE         = 28
	BeepBoopLexerFALSE        = 29
	BeepBoopLexerOR           = 30
	BeepBoopLexerAND          = 31
	BeepBoopLexerINT          = 32
	BeepBoopLexerQUOTED       = 33
	BeepBoopLexerLETTERS      = 34
	BeepBoopLexerSTRING       = 35
)
