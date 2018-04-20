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

func (v *BeepBoopVisitor) VisitBeepboop(ctx *parser.BoopContext) interface{} {

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
	return v.Visit(ctx.Statement())
}

func (v *BeepBoopVisitor) VisitFuncdefCode(ctx *parser.FuncdefCodeContext) interface{} {
	return v.Visit(ctx.Funcdef())
}

func (v *BeepBoopVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	return v.Visit(ctx.Assignstat())
}

func (v *BeepBoopVisitor) VisitFncallStatement(ctx *parser.FncallStatementContext) interface{} {
	return v.Visit(ctx.Fncall())
}

func (v *BeepBoopVisitor) VisitExprAssign(ctx *parser.ExprAssignContext) interface{} {

	// Grab expression value
	value := v.Visit(ctx.Term())
	label := ctx.Label().GetText()

	v.tree.Set(label, value)

	return nil
}

func (v *BeepBoopVisitor) VisitFncallAssign(ctx *parser.FncallAssignContext) interface{} {

	// Grab expression value
	value := v.Visit(ctx.Fncall())
	label := ctx.Label().GetText()

	v.tree.Set(label, value)

	return nil
}

func (v *BeepBoopVisitor) VisitFuncguts(ctx *parser.FuncgutsContext) interface{} {
	var val interface{}
	for i := 0; i < ctx.GetChildCount(); i++ {
		val = v.Visit(ctx.Statement(i))
	}
	return val
}

func (v *BeepBoopVisitor) VisitFuncdef(ctx *parser.FuncdefContext) interface{} {
	name := ctx.Label(0).GetText()

	args := []string{}
	for i, label := range ctx.AllLabel() {

		// Skip name
		if i == 0 {
			continue
		}

		// Rest args
		s, _ := AnythingToString(label.GetText())
		args = append(args, s)
	}

	guts := ctx.Funcguts()

	v.tree.Set(name, NewFunc(name, args, guts))

	return nil
}

func (v *BeepBoopVisitor) VisitFncall(ctx *parser.FncallContext) interface{} {

	// Collect args
	args := []string{}
	for _, term := range ctx.AllTerm() {
		s, _ := AnythingToString(v.Visit(term))
		args = append(args, s)
	}

	// Collect function name
	name := ctx.Label().GetText()
	fn, err := v.tree.Get(name)

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
	output := RunCmd(name, args...)

	// Echo is the only command allowed to output
	if name == "echo" {
		v.logger.Print(output)
	}

	return output
}

func (v *BeepBoopVisitor) VisitLabelTerm(ctx *parser.LabelTermContext) interface{} {
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
