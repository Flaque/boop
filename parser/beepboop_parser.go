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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 59, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 21, 10, 3, 13, 3, 14, 3, 22, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 31, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 41, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 46,
	10, 6, 12, 6, 14, 6, 49, 11, 6, 3, 7, 3, 7, 3, 7, 5, 7, 54, 10, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 2, 3, 10, 9, 2, 4, 6, 8, 10, 12, 14, 2, 3, 3, 2, 7, 8,
	2, 57, 2, 16, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 30, 3, 2, 2, 2, 8, 32,
	3, 2, 2, 2, 10, 40, 3, 2, 2, 2, 12, 53, 3, 2, 2, 2, 14, 55, 3, 2, 2, 2,
	16, 17, 5, 4, 3, 2, 17, 18, 7, 2, 2, 3, 18, 3, 3, 2, 2, 2, 19, 21, 5, 6,
	4, 2, 20, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23,
	3, 2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 25, 5, 10, 6, 2, 25, 26, 7, 4, 2, 2,
	26, 31, 3, 2, 2, 2, 27, 28, 5, 8, 5, 2, 28, 29, 7, 4, 2, 2, 29, 31, 3,
	2, 2, 2, 30, 24, 3, 2, 2, 2, 30, 27, 3, 2, 2, 2, 31, 7, 3, 2, 2, 2, 32,
	33, 5, 14, 8, 2, 33, 34, 7, 9, 2, 2, 34, 35, 5, 10, 6, 2, 35, 9, 3, 2,
	2, 2, 36, 37, 8, 6, 1, 2, 37, 41, 5, 12, 7, 2, 38, 39, 7, 8, 2, 2, 39,
	41, 5, 10, 6, 4, 40, 36, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 47, 3, 2,
	2, 2, 42, 43, 12, 3, 2, 2, 43, 44, 9, 2, 2, 2, 44, 46, 5, 10, 6, 4, 45,
	42, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2,
	2, 48, 11, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 54, 5, 14, 8, 2, 51, 54,
	7, 10, 2, 2, 52, 54, 7, 6, 2, 2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	53, 52, 3, 2, 2, 2, 54, 13, 3, 2, 2, 2, 55, 56, 7, 3, 2, 2, 56, 57, 7,
	10, 2, 2, 57, 15, 3, 2, 2, 2, 7, 22, 30, 40, 47, 53,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'$'", "", "' '", "", "'+'", "'-'", "'='",
}
var symbolicNames = []string{
	"", "", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "ASSIGN", "STRING",
}

var ruleNames = []string{
	"beepboop", "block", "statement", "assignment", "expr", "term", "label",
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
	BeepBoopParserNEWLINE    = 2
	BeepBoopParserWHITESPACE = 3
	BeepBoopParserINT        = 4
	BeepBoopParserPLUS       = 5
	BeepBoopParserMINUS      = 6
	BeepBoopParserASSIGN     = 7
	BeepBoopParserSTRING     = 8
)

// BeepBoopParser rules.
const (
	BeepBoopParserRULE_beepboop   = 0
	BeepBoopParserRULE_block      = 1
	BeepBoopParserRULE_statement  = 2
	BeepBoopParserRULE_assignment = 3
	BeepBoopParserRULE_expr       = 4
	BeepBoopParserRULE_term       = 5
	BeepBoopParserRULE_label      = 6
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

func (s *BeepboopContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BeepboopContext) EOF() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserEOF, 0)
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
		p.SetState(14)
		p.Block()
	}
	{
		p.SetState(15)
		p.Match(BeepBoopParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BeepBoopParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BeepBoopParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BeepBoopParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BeepBoopParserRULE_block)
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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BeepBoopParserT__0)|(1<<BeepBoopParserINT)|(1<<BeepBoopParserMINUS)|(1<<BeepBoopParserSTRING))) != 0) {
		{
			p.SetState(17)
			p.Statement()
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

type ExprStatementContext struct {
	*StatementContext
}

func NewExprStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStatementContext {
	var p = new(ExprStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExprStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(BeepBoopParserNEWLINE, 0)
}

func (s *ExprStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.EnterExprStatement(s)
	}
}

func (s *ExprStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BeepBoopListener); ok {
		listenerT.ExitExprStatement(s)
	}
}

func (s *ExprStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BeepBoopVisitor:
		return t.VisitExprStatement(s)

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

	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.expr(0)
		}
		{
			p.SetState(23)
			p.Match(BeepBoopParserNEWLINE)
		}

	case 2:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Assignment()
		}
		{
			p.SetState(26)
			p.Match(BeepBoopParserNEWLINE)
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
	p.EnterRule(localctx, 6, BeepBoopParserRULE_assignment)

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
		p.SetState(30)
		p.Label()
	}
	{
		p.SetState(31)
		p.Match(BeepBoopParserASSIGN)
	}
	{
		p.SetState(32)
		p.expr(0)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, BeepBoopParserRULE_expr, _p)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserT__0, BeepBoopParserINT, BeepBoopParserSTRING:
		localctx = NewTermExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(35)
			p.Term()
		}

	case BeepBoopParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.Match(BeepBoopParserMINUS)
		}
		{
			p.SetState(37)
			p.expr(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, BeepBoopParserRULE_expr)
			p.SetState(40)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(41)

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
				p.SetState(42)
				p.expr(2)
			}

		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 10, BeepBoopParserRULE_term)

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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BeepBoopParserT__0:
		localctx = NewLabelTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Label()
		}

	case BeepBoopParserSTRING:
		localctx = NewStringTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(BeepBoopParserSTRING)
		}

	case BeepBoopParserINT:
		localctx = NewIntTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
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
	p.EnterRule(localctx, 12, BeepBoopParserRULE_label)

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
		p.SetState(53)
		p.Match(BeepBoopParserT__0)
	}
	{
		p.SetState(54)
		p.Match(BeepBoopParserSTRING)
	}

	return localctx
}

func (p *BeepBoopParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
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
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
