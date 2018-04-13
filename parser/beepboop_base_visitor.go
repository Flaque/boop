// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBeepBoopVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBeepBoopVisitor) VisitBeepboop(ctx *BeepboopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitTermExpr(ctx *TermExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitMinusExpr(ctx *MinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeepBoopVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}
