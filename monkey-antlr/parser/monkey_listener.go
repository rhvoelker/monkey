// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

// MonkeyListener is a complete listener for a parse tree produced by MonkeyParser.
type MonkeyListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLetStatement is called when entering the LetStatement production.
	EnterLetStatement(c *LetStatementContext)

	// EnterReturnStatement is called when entering the ReturnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIfExpression is called when entering the IfExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterUnaryOperatorExpression is called when entering the UnaryOperatorExpression production.
	EnterUnaryOperatorExpression(c *UnaryOperatorExpressionContext)

	// EnterBooleanLiteral is called when entering the BooleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterGroupedExpression is called when entering the GroupedExpression production.
	EnterGroupedExpression(c *GroupedExpressionContext)

	// EnterEqualityBinaryExpression is called when entering the EqualityBinaryExpression production.
	EnterEqualityBinaryExpression(c *EqualityBinaryExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterLesGreBinaryExpression is called when entering the LesGreBinaryExpression production.
	EnterLesGreBinaryExpression(c *LesGreBinaryExpressionContext)

	// EnterIndexOperatorExpression is called when entering the IndexOperatorExpression production.
	EnterIndexOperatorExpression(c *IndexOperatorExpressionContext)

	// EnterAddSubBinaryExpression is called when entering the AddSubBinaryExpression production.
	EnterAddSubBinaryExpression(c *AddSubBinaryExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterArrayExpression is called when entering the ArrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterCallExpression is called when entering the CallExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterIntegerLiteral is called when entering the IntegerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterMulDivBinaryExpression is called when entering the MulDivBinaryExpression production.
	EnterMulDivBinaryExpression(c *MulDivBinaryExpressionContext)

	// EnterLet_stmt is called when entering the let_stmt production.
	EnterLet_stmt(c *Let_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterExpression_stmt is called when entering the expression_stmt production.
	EnterExpression_stmt(c *Expression_stmtContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterArgs_list is called when entering the args_list production.
	EnterArgs_list(c *Args_listContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLetStatement is called when exiting the LetStatement production.
	ExitLetStatement(c *LetStatementContext)

	// ExitReturnStatement is called when exiting the ReturnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIfExpression is called when exiting the IfExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitUnaryOperatorExpression is called when exiting the UnaryOperatorExpression production.
	ExitUnaryOperatorExpression(c *UnaryOperatorExpressionContext)

	// ExitBooleanLiteral is called when exiting the BooleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitGroupedExpression is called when exiting the GroupedExpression production.
	ExitGroupedExpression(c *GroupedExpressionContext)

	// ExitEqualityBinaryExpression is called when exiting the EqualityBinaryExpression production.
	ExitEqualityBinaryExpression(c *EqualityBinaryExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitLesGreBinaryExpression is called when exiting the LesGreBinaryExpression production.
	ExitLesGreBinaryExpression(c *LesGreBinaryExpressionContext)

	// ExitIndexOperatorExpression is called when exiting the IndexOperatorExpression production.
	ExitIndexOperatorExpression(c *IndexOperatorExpressionContext)

	// ExitAddSubBinaryExpression is called when exiting the AddSubBinaryExpression production.
	ExitAddSubBinaryExpression(c *AddSubBinaryExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitArrayExpression is called when exiting the ArrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitCallExpression is called when exiting the CallExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitIntegerLiteral is called when exiting the IntegerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitMulDivBinaryExpression is called when exiting the MulDivBinaryExpression production.
	ExitMulDivBinaryExpression(c *MulDivBinaryExpressionContext)

	// ExitLet_stmt is called when exiting the let_stmt production.
	ExitLet_stmt(c *Let_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitExpression_stmt is called when exiting the expression_stmt production.
	ExitExpression_stmt(c *Expression_stmtContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitArgs_list is called when exiting the args_list production.
	ExitArgs_list(c *Args_listContext)
}
