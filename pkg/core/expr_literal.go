package core

type Literal struct {
	Value interface{}
}

func (l Literal) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return expr.Value
}

func (l Literal) Accept() interface{} {
	return l.VisitLiteralExpr(l)
}
