// Code generated from parser/Do.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Do

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDoListener is a complete listener for a parse tree produced by DoParser.
type BaseDoListener struct{}

var _ DoListener = &BaseDoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDo is called when production do is entered.
func (s *BaseDoListener) EnterDo(ctx *DoContext) {}

// ExitDo is called when production do is exited.
func (s *BaseDoListener) ExitDo(ctx *DoContext) {}

// EnterMath is called when production math is entered.
func (s *BaseDoListener) EnterMath(ctx *MathContext) {}

// ExitMath is called when production math is exited.
func (s *BaseDoListener) ExitMath(ctx *MathContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseDoListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseDoListener) ExitTerm(ctx *TermContext) {}
