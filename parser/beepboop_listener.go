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

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterFncallStatement is called when entering the fncallStatement production.
	EnterFncallStatement(c *FncallStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterPipeStatement is called when entering the pipeStatement production.
	EnterPipeStatement(c *PipeStatementContext)

	// EnterExprIfStatement is called when entering the exprIfStatement production.
	EnterExprIfStatement(c *ExprIfStatementContext)

	// EnterFncallIfStatement is called when entering the fncallIfStatement production.
	EnterFncallIfStatement(c *FncallIfStatementContext)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterExprReturn is called when entering the exprReturn production.
	EnterExprReturn(c *ExprReturnContext)

	// EnterFncallReturn is called when entering the fncallReturn production.
	EnterFncallReturn(c *FncallReturnContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterTermExpr is called when entering the termExpr production.
	EnterTermExpr(c *TermExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

	// EnterFncall is called when entering the fncall production.
	EnterFncall(c *FncallContext)

	// EnterLabelTerm is called when entering the labelTerm production.
	EnterLabelTerm(c *LabelTermContext)

	// EnterStringTerm is called when entering the stringTerm production.
	EnterStringTerm(c *StringTermContext)

	// EnterIntTerm is called when entering the intTerm production.
	EnterIntTerm(c *IntTermContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// ExitBeepboop is called when exiting the beepboop production.
	ExitBeepboop(c *BeepboopContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitFncallStatement is called when exiting the fncallStatement production.
	ExitFncallStatement(c *FncallStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitPipeStatement is called when exiting the pipeStatement production.
	ExitPipeStatement(c *PipeStatementContext)

	// ExitExprIfStatement is called when exiting the exprIfStatement production.
	ExitExprIfStatement(c *ExprIfStatementContext)

	// ExitFncallIfStatement is called when exiting the fncallIfStatement production.
	ExitFncallIfStatement(c *FncallIfStatementContext)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitExprReturn is called when exiting the exprReturn production.
	ExitExprReturn(c *ExprReturnContext)

	// ExitFncallReturn is called when exiting the fncallReturn production.
	ExitFncallReturn(c *FncallReturnContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitTermExpr is called when exiting the termExpr production.
	ExitTermExpr(c *TermExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitFncall is called when exiting the fncall production.
	ExitFncall(c *FncallContext)

	// ExitLabelTerm is called when exiting the labelTerm production.
	ExitLabelTerm(c *LabelTermContext)

	// ExitStringTerm is called when exiting the stringTerm production.
	ExitStringTerm(c *StringTermContext)

	// ExitIntTerm is called when exiting the intTerm production.
	ExitIntTerm(c *IntTermContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
