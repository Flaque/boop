// Code generated from parser/Do.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Do

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDoVisitor) VisitDo(ctx *DoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDoVisitor) VisitMath(ctx *MathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDoVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}
