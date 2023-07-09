package core

import (
	"interpreter/pkg/data"
)

type Binary struct {
	Left     Expr
	Operator data.Token
	Right    Expr
}

func (b Binary) VisitBinaryExpr(expr Binary) interface{} {
	return Parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (b Binary) Accept() interface{} {
	return b.VisitBinaryExpr(b)
}
