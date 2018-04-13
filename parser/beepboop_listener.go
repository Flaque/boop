// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BeepBoopListener is a complete listener for a parse tree produced by BeepBoopParser.
type BeepBoopListener interface {
	antlr.ParseTreeListener

	// EnterBeepboop is called when entering the beepboop production.
	EnterBeepboop(c *BeepboopContext)

	// EnterMath is called when entering the math production.
	EnterMath(c *MathContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitBeepboop is called when exiting the beepboop production.
	ExitBeepboop(c *BeepboopContext)

	// ExitMath is called when exiting the math production.
	ExitMath(c *MathContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
