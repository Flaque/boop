// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BeepBoopParser.
type BeepBoopVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BeepBoopParser#boop.
	VisitBoop(ctx *BoopContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#funcdefCode.
	VisitFuncdefCode(ctx *FuncdefCodeContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#statementCode.
	VisitStatementCode(ctx *StatementCodeContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#funcdef.
	VisitFuncdef(ctx *FuncdefContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#funcguts.
	VisitFuncguts(ctx *FuncgutsContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#fncallStatement.
	VisitFncallStatement(ctx *FncallStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#pipeStatement.
	VisitPipeStatement(ctx *PipeStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#newlineStatement.
	VisitNewlineStatement(ctx *NewlineStatementContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#importstat.
	VisitImportstat(ctx *ImportstatContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#exportstat.
	VisitExportstat(ctx *ExportstatContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#ifstat.
	VisitIfstat(ctx *IfstatContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#returnstat.
	VisitReturnstat(ctx *ReturnstatContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#structexpr.
	VisitStructexpr(ctx *StructexprContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#assignstat.
	VisitAssignstat(ctx *AssignstatContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#exprAssign.
	VisitExprAssign(ctx *ExprAssignContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#fncallAssign.
	VisitFncallAssign(ctx *FncallAssignContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#paren_fncall.
	VisitParen_fncall(ctx *Paren_fncallContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#fncall.
	VisitFncall(ctx *FncallContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#fnargs.
	VisitFnargs(ctx *FnargsContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#labelTerm.
	VisitLabelTerm(ctx *LabelTermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#literalTerm.
	VisitLiteralTerm(ctx *LiteralTermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#mathTerm.
	VisitMathTerm(ctx *MathTermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#parenfncallTerm.
	VisitParenfncallTerm(ctx *ParenfncallTermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#structTerm.
	VisitStructTerm(ctx *StructTermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#listTerm.
	VisitListTerm(ctx *ListTermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#listterm.
	VisitListterm(ctx *ListtermContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#gthanequalsCond.
	VisitGthanequalsCond(ctx *GthanequalsCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#equalsCond.
	VisitEqualsCond(ctx *EqualsCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#orCond.
	VisitOrCond(ctx *OrCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#gthanCond.
	VisitGthanCond(ctx *GthanCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#andCond.
	VisitAndCond(ctx *AndCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#lthanCond.
	VisitLthanCond(ctx *LthanCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#lthanequalsCond.
	VisitLthanequalsCond(ctx *LthanequalsCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#boolCond.
	VisitBoolCond(ctx *BoolCondContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#soloMath.
	VisitSoloMath(ctx *SoloMathContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#multiplicativeMath.
	VisitMultiplicativeMath(ctx *MultiplicativeMathContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#additiveMath.
	VisitAdditiveMath(ctx *AdditiveMathContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#unarySubMath.
	VisitUnarySubMath(ctx *UnarySubMathContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#numLiteral.
	VisitNumLiteral(ctx *NumLiteralContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#wordsLiteral.
	VisitWordsLiteral(ctx *WordsLiteralContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#num.
	VisitNum(ctx *NumContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#words.
	VisitWords(ctx *WordsContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#boolexpr.
	VisitBoolexpr(ctx *BoolexprContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#pipe.
	VisitPipe(ctx *PipeContext) interface{}

	// Visit a parse tree produced by BeepBoopParser#label.
	VisitLabel(ctx *LabelContext) interface{}
}
