// Code generated from C:/Users/Sheep/Documents/Projects/monkey-interpreter/monkey-antlr/parser/Monkey.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Monkey
import "github.com/antlr4-go/antlr/v4"

type BaseMonkeyVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMonkeyVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitLetStatement(ctx *LetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitUnaryOperatorExpression(ctx *UnaryOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitGroupedExpression(ctx *GroupedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitEqualityBinaryExpression(ctx *EqualityBinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitLesGreBinaryExpression(ctx *LesGreBinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitIndexOperatorExpression(ctx *IndexOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitAddSubBinaryExpression(ctx *AddSubBinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitArrayExpression(ctx *ArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyVisitor) VisitMulDivBinaryExpression(ctx *MulDivBinaryExpressionContext) interface{} {
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
