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

	// EnterExportStatement is called when entering the exportStatement production.
	EnterExportStatement(c *ExportStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterFncallStatement is called when entering the fncallStatement production.
	EnterFncallStatement(c *FncallStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterPipeStatement is called when entering the pipeStatement production.
	EnterPipeStatement(c *PipeStatementContext)

	// EnterNewlineStatement is called when entering the newlineStatement production.
	EnterNewlineStatement(c *NewlineStatementContext)

	// EnterImportstat is called when entering the importstat production.
	EnterImportstat(c *ImportstatContext)

	// EnterExportstat is called when entering the exportstat production.
	EnterExportstat(c *ExportstatContext)

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

	// EnterFlagFnargs is called when entering the flagFnargs production.
	EnterFlagFnargs(c *FlagFnargsContext)

	// EnterMultFnargs is called when entering the multFnargs production.
	EnterMultFnargs(c *MultFnargsContext)

	// EnterDivFnargs is called when entering the divFnargs production.
	EnterDivFnargs(c *DivFnargsContext)

	// EnterTermFnargs is called when entering the termFnargs production.
	EnterTermFnargs(c *TermFnargsContext)

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

	// EnterMultiplicativeMath is called when entering the multiplicativeMath production.
	EnterMultiplicativeMath(c *MultiplicativeMathContext)

	// EnterAdditiveMath is called when entering the additiveMath production.
	EnterAdditiveMath(c *AdditiveMathContext)

	// EnterUnarySubMath is called when entering the unarySubMath production.
	EnterUnarySubMath(c *UnarySubMathContext)

	// EnterNumLiteral is called when entering the numLiteral production.
	EnterNumLiteral(c *NumLiteralContext)

	// EnterWordsLiteral is called when entering the wordsLiteral production.
	EnterWordsLiteral(c *WordsLiteralContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterIntNum is called when entering the intNum production.
	EnterIntNum(c *IntNumContext)

	// EnterFloatNum is called when entering the floatNum production.
	EnterFloatNum(c *FloatNumContext)

	// EnterLetterWords is called when entering the letterWords production.
	EnterLetterWords(c *LetterWordsContext)

	// EnterStringWords is called when entering the stringWords production.
	EnterStringWords(c *StringWordsContext)

	// EnterQuotedWords is called when entering the quotedWords production.
	EnterQuotedWords(c *QuotedWordsContext)

	// EnterBoolexpr is called when entering the boolexpr production.
	EnterBoolexpr(c *BoolexprContext)

	// EnterPipeToFncallPipe is called when entering the pipeToFncallPipe production.
	EnterPipeToFncallPipe(c *PipeToFncallPipeContext)

	// EnterPipeToPipe is called when entering the pipeToPipe production.
	EnterPipeToPipe(c *PipeToPipeContext)

	// EnterTermPipe is called when entering the termPipe production.
	EnterTermPipe(c *TermPipeContext)

	// EnterNewlinePipe is called when entering the newlinePipe production.
	EnterNewlinePipe(c *NewlinePipeContext)

	// EnterFncallPipe is called when entering the fncallPipe production.
	EnterFncallPipe(c *FncallPipeContext)

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

	// ExitExportStatement is called when exiting the exportStatement production.
	ExitExportStatement(c *ExportStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitFncallStatement is called when exiting the fncallStatement production.
	ExitFncallStatement(c *FncallStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitPipeStatement is called when exiting the pipeStatement production.
	ExitPipeStatement(c *PipeStatementContext)

	// ExitNewlineStatement is called when exiting the newlineStatement production.
	ExitNewlineStatement(c *NewlineStatementContext)

	// ExitImportstat is called when exiting the importstat production.
	ExitImportstat(c *ImportstatContext)

	// ExitExportstat is called when exiting the exportstat production.
	ExitExportstat(c *ExportstatContext)

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

	// ExitFlagFnargs is called when exiting the flagFnargs production.
	ExitFlagFnargs(c *FlagFnargsContext)

	// ExitMultFnargs is called when exiting the multFnargs production.
	ExitMultFnargs(c *MultFnargsContext)

	// ExitDivFnargs is called when exiting the divFnargs production.
	ExitDivFnargs(c *DivFnargsContext)

	// ExitTermFnargs is called when exiting the termFnargs production.
	ExitTermFnargs(c *TermFnargsContext)

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

	// ExitMultiplicativeMath is called when exiting the multiplicativeMath production.
	ExitMultiplicativeMath(c *MultiplicativeMathContext)

	// ExitAdditiveMath is called when exiting the additiveMath production.
	ExitAdditiveMath(c *AdditiveMathContext)

	// ExitUnarySubMath is called when exiting the unarySubMath production.
	ExitUnarySubMath(c *UnarySubMathContext)

	// ExitNumLiteral is called when exiting the numLiteral production.
	ExitNumLiteral(c *NumLiteralContext)

	// ExitWordsLiteral is called when exiting the wordsLiteral production.
	ExitWordsLiteral(c *WordsLiteralContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitIntNum is called when exiting the intNum production.
	ExitIntNum(c *IntNumContext)

	// ExitFloatNum is called when exiting the floatNum production.
	ExitFloatNum(c *FloatNumContext)

	// ExitLetterWords is called when exiting the letterWords production.
	ExitLetterWords(c *LetterWordsContext)

	// ExitStringWords is called when exiting the stringWords production.
	ExitStringWords(c *StringWordsContext)

	// ExitQuotedWords is called when exiting the quotedWords production.
	ExitQuotedWords(c *QuotedWordsContext)

	// ExitBoolexpr is called when exiting the boolexpr production.
	ExitBoolexpr(c *BoolexprContext)

	// ExitPipeToFncallPipe is called when exiting the pipeToFncallPipe production.
	ExitPipeToFncallPipe(c *PipeToFncallPipeContext)

	// ExitPipeToPipe is called when exiting the pipeToPipe production.
	ExitPipeToPipe(c *PipeToPipeContext)

	// ExitTermPipe is called when exiting the termPipe production.
	ExitTermPipe(c *TermPipeContext)

	// ExitNewlinePipe is called when exiting the newlinePipe production.
	ExitNewlinePipe(c *NewlinePipeContext)

	// ExitFncallPipe is called when exiting the fncallPipe production.
	ExitFncallPipe(c *FncallPipeContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
