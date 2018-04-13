// Code generated from parser/Do.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Do

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DoListener is a complete listener for a parse tree produced by DoParser.
type DoListener interface {
	antlr.ParseTreeListener

	// EnterDo is called when entering the do production.
	EnterDo(c *DoContext)

	// EnterMath is called when entering the math production.
	EnterMath(c *MathContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitDo is called when exiting the do production.
	ExitDo(c *DoContext)

	// ExitMath is called when exiting the math production.
	ExitMath(c *MathContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
