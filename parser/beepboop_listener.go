// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BeepBoopListener is a complete listener for a parse tree produced by BeepBoopParser.
type BeepBoopListener interface {
	antlr.ParseTreeListener

	// EnterBoop is called when entering the boop production.
	EnterBoop(c *BoopContext)

	// EnterFuncdefCode is called when entering the funcdefCode production.
	EnterFuncdefCode(c *FuncdefCodeContext)

	// EnterStatementCode is called when entering the statementCode production.
	EnterStatementCode(c *StatementCodeContext)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterFuncguts is called when entering the funcguts production.
	EnterFuncguts(c *FuncgutsContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterFncallStatement is called when entering the fncallStatement production.
	EnterFncallStatement(c *FncallStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterPipeStatement is called when entering the pipeStatement production.
	EnterPipeStatement(c *PipeStatementContext)

	// EnterNewlineStatement is called when entering the newlineStatement production.
	EnterNewlineStatement(c *NewlineStatementContext)

	// EnterIfstat is called when entering the ifstat production.
	EnterIfstat(c *IfstatContext)

	// EnterReturnstat is called when entering the returnstat production.
	EnterReturnstat(c *ReturnstatContext)

	// EnterStructexpr is called when entering the structexpr production.
	EnterStructexpr(c *StructexprContext)

	// EnterAssignstat is called when entering the assignstat production.
	EnterAssignstat(c *AssignstatContext)

	// EnterExprAssign is called when entering the exprAssign production.
	EnterExprAssign(c *ExprAssignContext)

	// EnterFncallAssign is called when entering the fncallAssign production.
	EnterFncallAssign(c *FncallAssignContext)

	// EnterParen_fncall is called when entering the paren_fncall production.
	EnterParen_fncall(c *Paren_fncallContext)

	// EnterFncall is called when entering the fncall production.
	EnterFncall(c *FncallContext)

	// EnterLabelTerm is called when entering the labelTerm production.
	EnterLabelTerm(c *LabelTermContext)

	// EnterLiteralTerm is called when entering the literalTerm production.
	EnterLiteralTerm(c *LiteralTermContext)

	// EnterMathTerm is called when entering the mathTerm production.
	EnterMathTerm(c *MathTermContext)

	// EnterParenfncallTerm is called when entering the parenfncallTerm production.
	EnterParenfncallTerm(c *ParenfncallTermContext)

	// EnterStructTerm is called when entering the structTerm production.
	EnterStructTerm(c *StructTermContext)

	// EnterListTerm is called when entering the listTerm production.
	EnterListTerm(c *ListTermContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterListterm is called when entering the listterm production.
	EnterListterm(c *ListtermContext)

	// EnterGthanequalsCond is called when entering the gthanequalsCond production.
	EnterGthanequalsCond(c *GthanequalsCondContext)

	// EnterEqualsCond is called when entering the equalsCond production.
	EnterEqualsCond(c *EqualsCondContext)

	// EnterOrCond is called when entering the orCond production.
	EnterOrCond(c *OrCondContext)

	// EnterGthanCond is called when entering the gthanCond production.
	EnterGthanCond(c *GthanCondContext)

	// EnterAndCond is called when entering the andCond production.
	EnterAndCond(c *AndCondContext)

	// EnterLthanCond is called when entering the lthanCond production.
	EnterLthanCond(c *LthanCondContext)

	// EnterLthanequalsCond is called when entering the lthanequalsCond production.
	EnterLthanequalsCond(c *LthanequalsCondContext)

	// EnterBoolCond is called when entering the boolCond production.
	EnterBoolCond(c *BoolCondContext)

	// EnterSoloMath is called when entering the soloMath production.
	EnterSoloMath(c *SoloMathContext)

	// EnterAdditiveMath is called when entering the additiveMath production.
	EnterAdditiveMath(c *AdditiveMathContext)

	// EnterUnarySubMath is called when entering the unarySubMath production.
	EnterUnarySubMath(c *UnarySubMathContext)

	// EnterNumLiteral is called when entering the numLiteral production.
	EnterNumLiteral(c *NumLiteralContext)

	// EnterLettersLiteral is called when entering the lettersLiteral production.
	EnterLettersLiteral(c *LettersLiteralContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterQuotedLiteral is called when entering the quotedLiteral production.
	EnterQuotedLiteral(c *QuotedLiteralContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterBoolexpr is called when entering the boolexpr production.
	EnterBoolexpr(c *BoolexprContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// ExitBoop is called when exiting the boop production.
	ExitBoop(c *BoopContext)

	// ExitFuncdefCode is called when exiting the funcdefCode production.
	ExitFuncdefCode(c *FuncdefCodeContext)

	// ExitStatementCode is called when exiting the statementCode production.
	ExitStatementCode(c *StatementCodeContext)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitFuncguts is called when exiting the funcguts production.
	ExitFuncguts(c *FuncgutsContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitFncallStatement is called when exiting the fncallStatement production.
	ExitFncallStatement(c *FncallStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitPipeStatement is called when exiting the pipeStatement production.
	ExitPipeStatement(c *PipeStatementContext)

	// ExitNewlineStatement is called when exiting the newlineStatement production.
	ExitNewlineStatement(c *NewlineStatementContext)

	// ExitIfstat is called when exiting the ifstat production.
	ExitIfstat(c *IfstatContext)

	// ExitReturnstat is called when exiting the returnstat production.
	ExitReturnstat(c *ReturnstatContext)

	// ExitStructexpr is called when exiting the structexpr production.
	ExitStructexpr(c *StructexprContext)

	// ExitAssignstat is called when exiting the assignstat production.
	ExitAssignstat(c *AssignstatContext)

	// ExitExprAssign is called when exiting the exprAssign production.
	ExitExprAssign(c *ExprAssignContext)

	// ExitFncallAssign is called when exiting the fncallAssign production.
	ExitFncallAssign(c *FncallAssignContext)

	// ExitParen_fncall is called when exiting the paren_fncall production.
	ExitParen_fncall(c *Paren_fncallContext)

	// ExitFncall is called when exiting the fncall production.
	ExitFncall(c *FncallContext)

	// ExitLabelTerm is called when exiting the labelTerm production.
	ExitLabelTerm(c *LabelTermContext)

	// ExitLiteralTerm is called when exiting the literalTerm production.
	ExitLiteralTerm(c *LiteralTermContext)

	// ExitMathTerm is called when exiting the mathTerm production.
	ExitMathTerm(c *MathTermContext)

	// ExitParenfncallTerm is called when exiting the parenfncallTerm production.
	ExitParenfncallTerm(c *ParenfncallTermContext)

	// ExitStructTerm is called when exiting the structTerm production.
	ExitStructTerm(c *StructTermContext)

	// ExitListTerm is called when exiting the listTerm production.
	ExitListTerm(c *ListTermContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitListterm is called when exiting the listterm production.
	ExitListterm(c *ListtermContext)

	// ExitGthanequalsCond is called when exiting the gthanequalsCond production.
	ExitGthanequalsCond(c *GthanequalsCondContext)

	// ExitEqualsCond is called when exiting the equalsCond production.
	ExitEqualsCond(c *EqualsCondContext)

	// ExitOrCond is called when exiting the orCond production.
	ExitOrCond(c *OrCondContext)

	// ExitGthanCond is called when exiting the gthanCond production.
	ExitGthanCond(c *GthanCondContext)

	// ExitAndCond is called when exiting the andCond production.
	ExitAndCond(c *AndCondContext)

	// ExitLthanCond is called when exiting the lthanCond production.
	ExitLthanCond(c *LthanCondContext)

	// ExitLthanequalsCond is called when exiting the lthanequalsCond production.
	ExitLthanequalsCond(c *LthanequalsCondContext)

	// ExitBoolCond is called when exiting the boolCond production.
	ExitBoolCond(c *BoolCondContext)

	// ExitSoloMath is called when exiting the soloMath production.
	ExitSoloMath(c *SoloMathContext)

	// ExitAdditiveMath is called when exiting the additiveMath production.
	ExitAdditiveMath(c *AdditiveMathContext)

	// ExitUnarySubMath is called when exiting the unarySubMath production.
	ExitUnarySubMath(c *UnarySubMathContext)

	// ExitNumLiteral is called when exiting the numLiteral production.
	ExitNumLiteral(c *NumLiteralContext)

	// ExitLettersLiteral is called when exiting the lettersLiteral production.
	ExitLettersLiteral(c *LettersLiteralContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitQuotedLiteral is called when exiting the quotedLiteral production.
	ExitQuotedLiteral(c *QuotedLiteralContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitBoolexpr is called when exiting the boolexpr production.
	ExitBoolexpr(c *BoolexprContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
