// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBeepBoopVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBeepBoopVisitor) VisitBeepboop(ctx *BeepboopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitStatementCode(ctx *StatementCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFuncdefCode(ctx *FuncdefCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncallStatement(ctx *FncallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitPipeStatement(ctx *PipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFuncdef(ctx *FuncdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitExprIfStatement(ctx *ExprIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncallIfStatement(ctx *FncallIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitExprReturn(ctx *ExprReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncallReturn(ctx *FncallReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitTermExpr(ctx *TermExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitPipe(ctx *PipeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitFncall(ctx *FncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLabelTerm(ctx *LabelTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitStringTerm(ctx *StringTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitIntTerm(ctx *IntTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}
