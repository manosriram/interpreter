package core

import (
	"interpreter/pkg/data"
)

type Unary struct {
	Operator data.Token
	Right    Expr
}

func (u Unary) VisitUnaryExpr(expr Unary) string {
	return Parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (u Unary) Accept() interface{} {
	return u.VisitUnaryExpr(u)
}
