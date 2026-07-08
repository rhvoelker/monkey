// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MonkeyParser.
type MonkeyVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MonkeyParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by MonkeyParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

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
