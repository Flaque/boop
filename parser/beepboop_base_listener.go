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

// EnterStatement is called when production statement is entered.
func (s *BaseBeepBoopListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBeepBoopListener) ExitStatement(ctx *StatementContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseBeepBoopListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseBeepBoopListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterTermExpr is called when production termExpr is entered.
func (s *BaseBeepBoopListener) EnterTermExpr(ctx *TermExprContext) {}

// ExitTermExpr is called when production termExpr is exited.
func (s *BaseBeepBoopListener) ExitTermExpr(ctx *TermExprContext) {}

// EnterMinusExpr is called when production minusExpr is entered.
func (s *BaseBeepBoopListener) EnterMinusExpr(ctx *MinusExprContext) {}

// ExitMinusExpr is called when production minusExpr is exited.
func (s *BaseBeepBoopListener) ExitMinusExpr(ctx *MinusExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBeepBoopListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBeepBoopListener) ExitTerm(ctx *TermContext) {}
