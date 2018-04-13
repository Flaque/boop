package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/beep/parser"
)

type TreeShapeListener struct {
	*parser.BaseBeepBoopListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBeepBoopLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeepBoopParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	tree := p.Math()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
