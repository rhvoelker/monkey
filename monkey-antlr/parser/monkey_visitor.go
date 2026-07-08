// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MonkeyParser.
type MonkeyVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MonkeyParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by MonkeyParser#LetStatement.
	VisitLetStatement(ctx *LetStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#ExpressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#IfExpression.
	VisitIfExpression(ctx *IfExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#UnaryOperatorExpression.
	VisitUnaryOperatorExpression(ctx *UnaryOperatorExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#BooleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#GroupedExpression.
	VisitGroupedExpression(ctx *GroupedExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#EqualityBinaryExpression.
	VisitEqualityBinaryExpression(ctx *EqualityBinaryExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#LesGreBinaryExpression.
	VisitLesGreBinaryExpression(ctx *LesGreBinaryExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#IndexOperatorExpression.
	VisitIndexOperatorExpression(ctx *IndexOperatorExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#AddSubBinaryExpression.
	VisitAddSubBinaryExpression(ctx *AddSubBinaryExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#ArrayExpression.
	VisitArrayExpression(ctx *ArrayExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#CallExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#IntegerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#MulDivBinaryExpression.
	VisitMulDivBinaryExpression(ctx *MulDivBinaryExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#let_stmt.
	VisitLet_stmt(ctx *Let_stmtContext) interface{}

	// Visit a parse tree produced by MonkeyParser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expression_stmt.
	VisitExpression_stmt(ctx *Expression_stmtContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expr_list.
	VisitExpr_list(ctx *Expr_listContext) interface{}

	// Visit a parse tree produced by MonkeyParser#args_list.
	VisitArgs_list(ctx *Args_listContext) interface{}
}
