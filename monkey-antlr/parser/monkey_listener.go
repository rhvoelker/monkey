// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

// MonkeyListener is a complete listener for a parse tree produced by MonkeyParser.
type MonkeyListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

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

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

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
