package runtime

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/beep/parser"
)

type BeepBoopVisitor struct {
	*parser.BaseBeepBoopVisitor

	// Stack of Hashmap frames to store variables
	stack *Stack
}

func (v *BeepBoopVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *BeepBoopVisitor) VisitBeepboop(ctx *parser.BeepboopContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *BeepBoopVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {

	v.stack.Push(NewFrame())
	val := v.Visit(ctx.Statement(0))
	v.stack.Pop()

	return val
}

func (v *BeepBoopVisitor) VisitExprStatement(ctx *parser.ExprStatementContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *BeepBoopVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	return v.Visit(ctx.Assignment())
}

func (v *BeepBoopVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	label := ctx.Label().GetText()
	value := ctx.Expr().GetText()

	v.stack.Add(label, value)

	return v.VisitChildren(ctx)
}

func (v *BeepBoopVisitor) VisitTermExpr(ctx *parser.TermExprContext) interface{} {
	return v.Visit(ctx.Term())
}

func (v *BeepBoopVisitor) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	a, oka := v.Visit(ctx.Expr(0)).(int)
	b, okb := v.Visit(ctx.Expr(1)).(int)

	isAdd := ctx.MINUS() == nil

	total := 0

	if oka {
		total += a
	}

	if okb {
		if isAdd {
			total += b
		} else {
			total -= b
		}
	}

	return total
}

func (v *BeepBoopVisitor) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) interface{} {
	a, ok := v.Visit(ctx.Expr()).(int)

	total := 0

	if ok {
		total -= a
	}

	return total
}

func (v *BeepBoopVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	i, _ := RuleToInt(ctx.INT())
	return i
}
