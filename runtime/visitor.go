package runtime

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/boop/parser"
)

type BeepBoopVisitor struct {
	*parser.BaseBeepBoopVisitor

	// Tree of Hashmap frames to store variables
	tree *Tree

	// where to print to
	logger *Logger
}

func (v *BeepBoopVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *BeepBoopVisitor) VisitBeepboop(ctx *parser.BeepboopContext) interface{} {

	// Create new tree context
	tree := NewTree(v.tree)
	v.tree = &tree

	val := v.Visit(ctx.Block())

	// Pop off this tree context
	v.tree = v.tree.Parent

	return val
}

func (v *BeepBoopVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {

	for i := 0; i < ctx.GetChildCount(); i++ {
		v.Visit(ctx.Statement(i))
	}

	return nil
}

func (v *BeepBoopVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	return v.Visit(ctx.Assignment())
}

func (v *BeepBoopVisitor) VisitFncallStatement(ctx *parser.FncallStatementContext) interface{} {
	return v.Visit(ctx.Fncall())
}

func (v *BeepBoopVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {

	// Make sure we visit the label, even though we might be using it
	v.Visit(ctx.Label())

	// Grab expression value
	value := v.Visit(ctx.Expr())

	v.tree.Set(ctx.Label().GetText(), value)

	return nil
}

func (v *BeepBoopVisitor) VisitFncall(ctx *parser.FncallContext) interface{} {
	fn, err := v.tree.Get(ctx.STRING().GetText())

	// TODO Check if a function doesn't exist
	if err != nil {

	}

	// Collect args
	args := []string{}
	for _, term := range ctx.AllTerm() {
		s, _ := AnythingToString(v.Visit(term))
		args = append(args, s)
	}

	if fn != nil {
		// TODO: do a real function
		return nil
	}

	name := ctx.STRING().GetText()

	output := RunCmd(name, args...)

	// Echo is the only command allowed to output
	if name == "echo" {
		v.logger.Print(output)
	}

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

	val, _ := v.tree.Get(label)

	// TODO throw label doesn't exist term here

	return val
}

func (v *BeepBoopVisitor) VisitStringTerm(ctx *parser.StringTermContext) interface{} {
	return ctx.STRING().GetText()
}

func (v *BeepBoopVisitor) VisitIntTerm(ctx *parser.IntTermContext) interface{} {
	i, _ := RuleToInt(ctx.INT())
	return i
}
