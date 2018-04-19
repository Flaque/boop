package runtime

import (
	"io"
	"strings"

	"github.com/Flaque/boop/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func prepCode(code string) string {
	// Code should have no new lines in the front
	code = strings.Trim(code, "\n")

	// Code must have a new line at the end for the grammar
	// So we'll fix it here just in case
	code += "\n"

	return code
}

func getParser(code string) *parser.BeepBoopParser {
	input := antlr.NewInputStream(code)
	lexer := parser.NewBeepBoopLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBeepBoopParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	return p
}

type Listener struct {
	*antlr.BaseParseTreeListener
}

func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// Do nothing
}

// Validate checks code is syntactically correct
func Validate(code string) {
	code = prepCode(code)

	p := getParser(code)
	listener := &Listener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Beepboop())
}

// Run runs code and returns visiter output
func Run(code string, out io.Writer) interface{} {
	code = prepCode(code)

	// Setup ANTLR's parser
	p := getParser(code)
	tree := p.Beepboop()

	// Stack of Hashtree's to store variables
	stack := NewTree(nil)
	logger := NewLogger(out)
	visits := []string{}

	visitor := BeepBoopVisitor{&parser.BaseBeepBoopVisitor{}, &stack, &logger, visits}
	output := visitor.Visit(tree)

	return output
}
