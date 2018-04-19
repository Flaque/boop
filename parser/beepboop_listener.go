// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BeepBoopListener is a complete listener for a parse tree produced by BeepBoopParser.
type BeepBoopListener interface {
	antlr.ParseTreeListener

	// EnterBeepboop is called when entering the beepboop production.
	EnterBeepboop(c *BeepboopContext)

	// EnterStatementCode is called when entering the statementCode production.
	EnterStatementCode(c *StatementCodeContext)

	// EnterFuncdefCode is called when entering the funcdefCode production.
	EnterFuncdefCode(c *FuncdefCodeContext)

	// EnterFncallStatement is called when entering the fncallStatement production.
	EnterFncallStatement(c *FncallStatementContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterPipeStatement is called when entering the pipeStatement production.
	EnterPipeStatement(c *PipeStatementContext)

	// EnterNoopStatement is called when entering the noopStatement production.
	EnterNoopStatement(c *NoopStatementContext)

	// EnterFuncguts is called when entering the funcguts production.
	EnterFuncguts(c *FuncgutsContext)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterFncall is called when entering the fncall production.
	EnterFncall(c *FncallContext)

	// EnterExprIfStatement is called when entering the exprIfStatement production.
	EnterExprIfStatement(c *ExprIfStatementContext)

	// EnterFncallIfStatement is called when entering the fncallIfStatement production.
	EnterFncallIfStatement(c *FncallIfStatementContext)

	// EnterExprReturn is called when entering the exprReturn production.
	EnterExprReturn(c *ExprReturnContext)

	// EnterFncallReturn is called when entering the fncallReturn production.
	EnterFncallReturn(c *FncallReturnContext)

	// EnterExprAssign is called when entering the exprAssign production.
	EnterExprAssign(c *ExprAssignContext)

	// EnterFncallAssign is called when entering the fncallAssign production.
	EnterFncallAssign(c *FncallAssignContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterTermExpr is called when entering the termExpr production.
	EnterTermExpr(c *TermExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

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

	// ExitStatementCode is called when exiting the statementCode production.
	ExitStatementCode(c *StatementCodeContext)

	// ExitFuncdefCode is called when exiting the funcdefCode production.
	ExitFuncdefCode(c *FuncdefCodeContext)

	// ExitFncallStatement is called when exiting the fncallStatement production.
	ExitFncallStatement(c *FncallStatementContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitPipeStatement is called when exiting the pipeStatement production.
	ExitPipeStatement(c *PipeStatementContext)

	// ExitNoopStatement is called when exiting the noopStatement production.
	ExitNoopStatement(c *NoopStatementContext)

	// ExitFuncguts is called when exiting the funcguts production.
	ExitFuncguts(c *FuncgutsContext)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitFncall is called when exiting the fncall production.
	ExitFncall(c *FncallContext)

	// ExitExprIfStatement is called when exiting the exprIfStatement production.
	ExitExprIfStatement(c *ExprIfStatementContext)

	// ExitFncallIfStatement is called when exiting the fncallIfStatement production.
	ExitFncallIfStatement(c *FncallIfStatementContext)

	// ExitExprReturn is called when exiting the exprReturn production.
	ExitExprReturn(c *ExprReturnContext)

	// ExitFncallReturn is called when exiting the fncallReturn production.
	ExitFncallReturn(c *FncallReturnContext)

	// ExitExprAssign is called when exiting the exprAssign production.
	ExitExprAssign(c *ExprAssignContext)

	// ExitFncallAssign is called when exiting the fncallAssign production.
	ExitFncallAssign(c *FncallAssignContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitTermExpr is called when exiting the termExpr production.
	ExitTermExpr(c *TermExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitLabelTerm is called when exiting the labelTerm production.
	ExitLabelTerm(c *LabelTermContext)

	// ExitStringTerm is called when exiting the stringTerm production.
	ExitStringTerm(c *StringTermContext)

	// ExitIntTerm is called when exiting the intTerm production.
	ExitIntTerm(c *IntTermContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
