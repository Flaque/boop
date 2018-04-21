// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BeepBoop

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBeepBoopListener is a complete listener for a parse tree produced by BeepBoopParser.
type BaseBeepBoopListener struct{}

var _ BeepBoopListener = &BaseBeepBoopListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBeepBoopListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBeepBoopListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBeepBoopListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBeepBoopListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBoop is called when production boop is entered.
func (s *BaseBeepBoopListener) EnterBoop(ctx *BoopContext) {}

// ExitBoop is called when production boop is exited.
func (s *BaseBeepBoopListener) ExitBoop(ctx *BoopContext) {}

// EnterFuncdefCode is called when production funcdefCode is entered.
func (s *BaseBeepBoopListener) EnterFuncdefCode(ctx *FuncdefCodeContext) {}

// ExitFuncdefCode is called when production funcdefCode is exited.
func (s *BaseBeepBoopListener) ExitFuncdefCode(ctx *FuncdefCodeContext) {}

// EnterStatementCode is called when production statementCode is entered.
func (s *BaseBeepBoopListener) EnterStatementCode(ctx *StatementCodeContext) {}

// ExitStatementCode is called when production statementCode is exited.
func (s *BaseBeepBoopListener) ExitStatementCode(ctx *StatementCodeContext) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BaseBeepBoopListener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BaseBeepBoopListener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterFuncguts is called when production funcguts is entered.
func (s *BaseBeepBoopListener) EnterFuncguts(ctx *FuncgutsContext) {}

// ExitFuncguts is called when production funcguts is exited.
func (s *BaseBeepBoopListener) ExitFuncguts(ctx *FuncgutsContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseBeepBoopListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseBeepBoopListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseBeepBoopListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseBeepBoopListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterFncallStatement is called when production fncallStatement is entered.
func (s *BaseBeepBoopListener) EnterFncallStatement(ctx *FncallStatementContext) {}

// ExitFncallStatement is called when production fncallStatement is exited.
func (s *BaseBeepBoopListener) ExitFncallStatement(ctx *FncallStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseBeepBoopListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseBeepBoopListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterPipeStatement is called when production pipeStatement is entered.
func (s *BaseBeepBoopListener) EnterPipeStatement(ctx *PipeStatementContext) {}

// ExitPipeStatement is called when production pipeStatement is exited.
func (s *BaseBeepBoopListener) ExitPipeStatement(ctx *PipeStatementContext) {}

// EnterNewlineStatement is called when production newlineStatement is entered.
func (s *BaseBeepBoopListener) EnterNewlineStatement(ctx *NewlineStatementContext) {}

// ExitNewlineStatement is called when production newlineStatement is exited.
func (s *BaseBeepBoopListener) ExitNewlineStatement(ctx *NewlineStatementContext) {}

// EnterIfstat is called when production ifstat is entered.
func (s *BaseBeepBoopListener) EnterIfstat(ctx *IfstatContext) {}

// ExitIfstat is called when production ifstat is exited.
func (s *BaseBeepBoopListener) ExitIfstat(ctx *IfstatContext) {}

// EnterReturnstat is called when production returnstat is entered.
func (s *BaseBeepBoopListener) EnterReturnstat(ctx *ReturnstatContext) {}

// ExitReturnstat is called when production returnstat is exited.
func (s *BaseBeepBoopListener) ExitReturnstat(ctx *ReturnstatContext) {}

// EnterStructexpr is called when production structexpr is entered.
func (s *BaseBeepBoopListener) EnterStructexpr(ctx *StructexprContext) {}

// ExitStructexpr is called when production structexpr is exited.
func (s *BaseBeepBoopListener) ExitStructexpr(ctx *StructexprContext) {}

// EnterAssignstat is called when production assignstat is entered.
func (s *BaseBeepBoopListener) EnterAssignstat(ctx *AssignstatContext) {}

// ExitAssignstat is called when production assignstat is exited.
func (s *BaseBeepBoopListener) ExitAssignstat(ctx *AssignstatContext) {}

// EnterExprAssign is called when production exprAssign is entered.
func (s *BaseBeepBoopListener) EnterExprAssign(ctx *ExprAssignContext) {}

// ExitExprAssign is called when production exprAssign is exited.
func (s *BaseBeepBoopListener) ExitExprAssign(ctx *ExprAssignContext) {}

// EnterFncallAssign is called when production fncallAssign is entered.
func (s *BaseBeepBoopListener) EnterFncallAssign(ctx *FncallAssignContext) {}

// ExitFncallAssign is called when production fncallAssign is exited.
func (s *BaseBeepBoopListener) ExitFncallAssign(ctx *FncallAssignContext) {}

// EnterParen_fncall is called when production paren_fncall is entered.
func (s *BaseBeepBoopListener) EnterParen_fncall(ctx *Paren_fncallContext) {}

// ExitParen_fncall is called when production paren_fncall is exited.
func (s *BaseBeepBoopListener) ExitParen_fncall(ctx *Paren_fncallContext) {}

// EnterFncall is called when production fncall is entered.
func (s *BaseBeepBoopListener) EnterFncall(ctx *FncallContext) {}

// ExitFncall is called when production fncall is exited.
func (s *BaseBeepBoopListener) ExitFncall(ctx *FncallContext) {}

// EnterLabelTerm is called when production labelTerm is entered.
func (s *BaseBeepBoopListener) EnterLabelTerm(ctx *LabelTermContext) {}

// ExitLabelTerm is called when production labelTerm is exited.
func (s *BaseBeepBoopListener) ExitLabelTerm(ctx *LabelTermContext) {}

// EnterLiteralTerm is called when production literalTerm is entered.
func (s *BaseBeepBoopListener) EnterLiteralTerm(ctx *LiteralTermContext) {}

// ExitLiteralTerm is called when production literalTerm is exited.
func (s *BaseBeepBoopListener) ExitLiteralTerm(ctx *LiteralTermContext) {}

// EnterMathTerm is called when production mathTerm is entered.
func (s *BaseBeepBoopListener) EnterMathTerm(ctx *MathTermContext) {}

// ExitMathTerm is called when production mathTerm is exited.
func (s *BaseBeepBoopListener) ExitMathTerm(ctx *MathTermContext) {}

// EnterParenfncallTerm is called when production parenfncallTerm is entered.
func (s *BaseBeepBoopListener) EnterParenfncallTerm(ctx *ParenfncallTermContext) {}

// ExitParenfncallTerm is called when production parenfncallTerm is exited.
func (s *BaseBeepBoopListener) ExitParenfncallTerm(ctx *ParenfncallTermContext) {}

// EnterStructTerm is called when production structTerm is entered.
func (s *BaseBeepBoopListener) EnterStructTerm(ctx *StructTermContext) {}

// ExitStructTerm is called when production structTerm is exited.
func (s *BaseBeepBoopListener) ExitStructTerm(ctx *StructTermContext) {}

// EnterListTerm is called when production listTerm is entered.
func (s *BaseBeepBoopListener) EnterListTerm(ctx *ListTermContext) {}

// ExitListTerm is called when production listTerm is exited.
func (s *BaseBeepBoopListener) ExitListTerm(ctx *ListTermContext) {}

// EnterList is called when production list is entered.
func (s *BaseBeepBoopListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseBeepBoopListener) ExitList(ctx *ListContext) {}

// EnterListterm is called when production listterm is entered.
func (s *BaseBeepBoopListener) EnterListterm(ctx *ListtermContext) {}

// ExitListterm is called when production listterm is exited.
func (s *BaseBeepBoopListener) ExitListterm(ctx *ListtermContext) {}

// EnterGthanequalsCond is called when production gthanequalsCond is entered.
func (s *BaseBeepBoopListener) EnterGthanequalsCond(ctx *GthanequalsCondContext) {}

// ExitGthanequalsCond is called when production gthanequalsCond is exited.
func (s *BaseBeepBoopListener) ExitGthanequalsCond(ctx *GthanequalsCondContext) {}

// EnterEqualsCond is called when production equalsCond is entered.
func (s *BaseBeepBoopListener) EnterEqualsCond(ctx *EqualsCondContext) {}

// ExitEqualsCond is called when production equalsCond is exited.
func (s *BaseBeepBoopListener) ExitEqualsCond(ctx *EqualsCondContext) {}

// EnterOrCond is called when production orCond is entered.
func (s *BaseBeepBoopListener) EnterOrCond(ctx *OrCondContext) {}

// ExitOrCond is called when production orCond is exited.
func (s *BaseBeepBoopListener) ExitOrCond(ctx *OrCondContext) {}

// EnterGthanCond is called when production gthanCond is entered.
func (s *BaseBeepBoopListener) EnterGthanCond(ctx *GthanCondContext) {}

// ExitGthanCond is called when production gthanCond is exited.
func (s *BaseBeepBoopListener) ExitGthanCond(ctx *GthanCondContext) {}

// EnterAndCond is called when production andCond is entered.
func (s *BaseBeepBoopListener) EnterAndCond(ctx *AndCondContext) {}

// ExitAndCond is called when production andCond is exited.
func (s *BaseBeepBoopListener) ExitAndCond(ctx *AndCondContext) {}

// EnterLthanCond is called when production lthanCond is entered.
func (s *BaseBeepBoopListener) EnterLthanCond(ctx *LthanCondContext) {}

// ExitLthanCond is called when production lthanCond is exited.
func (s *BaseBeepBoopListener) ExitLthanCond(ctx *LthanCondContext) {}

// EnterLthanequalsCond is called when production lthanequalsCond is entered.
func (s *BaseBeepBoopListener) EnterLthanequalsCond(ctx *LthanequalsCondContext) {}

// ExitLthanequalsCond is called when production lthanequalsCond is exited.
func (s *BaseBeepBoopListener) ExitLthanequalsCond(ctx *LthanequalsCondContext) {}

// EnterBoolCond is called when production boolCond is entered.
func (s *BaseBeepBoopListener) EnterBoolCond(ctx *BoolCondContext) {}

// ExitBoolCond is called when production boolCond is exited.
func (s *BaseBeepBoopListener) ExitBoolCond(ctx *BoolCondContext) {}

// EnterSoloMath is called when production soloMath is entered.
func (s *BaseBeepBoopListener) EnterSoloMath(ctx *SoloMathContext) {}

// ExitSoloMath is called when production soloMath is exited.
func (s *BaseBeepBoopListener) ExitSoloMath(ctx *SoloMathContext) {}

// EnterAdditiveMath is called when production additiveMath is entered.
func (s *BaseBeepBoopListener) EnterAdditiveMath(ctx *AdditiveMathContext) {}

// ExitAdditiveMath is called when production additiveMath is exited.
func (s *BaseBeepBoopListener) ExitAdditiveMath(ctx *AdditiveMathContext) {}

// EnterUnarySubMath is called when production unarySubMath is entered.
func (s *BaseBeepBoopListener) EnterUnarySubMath(ctx *UnarySubMathContext) {}

// ExitUnarySubMath is called when production unarySubMath is exited.
func (s *BaseBeepBoopListener) ExitUnarySubMath(ctx *UnarySubMathContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseBeepBoopListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseBeepBoopListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNum is called when production num is entered.
func (s *BaseBeepBoopListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseBeepBoopListener) ExitNum(ctx *NumContext) {}

// EnterBoolexpr is called when production boolexpr is entered.
func (s *BaseBeepBoopListener) EnterBoolexpr(ctx *BoolexprContext) {}

// ExitBoolexpr is called when production boolexpr is exited.
func (s *BaseBeepBoopListener) ExitBoolexpr(ctx *BoolexprContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BaseBeepBoopListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BaseBeepBoopListener) ExitPipe(ctx *PipeContext) {}
