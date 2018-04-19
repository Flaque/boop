package runtime

import (
	"strconv"

	"github.com/Flaque/boop/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type BeepBoopVisitor struct {
	*parser.BaseBeepBoopVisitor

	// Tree of Hashmap frames to store variables
	tree *Tree

	// where to print to
	logger *Logger

	visits []string
}

func (v *BeepBoopVisitor) StartScope() {
	tree := NewTree(v.tree)
	v.tree = &tree
}

func (v *BeepBoopVisitor) EndScope() {
	v.tree = v.tree.Parent
}

func (v *BeepBoopVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *BeepBoopVisitor) VisitBeepboop(ctx *parser.BeepboopContext) interface{} {
	v.visits = append(v.visits, "beepboop")

	// Create new tree context
	v.StartScope()

	for i := 0; i < ctx.GetChildCount(); i++ {
		if c := ctx.Code(i); c != nil {
			v.Visit(ctx.Code(i))
		}
	}

	// Pop off this tree context
	v.EndScope()
	return nil
}

func (v *BeepBoopVisitor) VisitStatementCode(ctx *parser.StatementCodeContext) interface{} {
	v.visits = append(v.visits, "statementcode")
	return v.Visit(ctx.Statement())
}

func (v *BeepBoopVisitor) VisitFuncdefCode(ctx *parser.FuncdefCodeContext) interface{} {
	v.visits = append(v.visits, "funcdefcode")
	return v.Visit(ctx.Funcdef())
}

func (v *BeepBoopVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	v.visits = append(v.visits, "assignstatement")
	return v.Visit(ctx.Assignment())
}

func (v *BeepBoopVisitor) VisitFncallStatement(ctx *parser.FncallStatementContext) interface{} {
	v.visits = append(v.visits, "fncallstatement")
	return v.Visit(ctx.Fncall())
}

func (v *BeepBoopVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	v.visits = append(v.visits, "assignment")

	// Grab expression value
	value := v.Visit(ctx.Expr())
	label := ctx.Label().GetText()

	v.tree.Set(label, value)

	return nil
}

func (v *BeepBoopVisitor) VisitFuncguts(ctx *parser.FuncgutsContext) interface{} {
	for i := 0; i < ctx.GetChildCount(); i++ {
		v.Visit(ctx.Statement(i))
	}
	return nil
}

func (v *BeepBoopVisitor) VisitFuncdef(ctx *parser.FuncdefContext) interface{} {
	name := ctx.STRING().GetText()

	args := []string{}
	for _, label := range ctx.AllLabel() {
		s, _ := AnythingToString(label.GetText())
		args = append(args, s)
	}

	guts := ctx.Funcguts()

	v.tree.Set(name, NewFunc(name, args, guts))

	return nil
}

func (v *BeepBoopVisitor) VisitFncall(ctx *parser.FncallContext) interface{} {
	v.visits = append(v.visits, "fncall")

	// Collect args
	args := []string{}
	for _, term := range ctx.AllTerm() {
		s, _ := AnythingToString(v.Visit(term))
		args = append(args, s)
	}

	// Collect function name
	fn, err := v.tree.Get(ctx.STRING().GetText())

	// Run a real function
	if err == nil {
		function, _ := fn.(Function)

		if len(args) != len(function.args) {
			ThrowIncorrectFunctionCall(v.logger, function.name, len(function.args), len(args))
			return nil
		}

		// Inject arguments into a new scope
		v.StartScope()
		for i, label := range function.args {
			v.tree.Set(label, args[i])
		}

		if function.guts != nil {
			v.Visit(function.guts)
		}
		v.EndScope()

		return nil
	}

	// Run a command line function
	name := ctx.STRING().GetText()

	output := RunCmd(name, args...)

	// Echo is the only command allowed to output
	if name == "echo" {
		v.logger.Print(output)
	}

	return output
}

func (v *BeepBoopVisitor) VisitTermExpr(ctx *parser.TermExprContext) interface{} {
	v.visits = append(v.visits, "termexpr")
	return v.Visit(ctx.Term())
}

func (v *BeepBoopVisitor) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	v.visits = append(v.visits, "additiveexpr")

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
	v.visits = append(v.visits, "unaryminusexpr")

	a, ok := v.Visit(ctx.Expr()).(int)

	total := 0

	if ok {
		total -= a
	}

	return total
}

func (v *BeepBoopVisitor) VisitLabelTerm(ctx *parser.LabelTermContext) interface{} {
	v.visits = append(v.visits, "labelterm")
	label := ctx.Label().GetText()

	val, err := v.tree.Get(label)

	if err != nil {
		ThrowRuntimeError(v.logger, err.Error())
		return nil
	}

	// Attempt to get it as an int
	if vint, ok := val.(int); ok {
		return vint
	}

	// Attempt to get it as a string
	str := val.(string)

	// Attempt to parse as an int if we can
	if v, err := strconv.Atoi(str); err == nil {
		return v
	}

	return str
}

func (v *BeepBoopVisitor) VisitStringTerm(ctx *parser.StringTermContext) interface{} {
	v.visits = append(v.visits, "stringterm")
	return ctx.STRING().GetText()
}

func (v *BeepBoopVisitor) VisitIntTerm(ctx *parser.IntTermContext) interface{} {
	v.visits = append(v.visits, "intterm")
	i, _ := RuleToInt(ctx.INT())
	return i
}
