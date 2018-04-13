// Code generated from parser/Do.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Do

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DoParser.
type DoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DoParser#do.
	VisitDo(ctx *DoContext) interface{}

	// Visit a parse tree produced by DoParser#math.
	VisitMath(ctx *MathContext) interface{}

	// Visit a parse tree produced by DoParser#term.
	VisitTerm(ctx *TermContext) interface{}
}
