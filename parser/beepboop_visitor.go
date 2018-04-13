// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BeepBoopParser.
type BeepBoopVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BeepBoopParser#beepboop.
	VisitBeepboop(ctx *BeepboopContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#math.
	VisitMath(ctx *MathContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#term.
	VisitTerm(ctx *TermContext) interface{}
}
