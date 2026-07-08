// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

// BaseMonkeyListener is a complete listener for a parse tree produced by MonkeyParser.
type BaseMonkeyListener struct{}

var _ MonkeyListener = &BaseMonkeyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMonkeyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMonkeyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMonkeyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMonkeyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseMonkeyListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMonkeyListener) ExitProg(ctx *ProgContext) {}

// EnterLetStatement is called when production LetStatement is entered.
func (s *BaseMonkeyListener) EnterLetStatement(ctx *LetStatementContext) {}

// ExitLetStatement is called when production LetStatement is exited.
func (s *BaseMonkeyListener) ExitLetStatement(ctx *LetStatementContext) {}

// EnterReturnStatement is called when production ReturnStatement is entered.
func (s *BaseMonkeyListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production ReturnStatement is exited.
func (s *BaseMonkeyListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExpressionStatement is called when production ExpressionStatement is entered.
func (s *BaseMonkeyListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production ExpressionStatement is exited.
func (s *BaseMonkeyListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfExpression is called when production IfExpression is entered.
func (s *BaseMonkeyListener) EnterIfExpression(ctx *IfExpressionContext) {}

// ExitIfExpression is called when production IfExpression is exited.
func (s *BaseMonkeyListener) ExitIfExpression(ctx *IfExpressionContext) {}

// EnterUnaryOperatorExpression is called when production UnaryOperatorExpression is entered.
func (s *BaseMonkeyListener) EnterUnaryOperatorExpression(ctx *UnaryOperatorExpressionContext) {}

// ExitUnaryOperatorExpression is called when production UnaryOperatorExpression is exited.
func (s *BaseMonkeyListener) ExitUnaryOperatorExpression(ctx *UnaryOperatorExpressionContext) {}

// EnterBooleanLiteral is called when production BooleanLiteral is entered.
func (s *BaseMonkeyListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production BooleanLiteral is exited.
func (s *BaseMonkeyListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterGroupedExpression is called when production GroupedExpression is entered.
func (s *BaseMonkeyListener) EnterGroupedExpression(ctx *GroupedExpressionContext) {}

// ExitGroupedExpression is called when production GroupedExpression is exited.
func (s *BaseMonkeyListener) ExitGroupedExpression(ctx *GroupedExpressionContext) {}

// EnterEqualityBinaryExpression is called when production EqualityBinaryExpression is entered.
func (s *BaseMonkeyListener) EnterEqualityBinaryExpression(ctx *EqualityBinaryExpressionContext) {}

// ExitEqualityBinaryExpression is called when production EqualityBinaryExpression is exited.
func (s *BaseMonkeyListener) ExitEqualityBinaryExpression(ctx *EqualityBinaryExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseMonkeyListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseMonkeyListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterLesGreBinaryExpression is called when production LesGreBinaryExpression is entered.
func (s *BaseMonkeyListener) EnterLesGreBinaryExpression(ctx *LesGreBinaryExpressionContext) {}

// ExitLesGreBinaryExpression is called when production LesGreBinaryExpression is exited.
func (s *BaseMonkeyListener) ExitLesGreBinaryExpression(ctx *LesGreBinaryExpressionContext) {}

// EnterIndexOperatorExpression is called when production IndexOperatorExpression is entered.
func (s *BaseMonkeyListener) EnterIndexOperatorExpression(ctx *IndexOperatorExpressionContext) {}

// ExitIndexOperatorExpression is called when production IndexOperatorExpression is exited.
func (s *BaseMonkeyListener) ExitIndexOperatorExpression(ctx *IndexOperatorExpressionContext) {}

// EnterAddSubBinaryExpression is called when production AddSubBinaryExpression is entered.
func (s *BaseMonkeyListener) EnterAddSubBinaryExpression(ctx *AddSubBinaryExpressionContext) {}

// ExitAddSubBinaryExpression is called when production AddSubBinaryExpression is exited.
func (s *BaseMonkeyListener) ExitAddSubBinaryExpression(ctx *AddSubBinaryExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseMonkeyListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseMonkeyListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseMonkeyListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseMonkeyListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterArrayExpression is called when production ArrayExpression is entered.
func (s *BaseMonkeyListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production ArrayExpression is exited.
func (s *BaseMonkeyListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterCallExpression is called when production CallExpression is entered.
func (s *BaseMonkeyListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production CallExpression is exited.
func (s *BaseMonkeyListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterIntegerLiteral is called when production IntegerLiteral is entered.
func (s *BaseMonkeyListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production IntegerLiteral is exited.
func (s *BaseMonkeyListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterMulDivBinaryExpression is called when production MulDivBinaryExpression is entered.
func (s *BaseMonkeyListener) EnterMulDivBinaryExpression(ctx *MulDivBinaryExpressionContext) {}

// ExitMulDivBinaryExpression is called when production MulDivBinaryExpression is exited.
func (s *BaseMonkeyListener) ExitMulDivBinaryExpression(ctx *MulDivBinaryExpressionContext) {}

// EnterLet_stmt is called when production let_stmt is entered.
func (s *BaseMonkeyListener) EnterLet_stmt(ctx *Let_stmtContext) {}

// ExitLet_stmt is called when production let_stmt is exited.
func (s *BaseMonkeyListener) ExitLet_stmt(ctx *Let_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseMonkeyListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseMonkeyListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterExpression_stmt is called when production expression_stmt is entered.
func (s *BaseMonkeyListener) EnterExpression_stmt(ctx *Expression_stmtContext) {}

// ExitExpression_stmt is called when production expression_stmt is exited.
func (s *BaseMonkeyListener) ExitExpression_stmt(ctx *Expression_stmtContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BaseMonkeyListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BaseMonkeyListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterArgs_list is called when production args_list is entered.
func (s *BaseMonkeyListener) EnterArgs_list(ctx *Args_listContext) {}

// ExitArgs_list is called when production args_list is exited.
func (s *BaseMonkeyListener) ExitArgs_list(ctx *Args_listContext) {}
