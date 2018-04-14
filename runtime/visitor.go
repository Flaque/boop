package runtime

import (
	"bytes"
	"fmt"
	"os/exec"

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

	for i := 0; i < ctx.GetChildCount(); i++ {
		v.Visit(ctx.Statement(0))
	}

	v.stack.Pop()

	return nil
}

func (v *BeepBoopVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	return v.Visit(ctx.Assignment())
}

func (v *BeepBoopVisitor) VisitFncallStatement(ctx *parser.FncallStatementContext) interface{} {
	return v.Visit(ctx.Fncall())
}

func (v *BeepBoopVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	label := ctx.Label().GetText()
	value := ctx.Expr().GetText()

	v.stack.Set(label, value)

	return v.VisitChildren(ctx)
}

func (v *BeepBoopVisitor) VisitFncall(ctx *parser.FncallContext) interface{} {
	fn := v.stack.Get(ctx.STRING().GetText())

	// Collect args
	args := []string{}
	for _, expr := range ctx.AllExpr() {
		args = append(args, expr.GetText())
	}

	if fn != nil {
		// TODO: do a real function
		return nil
	}

	name := ctx.STRING().GetText()

	// Attempt to call if from the command line
	cmd := exec.Command(name, args...)

	//TODO Set stdin

	// Set stdout
	var out bytes.Buffer
	cmd.Stdout = &out

	// Launch command
	cmd.Run()

	output := out.String()
	fmt.Print(output)

	return output
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

func (v *BeepBoopVisitor) VisitLabelTerm(ctx *parser.LabelTermContext) interface{} {
	label := ctx.Label().GetText()
	// TODO, throw a runtime error here if the label isn't in the stack
	return v.stack.Get(label)
}

func (v *BeepBoopVisitor) VisitStringTerm(ctx *parser.StringTermContext) interface{} {
	return ctx.STRING().GetText()
}

func (v *BeepBoopVisitor) VisitIntTerm(ctx *parser.IntTermContext) interface{} {
	i, _ := RuleToInt(ctx.INT())
	return i
}
