package core

type Visitor interface {
	VisitBinaryExpr(expr Binary) interface{}
}
