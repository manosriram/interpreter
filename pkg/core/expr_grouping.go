package core

type Grouping struct {
	Expression Expr
}

func (g Grouping) VisitGroupingExpr(expr Grouping) string {
	return Parenthesize("group", expr.Expression)
}

func (g Grouping) Accept() interface{} {
	return g.VisitGroupingExpr(g)
}
