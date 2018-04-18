// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBeepBoopListener is a complete listener for a parse tree produced by BeepBoopParser.
type BaseBeepBoopListener struct{}

var _ BeepBoopListener = &BaseBeepBoopListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBeepBoopListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBeepBoopListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBeepBoopListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBeepBoopListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBeepboop is called when production beepboop is entered.
func (s *BaseBeepBoopListener) EnterBeepboop(ctx *BeepboopContext) {}

// ExitBeepboop is called when production beepboop is exited.
func (s *BaseBeepBoopListener) ExitBeepboop(ctx *BeepboopContext) {}

// EnterStatementCode is called when production statementCode is entered.
func (s *BaseBeepBoopListener) EnterStatementCode(ctx *StatementCodeContext) {}

// ExitStatementCode is called when production statementCode is exited.
func (s *BaseBeepBoopListener) ExitStatementCode(ctx *StatementCodeContext) {}

// EnterFuncdefCode is called when production funcdefCode is entered.
func (s *BaseBeepBoopListener) EnterFuncdefCode(ctx *FuncdefCodeContext) {}

// ExitFuncdefCode is called when production funcdefCode is exited.
func (s *BaseBeepBoopListener) ExitFuncdefCode(ctx *FuncdefCodeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBeepBoopListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBeepBoopListener) ExitBlock(ctx *BlockContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseBeepBoopListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseBeepBoopListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterFncallStatement is called when production fncallStatement is entered.
func (s *BaseBeepBoopListener) EnterFncallStatement(ctx *FncallStatementContext) {}

// ExitFncallStatement is called when production fncallStatement is exited.
func (s *BaseBeepBoopListener) ExitFncallStatement(ctx *FncallStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseBeepBoopListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseBeepBoopListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterPipeStatement is called when production pipeStatement is entered.
func (s *BaseBeepBoopListener) EnterPipeStatement(ctx *PipeStatementContext) {}

// ExitPipeStatement is called when production pipeStatement is exited.
func (s *BaseBeepBoopListener) ExitPipeStatement(ctx *PipeStatementContext) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BaseBeepBoopListener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BaseBeepBoopListener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterExprIfStatement is called when production exprIfStatement is entered.
func (s *BaseBeepBoopListener) EnterExprIfStatement(ctx *ExprIfStatementContext) {}

// ExitExprIfStatement is called when production exprIfStatement is exited.
func (s *BaseBeepBoopListener) ExitExprIfStatement(ctx *ExprIfStatementContext) {}

// EnterFncallIfStatement is called when production fncallIfStatement is entered.
func (s *BaseBeepBoopListener) EnterFncallIfStatement(ctx *FncallIfStatementContext) {}

// ExitFncallIfStatement is called when production fncallIfStatement is exited.
func (s *BaseBeepBoopListener) ExitFncallIfStatement(ctx *FncallIfStatementContext) {}

// EnterExprReturn is called when production exprReturn is entered.
func (s *BaseBeepBoopListener) EnterExprReturn(ctx *ExprReturnContext) {}

// ExitExprReturn is called when production exprReturn is exited.
func (s *BaseBeepBoopListener) ExitExprReturn(ctx *ExprReturnContext) {}

// EnterFncallReturn is called when production fncallReturn is entered.
func (s *BaseBeepBoopListener) EnterFncallReturn(ctx *FncallReturnContext) {}

// ExitFncallReturn is called when production fncallReturn is exited.
func (s *BaseBeepBoopListener) ExitFncallReturn(ctx *FncallReturnContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBeepBoopListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBeepBoopListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseBeepBoopListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseBeepBoopListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterTermExpr is called when production termExpr is entered.
func (s *BaseBeepBoopListener) EnterTermExpr(ctx *TermExprContext) {}

// ExitTermExpr is called when production termExpr is exited.
func (s *BaseBeepBoopListener) ExitTermExpr(ctx *TermExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseBeepBoopListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseBeepBoopListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BaseBeepBoopListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BaseBeepBoopListener) ExitPipe(ctx *PipeContext) {}

// EnterFncall is called when production fncall is entered.
func (s *BaseBeepBoopListener) EnterFncall(ctx *FncallContext) {}

// ExitFncall is called when production fncall is exited.
func (s *BaseBeepBoopListener) ExitFncall(ctx *FncallContext) {}

// EnterLabelTerm is called when production labelTerm is entered.
func (s *BaseBeepBoopListener) EnterLabelTerm(ctx *LabelTermContext) {}

// ExitLabelTerm is called when production labelTerm is exited.
func (s *BaseBeepBoopListener) ExitLabelTerm(ctx *LabelTermContext) {}

// EnterStringTerm is called when production stringTerm is entered.
func (s *BaseBeepBoopListener) EnterStringTerm(ctx *StringTermContext) {}

// ExitStringTerm is called when production stringTerm is exited.
func (s *BaseBeepBoopListener) ExitStringTerm(ctx *StringTermContext) {}

// EnterIntTerm is called when production intTerm is entered.
func (s *BaseBeepBoopListener) EnterIntTerm(ctx *IntTermContext) {}

// ExitIntTerm is called when production intTerm is exited.
func (s *BaseBeepBoopListener) ExitIntTerm(ctx *IntTermContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseBeepBoopListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseBeepBoopListener) ExitLabel(ctx *LabelContext) {}
