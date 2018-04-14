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

// EnterBlock is called when production block is entered.
func (s *BaseBeepBoopListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBeepBoopListener) ExitBlock(ctx *BlockContext) {}

// EnterExprStatement is called when production exprStatement is entered.
func (s *BaseBeepBoopListener) EnterExprStatement(ctx *ExprStatementContext) {}

// ExitExprStatement is called when production exprStatement is exited.
func (s *BaseBeepBoopListener) ExitExprStatement(ctx *ExprStatementContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseBeepBoopListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseBeepBoopListener) ExitAssignStatement(ctx *AssignStatementContext) {}

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

// EnterTerm is called when production term is entered.
func (s *BaseBeepBoopListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBeepBoopListener) ExitTerm(ctx *TermContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseBeepBoopListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseBeepBoopListener) ExitLabel(ctx *LabelContext) {}
