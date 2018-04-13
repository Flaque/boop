package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/beep/parser"
	s "github.com/golang-collections/collections/stack"
)

var stack *s.Stack

func RuleToInt(rule antlr.ParseTree) (int, error) {
	return strconv.Atoi(rule.GetText())
}

type BeepBoopVisitor struct {
	*parser.BaseBeepBoopVisitor
}

func (v *BeepBoopVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *BeepBoopVisitor) VisitBeepboop(ctx *parser.BeepboopContext) interface{} {
	return v.Visit(ctx.Statement(0))
}

func (v *BeepBoopVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {

	return v.Visit(ctx.Expr(0))
}

func (v *BeepBoopVisitor) VisitTermExpr(ctx *parser.TermExprContext) interface{} {
	return v.Visit(ctx.Term())
}

func (v *BeepBoopVisitor) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	a, oka := v.Visit(ctx.Expr(0)).(int)
	b, okb := v.Visit(ctx.Expr(1)).(int)

	total := 0

	if oka {
		total += a
	}

	if okb {
		total += b
	}

	return total
}

func (v *BeepBoopVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	i, _ := RuleToInt(ctx.INT())
	return i
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("File required as input")
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}

	stack = s.New()

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBeepBoopLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeepBoopParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	tree := p.Beepboop()

	visitor := BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}}
	x := visitor.Visit(tree)

	// Final output
	fmt.Println(x)
}
