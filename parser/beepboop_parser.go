// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 173,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 30, 10, 2, 13, 2, 14, 2, 31, 3, 2, 6,
	2, 35, 10, 2, 13, 2, 14, 2, 36, 3, 2, 3, 2, 3, 2, 6, 2, 42, 10, 2, 13,
	2, 14, 2, 43, 3, 2, 3, 2, 5, 2, 48, 10, 2, 3, 3, 3, 3, 5, 3, 52, 10, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 67, 10, 4, 3, 5, 6, 5, 70, 10, 5, 13, 5, 14, 5, 71, 3, 6, 3,
	6, 3, 6, 6, 6, 77, 10, 6, 13, 6, 14, 6, 78, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 5, 6, 98, 10, 6, 3, 7, 3, 7, 3, 7, 6, 7, 103, 10, 7, 13, 7, 14, 7, 104,
	5, 7, 107, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 121, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 127,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 137,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 143, 10, 11, 3, 11, 3, 11, 3,
	11, 7, 11, 148, 10, 11, 12, 11, 14, 11, 151, 11, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 6, 12, 159, 10, 12, 13, 12, 14, 12, 160, 5, 12, 163,
	10, 12, 3, 13, 3, 13, 3, 13, 5, 13, 168, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 2, 3, 20, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 3,
	3, 2, 14, 15, 2, 183, 2, 47, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 66, 3, 2,
	2, 2, 8, 69, 3, 2, 2, 2, 10, 97, 3, 2, 2, 2, 12, 106, 3, 2, 2, 2, 14, 120,
	3, 2, 2, 2, 16, 126, 3, 2, 2, 2, 18, 136, 3, 2, 2, 2, 20, 142, 3, 2, 2,
	2, 22, 162, 3, 2, 2, 2, 24, 167, 3, 2, 2, 2, 26, 169, 3, 2, 2, 2, 28, 30,
	7, 5, 2, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2,
	31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 35, 5, 4, 3, 2, 34, 33, 3,
	2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 39, 7, 2, 2, 3, 39, 48, 3, 2, 2, 2, 40, 42, 5, 4, 3,
	2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 7, 2, 2, 3, 46, 48, 3, 2, 2, 2,
	47, 29, 3, 2, 2, 2, 47, 41, 3, 2, 2, 2, 48, 3, 3, 2, 2, 2, 49, 52, 5, 6,
	4, 2, 50, 52, 5, 10, 6, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	5, 3, 2, 2, 2, 53, 54, 5, 12, 7, 2, 54, 55, 7, 5, 2, 2, 55, 67, 3, 2, 2,
	2, 56, 57, 5, 18, 10, 2, 57, 58, 7, 5, 2, 2, 58, 67, 3, 2, 2, 2, 59, 60,
	5, 16, 9, 2, 60, 61, 7, 5, 2, 2, 61, 67, 3, 2, 2, 2, 62, 63, 5, 22, 12,
	2, 63, 64, 7, 5, 2, 2, 64, 67, 3, 2, 2, 2, 65, 67, 7, 5, 2, 2, 66, 53,
	3, 2, 2, 2, 66, 56, 3, 2, 2, 2, 66, 59, 3, 2, 2, 2, 66, 62, 3, 2, 2, 2,
	66, 65, 3, 2, 2, 2, 67, 7, 3, 2, 2, 2, 68, 70, 5, 6, 4, 2, 69, 68, 3, 2,
	2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 9,
	3, 2, 2, 2, 73, 74, 7, 10, 2, 2, 74, 76, 7, 23, 2, 2, 75, 77, 5, 26, 14,
	2, 76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79,
	3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 7, 8, 2, 2, 81, 82, 5, 8, 5, 2,
	82, 83, 7, 9, 2, 2, 83, 84, 7, 5, 2, 2, 84, 98, 3, 2, 2, 2, 85, 86, 7,
	10, 2, 2, 86, 87, 7, 23, 2, 2, 87, 88, 7, 8, 2, 2, 88, 89, 5, 8, 5, 2,
	89, 90, 7, 9, 2, 2, 90, 91, 7, 5, 2, 2, 91, 98, 3, 2, 2, 2, 92, 93, 7,
	10, 2, 2, 93, 94, 7, 23, 2, 2, 94, 95, 7, 8, 2, 2, 95, 96, 7, 9, 2, 2,
	96, 98, 7, 5, 2, 2, 97, 73, 3, 2, 2, 2, 97, 85, 3, 2, 2, 2, 97, 92, 3,
	2, 2, 2, 98, 11, 3, 2, 2, 2, 99, 107, 7, 23, 2, 2, 100, 102, 7, 23, 2,
	2, 101, 103, 5, 24, 13, 2, 102, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2,
	104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106,
	99, 3, 2, 2, 2, 106, 100, 3, 2, 2, 2, 107, 13, 3, 2, 2, 2, 108, 109, 7,
	7, 2, 2, 109, 110, 5, 20, 11, 2, 110, 111, 7, 8, 2, 2, 111, 112, 5, 4,
	3, 2, 112, 113, 7, 9, 2, 2, 113, 121, 3, 2, 2, 2, 114, 115, 7, 7, 2, 2,
	115, 116, 5, 12, 7, 2, 116, 117, 7, 8, 2, 2, 117, 118, 5, 4, 3, 2, 118,
	119, 7, 9, 2, 2, 119, 121, 3, 2, 2, 2, 120, 108, 3, 2, 2, 2, 120, 114,
	3, 2, 2, 2, 121, 15, 3, 2, 2, 2, 122, 123, 7, 11, 2, 2, 123, 127, 5, 20,
	11, 2, 124, 125, 7, 11, 2, 2, 125, 127, 5, 12, 7, 2, 126, 122, 3, 2, 2,
	2, 126, 124, 3, 2, 2, 2, 127, 17, 3, 2, 2, 2, 128, 129, 5, 26, 14, 2, 129,
	130, 7, 18, 2, 2, 130, 131, 5, 20, 11, 2, 131, 137, 3, 2, 2, 2, 132, 133,
	5, 26, 14, 2, 133, 134, 7, 18, 2, 2, 134, 135, 5, 12, 7, 2, 135, 137, 3,
	2, 2, 2, 136, 128, 3, 2, 2, 2, 136, 132, 3, 2, 2, 2, 137, 19, 3, 2, 2,
	2, 138, 139, 8, 11, 1, 2, 139, 140, 7, 15, 2, 2, 140, 143, 5, 20, 11, 4,
	141, 143, 5, 24, 13, 2, 142, 138, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143,
	149, 3, 2, 2, 2, 144, 145, 12, 5, 2, 2, 145, 146, 9, 2, 2, 2, 146, 148,
	5, 20, 11, 6, 147, 144, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3,
	2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 21, 3, 2, 2, 2, 151, 149, 3, 2, 2,
	2, 152, 153, 5, 12, 7, 2, 153, 154, 7, 19, 2, 2, 154, 155, 5, 12, 7, 2,
	155, 163, 3, 2, 2, 2, 156, 158, 5, 12, 7, 2, 157, 159, 7, 5, 2, 2, 158,
	157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161,
	3, 2, 2, 2, 161, 163, 3, 2, 2, 2, 162, 152, 3, 2, 2, 2, 162, 156, 3, 2,
	2, 2, 163, 23, 3, 2, 2, 2, 164, 168, 5, 26, 14, 2, 165, 168, 7, 23, 2,
	2, 166, 168, 7, 22, 2, 2, 167, 164, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167,
	166, 3, 2, 2, 2, 168, 25, 3, 2, 2, 2, 169, 170, 7, 3, 2, 2, 170, 171, 7,
	23, 2, 2, 171, 27, 3, 2, 2, 2, 21, 31, 36, 43, 47, 51, 66, 71, 78, 97,
	104, 106, 120, 126, 136, 142, 149, 160, 162, 167,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'$'", "", "", "", "'if'", "'do'", "'end'", "'func'", "'return'", "'for'",
	"'is'", "'+'", "'-'", "'*'", "'/'", "'='", "'|'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", "FUNC",
	"RETURN", "FOR", "IS", "PLUS", "MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE",
	"LPAREN", "RPAREN", "INT", "STRING",
}

var ruleNames = []string{
	"beepboop", "code", "statement", "funcguts", "funcdef", "fncall", "ifstat",
	"returnStat", "assignment", "expr", "pipe", "term", "label",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BeepBoopParser struct {
	*antlr.BaseParser
}

func NewBeepBoopParser(input antlr.TokenStream) *BeepBoopParser {
	this := new(BeepBoopParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BeepBoop.g4"

	return this
}

// BeepBoopParser tokens.
const (
	BeepBoopParserEOF        = antlr.TokenEOF
	BeepBoopParserT__0       = 1
	BeepBoopParserCOMMENT    = 2
	BeepBoopParserNEWLINE    = 3
	BeepBoopParserWHITESPACE = 4
	BeepBoopParserIF         = 5
	BeepBoopParserDO         = 6
	BeepBoopParserEND        = 7
	BeepBoopParserFUNC       = 8
	BeepBoopParserRETURN     = 9
	BeepBoopParserFOR        = 10
	BeepBoopParserIS         = 11
	BeepBoopParserPLUS       = 12
	BeepBoopParserMINUS      = 13
	BeepBoopParserMULT       = 14
	BeepBoopParserDIVIDE     = 15
	BeepBoopParserASSIGN     = 16
	BeepBoopParserPIPE       = 17
	BeepBoopParserLPAREN     = 18
	BeepBoopParserRPAREN     = 19
	BeepBoopParserINT        = 20
	BeepBoopParserSTRING     = 21
)

// BeepBoopParser rules.
const (
	BeepBoopParserRULE_beepboop   = 0
	BeepBoopParserRULE_code       = 1
	BeepBoopParserRULE_statement  = 2
	BeepBoopParserRULE_funcguts   = 3
	BeepBoopParserRULE_funcdef    = 4
	BeepBoopParserRULE_fncall     = 5
	BeepBoopParserRULE_ifstat     = 6
	BeepBoopParserRULE_returnStat = 7
	BeepBoopParserRULE_assignment = 8
	BeepBoopParserRULE_expr       = 9
	BeepBoopParserRULE_pipe       = 10
	BeepBoopParserRULE_term       = 11
	BeepBoopParserRULE_label      = 12
)

// IBeepboopContext is an interface to support dynamic dispatch.
type IBeepboopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeepboopContext differentiates from other interfaces.
	IsBeepboopContext()
}

type BeepboopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeepboopContext() *BeepboopContext {
	var p = new(BeepboopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_beepboop
	return p
}

func (*BeepboopContext) IsBeepboopContext() {}

func NewBeepboopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeepboopContext {
	var p = new(BeepboopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_beepboop

	return p
}

func (s *BeepboopContext) GetParser() antlr.Parser { return s.parser }

func (s *BeepboopContext) EOF() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEOF, 0)
}

func (s *BeepboopContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(BeepBoopParserNEWLINE)
}

func (s *BeepboopContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, i)
}

func (s *BeepboopContext) AllCode() []ICodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICodeContext)(nil)).Elem())
	var tst = make([]ICodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICodeContext)
		}
	}

	return tst
}

func (s *BeepboopContext) Code(i int) ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *BeepboopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeepboopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeepboopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterBeepboop(s)
	}
}

func (s *BeepboopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitBeepboop(s)
	}
}

func (s *BeepboopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitBeepboop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Beepboop() (localctx IBeepboopContext) {
	localctx = NewBeepboopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BeepBoopParserRULE_beepboop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(26)
					p.Match(BeepBoopParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserFUNC)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSTRING))) != 0) {
			{
				p.SetState(31)
				p.Code()
			}

			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(36)
			p.Match(BeepBoopParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserFUNC)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSTRING))) != 0) {
			{
				p.SetState(38)
				p.Code()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(43)
			p.Match(BeepBoopParserEOF)
		}

	}

	return localctx
}

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_code
	return p
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) CopyFrom(ctx *CodeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementCodeContext struct {
	*CodeContext
}

func NewStatementCodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementCodeContext {
	var p = new(StatementCodeContext)

	p.CodeContext = NewEmptyCodeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CodeContext))

	return p
}

func (s *StatementCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementCodeContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterStatementCode(s)
	}
}

func (s *StatementCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitStatementCode(s)
	}
}

func (s *StatementCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitStatementCode(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncdefCodeContext struct {
	*CodeContext
}

func NewFuncdefCodeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncdefCodeContext {
	var p = new(FuncdefCodeContext)

	p.CodeContext = NewEmptyCodeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CodeContext))

	return p
}

func (s *FuncdefCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefCodeContext) Funcdef() IFuncdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncdefContext)
}

func (s *FuncdefCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFuncdefCode(s)
	}
}

func (s *FuncdefCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFuncdefCode(s)
	}
}

func (s *FuncdefCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFuncdefCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BeepBoopParserRULE_code)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserT__0, BeepBoopParserNEWLINE, BeepBoopParserRETURN, BeepBoopParserSTRING:
		localctx = NewStatementCodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Statement()
		}

	case BeepBoopParserFUNC:
		localctx = NewFuncdefCodeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Funcdef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PipeStatementContext struct {
	*StatementContext
}

func NewPipeStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeStatementContext {
	var p = new(PipeStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PipeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeStatementContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *PipeStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *PipeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterPipeStatement(s)
	}
}

func (s *PipeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitPipeStatement(s)
	}
}

func (s *PipeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitPipeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStatementContext struct {
	*StatementContext
}

func NewAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatementContext {
	var p = new(AssignStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type NoopStatementContext struct {
	*StatementContext
}

func NewNoopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NoopStatementContext {
	var p = new(NoopStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoopStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *NoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterNoopStatement(s)
	}
}

func (s *NoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitNoopStatement(s)
	}
}

func (s *NoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitNoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	*StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ReturnStat() IReturnStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatContext)
}

func (s *ReturnStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FncallStatementContext struct {
	*StatementContext
}

func NewFncallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FncallStatementContext {
	var p = new(FncallStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FncallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallStatementContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *FncallStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *FncallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncallStatement(s)
	}
}

func (s *FncallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncallStatement(s)
	}
}

func (s *FncallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BeepBoopParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFncallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Fncall()
		}
		{
			p.SetState(52)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 2:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Assignment()
		}
		{
			p.SetState(55)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 3:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.ReturnStat()
		}
		{
			p.SetState(58)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 4:
		localctx = NewPipeStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(60)
			p.Pipe()
		}
		{
			p.SetState(61)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 5:
		localctx = NewNoopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Match(BeepBoopParserNEWLINE)
		}

	}

	return localctx
}

// IFuncgutsContext is an interface to support dynamic dispatch.
type IFuncgutsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncgutsContext differentiates from other interfaces.
	IsFuncgutsContext()
}

type FuncgutsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncgutsContext() *FuncgutsContext {
	var p = new(FuncgutsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_funcguts
	return p
}

func (*FuncgutsContext) IsFuncgutsContext() {}

func NewFuncgutsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncgutsContext {
	var p = new(FuncgutsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_funcguts

	return p
}

func (s *FuncgutsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncgutsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FuncgutsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FuncgutsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncgutsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncgutsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFuncguts(s)
	}
}

func (s *FuncgutsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFuncguts(s)
	}
}

func (s *FuncgutsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFuncguts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Funcguts() (localctx IFuncgutsContext) {
	localctx = NewFuncgutsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BeepBoopParserRULE_funcguts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserNEWLINE)|(1<<BeepBoopParserRETURN)|(1<<BeepBoopParserSTRING))) != 0) {
		{
			p.SetState(66)
			p.Statement()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncdefContext is an interface to support dynamic dispatch.
type IFuncdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncdefContext differentiates from other interfaces.
	IsFuncdefContext()
}

type FuncdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdefContext() *FuncdefContext {
	var p = new(FuncdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_funcdef
	return p
}

func (*FuncdefContext) IsFuncdefContext() {}

func NewFuncdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdefContext {
	var p = new(FuncdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_funcdef

	return p
}

func (s *FuncdefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserFUNC, 0)
}

func (s *FuncdefContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSTRING, 0)
}

func (s *FuncdefContext) DO() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDO, 0)
}

func (s *FuncdefContext) Funcguts() IFuncgutsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncgutsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncgutsContext)
}

func (s *FuncdefContext) END() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEND, 0)
}

func (s *FuncdefContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *FuncdefContext) AllLabel() []ILabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILabelContext)(nil)).Elem())
	var tst = make([]ILabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILabelContext)
		}
	}

	return tst
}

func (s *FuncdefContext) Label(i int) ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *FuncdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFuncdef(s)
	}
}

func (s *FuncdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFuncdef(s)
	}
}

func (s *FuncdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFuncdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Funcdef() (localctx IFuncdefContext) {
	localctx = NewFuncdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BeepBoopParserRULE_funcdef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(72)
			p.Match(BeepBoopParserSTRING)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == BeepBoopParserT__0 {
			{
				p.SetState(73)
				p.Label()
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(78)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(79)
			p.Funcguts()
		}
		{
			p.SetState(80)
			p.Match(BeepBoopParserEND)
		}
		{
			p.SetState(81)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(84)
			p.Match(BeepBoopParserSTRING)
		}
		{
			p.SetState(85)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(86)
			p.Funcguts()
		}
		{
			p.SetState(87)
			p.Match(BeepBoopParserEND)
		}
		{
			p.SetState(88)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(BeepBoopParserFUNC)
		}
		{
			p.SetState(91)
			p.Match(BeepBoopParserSTRING)
		}
		{
			p.SetState(92)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(93)
			p.Match(BeepBoopParserEND)
		}
		{
			p.SetState(94)
			p.Match(BeepBoopParserNEWLINE)
		}

	}

	return localctx
}

// IFncallContext is an interface to support dynamic dispatch.
type IFncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFncallContext differentiates from other interfaces.
	IsFncallContext()
}

type FncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFncallContext() *FncallContext {
	var p = new(FncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_fncall
	return p
}

func (*FncallContext) IsFncallContext() {}

func NewFncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FncallContext {
	var p = new(FncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_fncall

	return p
}

func (s *FncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FncallContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSTRING, 0)
}

func (s *FncallContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FncallContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncall(s)
	}
}

func (s *FncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncall(s)
	}
}

func (s *FncallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Fncall() (localctx IFncallContext) {
	localctx = NewFncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BeepBoopParserRULE_fncall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(BeepBoopParserSTRING)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(BeepBoopParserSTRING)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserINT)|(1<<BeepBoopParserSTRING))) != 0) {
			{
				p.SetState(99)
				p.Term()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) CopyFrom(ctx *IfstatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprIfStatementContext struct {
	*IfstatContext
}

func NewExprIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIfStatementContext {
	var p = new(ExprIfStatementContext)

	p.IfstatContext = NewEmptyIfstatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IfstatContext))

	return p
}

func (s *ExprIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserIF, 0)
}

func (s *ExprIfStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprIfStatementContext) DO() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDO, 0)
}

func (s *ExprIfStatementContext) Code() ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *ExprIfStatementContext) END() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEND, 0)
}

func (s *ExprIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterExprIfStatement(s)
	}
}

func (s *ExprIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitExprIfStatement(s)
	}
}

func (s *ExprIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitExprIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FncallIfStatementContext struct {
	*IfstatContext
}

func NewFncallIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FncallIfStatementContext {
	var p = new(FncallIfStatementContext)

	p.IfstatContext = NewEmptyIfstatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IfstatContext))

	return p
}

func (s *FncallIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallIfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserIF, 0)
}

func (s *FncallIfStatementContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *FncallIfStatementContext) DO() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserDO, 0)
}

func (s *FncallIfStatementContext) Code() ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *FncallIfStatementContext) END() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEND, 0)
}

func (s *FncallIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncallIfStatement(s)
	}
}

func (s *FncallIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncallIfStatement(s)
	}
}

func (s *FncallIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncallIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BeepBoopParserRULE_ifstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(BeepBoopParserIF)
		}
		{
			p.SetState(107)
			p.expr(0)
		}
		{
			p.SetState(108)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(109)
			p.Code()
		}
		{
			p.SetState(110)
			p.Match(BeepBoopParserEND)
		}

	case 2:
		localctx = NewFncallIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(BeepBoopParserIF)
		}
		{
			p.SetState(113)
			p.Fncall()
		}
		{
			p.SetState(114)
			p.Match(BeepBoopParserDO)
		}
		{
			p.SetState(115)
			p.Code()
		}
		{
			p.SetState(116)
			p.Match(BeepBoopParserEND)
		}

	}

	return localctx
}

// IReturnStatContext is an interface to support dynamic dispatch.
type IReturnStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatContext differentiates from other interfaces.
	IsReturnStatContext()
}

type ReturnStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatContext() *ReturnStatContext {
	var p = new(ReturnStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_returnStat
	return p
}

func (*ReturnStatContext) IsReturnStatContext() {}

func NewReturnStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatContext {
	var p = new(ReturnStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_returnStat

	return p
}

func (s *ReturnStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatContext) CopyFrom(ctx *ReturnStatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ReturnStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FncallReturnContext struct {
	*ReturnStatContext
}

func NewFncallReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FncallReturnContext {
	var p = new(FncallReturnContext)

	p.ReturnStatContext = NewEmptyReturnStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReturnStatContext))

	return p
}

func (s *FncallReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FncallReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserRETURN, 0)
}

func (s *FncallReturnContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *FncallReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterFncallReturn(s)
	}
}

func (s *FncallReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitFncallReturn(s)
	}
}

func (s *FncallReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitFncallReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprReturnContext struct {
	*ReturnStatContext
}

func NewExprReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprReturnContext {
	var p = new(ExprReturnContext)

	p.ReturnStatContext = NewEmptyReturnStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ReturnStatContext))

	return p
}

func (s *ExprReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserRETURN, 0)
}

func (s *ExprReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterExprReturn(s)
	}
}

func (s *ExprReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitExprReturn(s)
	}
}

func (s *ExprReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitExprReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) ReturnStat() (localctx IReturnStatContext) {
	localctx = NewReturnStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BeepBoopParserRULE_returnStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(BeepBoopParserRETURN)
		}
		{
			p.SetState(121)
			p.expr(0)
		}

	case 2:
		localctx = NewFncallReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(BeepBoopParserRETURN)
		}
		{
			p.SetState(123)
			p.Fncall()
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserASSIGN, 0)
}

func (s *AssignmentContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) Fncall() IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BeepBoopParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Label()
		}
		{
			p.SetState(127)
			p.Match(BeepBoopParserASSIGN)
		}
		{
			p.SetState(128)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Label()
		}
		{
			p.SetState(131)
			p.Match(BeepBoopParserASSIGN)
		}
		{
			p.SetState(132)
			p.Fncall()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryMinusExprContext struct {
	*ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserMINUS, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitUnaryMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TermExprContext struct {
	*ExprContext
}

func NewTermExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermExprContext {
	var p = new(TermExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TermExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermExprContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterTermExpr(s)
	}
}

func (s *TermExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitTermExpr(s)
	}
}

func (s *TermExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitTermExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AdditiveExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AdditiveExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPLUS, 0)
}

func (s *AdditiveExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserMINUS, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *BeepBoopParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, BeepBoopParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(137)
			p.Match(BeepBoopParserMINUS)
		}
		{
			p.SetState(138)
			p.expr(2)
		}

	case BeepBoopParserT__0, BeepBoopParserINT, BeepBoopParserSTRING:
		localctx = NewTermExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(139)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_expr)
			p.SetState(142)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(143)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AdditiveExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == BeepBoopParserPLUS || _la == BeepBoopParserMINUS) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AdditiveExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(144)
				p.expr(4)
			}

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IPipeContext is an interface to support dynamic dispatch.
type IPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeContext differentiates from other interfaces.
	IsPipeContext()
}

type PipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeContext() *PipeContext {
	var p = new(PipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_pipe
	return p
}

func (*PipeContext) IsPipeContext() {}

func NewPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeContext {
	var p = new(PipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_pipe

	return p
}

func (s *PipeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeContext) AllFncall() []IFncallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFncallContext)(nil)).Elem())
	var tst = make([]IFncallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFncallContext)
		}
	}

	return tst
}

func (s *PipeContext) Fncall(i int) IFncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFncallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFncallContext)
}

func (s *PipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserPIPE, 0)
}

func (s *PipeContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(BeepBoopParserNEWLINE)
}

func (s *PipeContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, i)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (s *PipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Pipe() (localctx IPipeContext) {
	localctx = NewPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BeepBoopParserRULE_pipe)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Fncall()
		}
		{
			p.SetState(151)
			p.Match(BeepBoopParserPIPE)
		}
		{
			p.SetState(152)
			p.Fncall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Fncall()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(155)
					p.Match(BeepBoopParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyFrom(ctx *TermContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringTermContext struct {
	*TermContext
}

func NewStringTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringTermContext {
	var p = new(StringTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *StringTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTermContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSTRING, 0)
}

func (s *StringTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterStringTerm(s)
	}
}

func (s *StringTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitStringTerm(s)
	}
}

func (s *StringTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitStringTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type LabelTermContext struct {
	*TermContext
}

func NewLabelTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelTermContext {
	var p = new(LabelTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *LabelTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelTermContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LabelTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLabelTerm(s)
	}
}

func (s *LabelTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLabelTerm(s)
	}
}

func (s *LabelTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLabelTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntTermContext struct {
	*TermContext
}

func NewIntTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTermContext {
	var p = new(IntTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *IntTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTermContext) INT() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserINT, 0)
}

func (s *IntTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterIntTerm(s)
	}
}

func (s *IntTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitIntTerm(s)
	}
}

func (s *IntTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitIntTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BeepBoopParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserT__0:
		localctx = NewLabelTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Label()
		}

	case BeepBoopParserSTRING:
		localctx = NewStringTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(BeepBoopParserSTRING)
		}

	case BeepBoopParserINT:
		localctx = NewIntTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
			p.Match(BeepBoopParserINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) STRING() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserSTRING, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BeepBoopParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(BeepBoopParserT__0)
	}
	{
		p.SetState(168)
		p.Match(BeepBoopParserSTRING)
	}

	return localctx
}

func (p *BeepBoopParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BeepBoopParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
