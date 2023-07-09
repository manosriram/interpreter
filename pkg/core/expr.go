package core

type Expr interface {
	Accept() interface{}
}
