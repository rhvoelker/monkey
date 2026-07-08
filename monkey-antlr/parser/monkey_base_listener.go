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

// EnterStmt is called when production stmt is entered.
func (s *BaseMonkeyListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseMonkeyListener) ExitStmt(ctx *StmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMonkeyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMonkeyListener) ExitExpr(ctx *ExprContext) {}

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
