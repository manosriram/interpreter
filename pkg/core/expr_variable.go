package core

type Variable struct {
	Name string
}

func (v Variable) VisitVariableExpr() string {
	return Parenthesize(v.Name)
}

func (v Variable) Accept() interface{} {
	return v.VisitVariableExpr()
}
