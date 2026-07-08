// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

type BaseMonkeyVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMonkeyVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitLet_stmt(ctx *Let_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitExpression_stmt(ctx *Expression_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitExpr_list(ctx *Expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitArgs_list(ctx *Args_listContext) interface{} {
	return v.VisitChildren(ctx)
}
