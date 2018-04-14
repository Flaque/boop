// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BeepBoopParser.
type BeepBoopVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BeepBoopParser#beepboop.
	VisitBeepboop(ctx *BeepboopContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#exprStatement.
	VisitExprStatement(ctx *ExprStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#unaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#termExpr.
	VisitTermExpr(ctx *TermExprContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#label.
	VisitLabel(ctx *LabelContext) interface{}
}
