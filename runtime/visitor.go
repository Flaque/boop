package runtime

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Flaque/boop/parser"
	"github.com/Flaque/boop/util"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type BeepBoopVisitor struct {
	*parser.BaseBeepBoopVisitor

	// Tree of Hashmap frames to store variables
	tree *Tree

	// where to print to
	logger *Logger

	// Errors
	errs []error
}

func NewBeepBoopVisitor(out io.Writer) *BeepBoopVisitor {
	tree := NewTree(nil)
	logger := NewLogger(out)
	return &BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}, tree, logger, []error{}}
}

func (v *BeepBoopVisitor) StartScope() {
	tree := NewTree(v.tree)
	v.tree = tree
}

func (v *BeepBoopVisitor) EndScope() {
	v.tree = v.tree.Parent
}

/*
Visitor Implementation
*/

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

func (v *BeepBoopVisitor) VisitNumLiteral(ctx *parser.NumLiteralContext) interface{} {
	return v.Visit(ctx.Num())
}

func (v *BeepBoopVisitor) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	val := ctx.Boolexpr().GetText()
	return GetVariableFromString(val)
}

func (v *BeepBoopVisitor) VisitNum(ctx *parser.NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BeepBoopVisitor) VisitBoolexpr(ctx *parser.BoolexprContext) interface{} {
	// TODO: Throw error
	return GetVariableFromString(ctx.GetText())
}

func (v *BeepBoopVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitBoop(ctx *parser.BoopContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFuncdefCode(ctx *parser.FuncdefCodeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitStatementCode(ctx *parser.StatementCodeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFuncdef(ctx *parser.FuncdefContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFuncguts(ctx *parser.FuncgutsContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitExportStatement(ctx *parser.ExportStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFncallStatement(ctx *parser.FncallStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitPipeStatement(ctx *parser.PipeStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitNewlineStatement(ctx *parser.NewlineStatementContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitImportstat(ctx *parser.ImportstatContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitExportstat(ctx *parser.ExportstatContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitIfstat(ctx *parser.IfstatContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitReturnstat(ctx *parser.ReturnstatContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitStructexpr(ctx *parser.StructexprContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitAssignstat(ctx *parser.AssignstatContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitExprAssign(ctx *parser.ExprAssignContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFncallAssign(ctx *parser.FncallAssignContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitParen_fncall(ctx *parser.Paren_fncallContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFncall(ctx *parser.FncallContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFlagFnargs(ctx *parser.FlagFnargsContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitMultFnargs(ctx *parser.MultFnargsContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitDivFnargs(ctx *parser.DivFnargsContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitTermFnargs(ctx *parser.TermFnargsContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitLabelTerm(ctx *parser.LabelTermContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitLiteralTerm(ctx *parser.LiteralTermContext) interface{} {
	return v.Visit(ctx.TermContext)
}

func (v *BeepBoopVisitor) VisitMathTerm(ctx *parser.MathTermContext) interface{} {
	return v.Visit(ctx.Math())
}

func (v *BeepBoopVisitor) VisitParenfncallTerm(ctx *parser.ParenfncallTermContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitStructTerm(ctx *parser.StructTermContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitListTerm(ctx *parser.ListTermContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitList(ctx *parser.ListContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitListterm(ctx *parser.ListtermContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitGthanequalsCond(ctx *parser.GthanequalsCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitEqualsCond(ctx *parser.EqualsCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitOrCond(ctx *parser.OrCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitGthanCond(ctx *parser.GthanCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitAndCond(ctx *parser.AndCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitLthanCond(ctx *parser.LthanCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitLthanequalsCond(ctx *parser.LthanequalsCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitBoolCond(ctx *parser.BoolCondContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitSoloMath(ctx *parser.SoloMathContext) interface{} {
	return v.Visit(ctx.Num())
}

func (v *BeepBoopVisitor) VisitMultiplicativeMath(ctx *parser.MultiplicativeMathContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitAdditiveMath(ctx *parser.AdditiveMathContext) interface{} {

	va := GetVariable(v.Visit(ctx.Math(0)))
	vb := GetVariable(v.Visit(ctx.Math(1)))

	fmt.Println("add", va.AsInt(), vb.AsInt())

	if ctx.PLUS() != nil {
		return va.Plus(vb)
	} else {
		return va.Minus(vb)
	}
}

func (v *BeepBoopVisitor) VisitUnarySubMath(ctx *parser.UnarySubMathContext) interface{} {
	va := GetVariable(v.Visit(ctx.Math()))
	total := va.Negate()
	fmt.Println("un", total.AsInt())

	if total.Is(UNKNOWN) {
		v.errs = append(v.errs, fmt.Errorf("unexpected UNKNOWN value when attempting to negate at %s", ctx.GetText()))
	}

	return total
}

func (v *BeepBoopVisitor) VisitWordsLiteral(ctx *parser.WordsLiteralContext) interface{} {
	return v.Visit(ctx.Words())
}

func (v *BeepBoopVisitor) VisitIntNum(ctx *parser.IntNumContext) interface{} {
	val, err := strconv.Atoi(ctx.GetText())
	v.errs = append(v.errs, err)

	return NewVariable(INT, val)
}

func (v *BeepBoopVisitor) VisitFloatNum(ctx *parser.FloatNumContext) interface{} {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	v.errs = append(v.errs, err)

	return NewVariable(FLOAT, val)
}

func (v *BeepBoopVisitor) VisitLetterWords(ctx *parser.LetterWordsContext) interface{} {
	return ctx.GetText()
}

func (v *BeepBoopVisitor) VisitStringWords(ctx *parser.StringWordsContext) interface{} {
	return ctx.GetText()
}

func (v *BeepBoopVisitor) VisitQuotedWords(ctx *parser.QuotedWordsContext) interface{} {
	return util.RemoveQuotes(ctx.GetText())
}

func (v *BeepBoopVisitor) VisitPipeToFncallPipe(ctx *parser.PipeToFncallPipeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitPipeToPipe(ctx *parser.PipeToPipeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitTermPipe(ctx *parser.TermPipeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitNewlinePipe(ctx *parser.NewlinePipeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitFncallPipe(ctx *parser.FncallPipeContext) interface{} {
	panic("not implemented")
}

func (v *BeepBoopVisitor) VisitLabel(ctx *parser.LabelContext) interface{} {
	panic("not implemented")
}
