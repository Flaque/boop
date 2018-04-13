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

type BeepBoopVisitor struct {
	*parser.BaseBeepBoopVisitor
}

func (v *BeepBoopVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *BeepBoopVisitor) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	fmt.Println("Math")

	a, _ := strconv.Atoi(ctx.Expr(0).GetText())
	b, _ := strconv.Atoi(ctx.Expr(1).GetText())

	fmt.Println(a + b)

	return v.VisitChildren(ctx) // TODO: Change to v.Visit(INT) + v.VISIT(INT) or something
}

func (v *BeepBoopVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	return ctx.INT()
}

func main() {
	stack = s.New()

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBeepBoopLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeepBoopParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	tree := p.Expr()

	visitor := BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}}
	visitor.Visit(tree)
}
