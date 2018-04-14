package runtime

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/beep/parser"
)

func Run(filename string) interface{} {
	input, _ := antlr.NewFileStream(filename)
	lexer := parser.NewBeepBoopLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeepBoopParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	tree := p.Beepboop()

	stack := NewStack()

	visitor := BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}, &stack}
	return visitor.Visit(tree)
}
