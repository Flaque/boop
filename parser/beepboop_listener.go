// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BeepBoopListener is a complete listener for a parse tree produced by BeepBoopParser.
type BeepBoopListener interface {
	antlr.ParseTreeListener

	// EnterBeepboop is called when entering the beepboop production.
	EnterBeepboop(c *BeepboopContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExprStatement is called when entering the exprStatement production.
	EnterExprStatement(c *ExprStatementContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterTermExpr is called when entering the termExpr production.
	EnterTermExpr(c *TermExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// ExitBeepboop is called when exiting the beepboop production.
	ExitBeepboop(c *BeepboopContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExprStatement is called when exiting the exprStatement production.
	ExitExprStatement(c *ExprStatementContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitTermExpr is called when exiting the termExpr production.
	ExitTermExpr(c *TermExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
