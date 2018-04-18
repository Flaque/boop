package runtime

import (
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flaque/boop/parser"
)

func Run(code string, out io.Writer) interface{} {

	// Code should have no new lines in the front
	code = strings.Trim(code, "\n")

	// Code must have a new line at the end for the grammar
	// So we'll fix it here just in case
	code += "\n"

	// Setup ANTLR's parser
	input := antlr.NewInputStream(code)
	lexer := parser.NewBeepBoopLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeepBoopParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Beepboop()

	// Stack of Hashtree's to store variables
	stack := NewTree(nil)
	logger := NewLogger(out)

	visitor := BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}, &stack, &logger}
	return visitor.Visit(tree)
}
