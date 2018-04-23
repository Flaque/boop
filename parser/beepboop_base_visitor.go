// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBeepBoopVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBeepBoopVisitor) VisitBoop(ctx *BoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFuncdefCode(ctx *FuncdefCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitStatementCode(ctx *StatementCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFuncdef(ctx *FuncdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFuncguts(ctx *FuncgutsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncallStatement(ctx *FncallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitPipeStatement(ctx *PipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitNewlineStatement(ctx *NewlineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitImportstat(ctx *ImportstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitExportstat(ctx *ExportstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitIfstat(ctx *IfstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitReturnstat(ctx *ReturnstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitStructexpr(ctx *StructexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAssignstat(ctx *AssignstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitExprAssign(ctx *ExprAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncallAssign(ctx *FncallAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitParen_fncall(ctx *Paren_fncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncall(ctx *FncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFnargs(ctx *FnargsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLabelTerm(ctx *LabelTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLiteralTerm(ctx *LiteralTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitMathTerm(ctx *MathTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitParenfncallTerm(ctx *ParenfncallTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitStructTerm(ctx *StructTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitListTerm(ctx *ListTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitListterm(ctx *ListtermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitGthanequalsCond(ctx *GthanequalsCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitEqualsCond(ctx *EqualsCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitOrCond(ctx *OrCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitGthanCond(ctx *GthanCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAndCond(ctx *AndCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLthanCond(ctx *LthanCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLthanequalsCond(ctx *LthanequalsCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitBoolCond(ctx *BoolCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitSoloMath(ctx *SoloMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitMultiplicativeMath(ctx *MultiplicativeMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAdditiveMath(ctx *AdditiveMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitUnarySubMath(ctx *UnarySubMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitNumLiteral(ctx *NumLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitWordsLiteral(ctx *WordsLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitWords(ctx *WordsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitBoolexpr(ctx *BoolexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitPipe(ctx *PipeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}
