package core

import "fmt"

type Literal struct {
	Value interface{}
}

func (l Literal) VisitLiteralExpr() string {
	if l.Value == nil {
		return "nil"
	}

	v := fmt.Sprint(l.Value)
	return v
}

func (l Literal) Accept() interface{} {
	return l.VisitLiteralExpr()
}
