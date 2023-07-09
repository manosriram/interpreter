package core

import (
	"bytes"
)

func PrintExpr(e Expr) string {
	switch t := e.(type) {
	case *Binary:
		return Parenthesize(t.Operator.Lexeme, t.Left, t.Right)
	case *Unary:
		return Parenthesize(t.Operator.Lexeme, t.Right)
	case *Literal:
		if t.Value == nil {
			return "nil"
		}
		return t.Value.(string)
	case *Grouping:
		return Parenthesize("group", t.Expression)
	default:
		return "nil"
	}
}

func Parenthesize(name string, exprs ...Expr) string {
	buf := bytes.Buffer{}

	buf.Write([]byte("("))
	buf.Write([]byte(name))

	for _, v := range exprs {
		buf.Write([]byte(" "))
		buf.Write([]byte(v.Accept().([]byte)))
		// buf.Write([]byte(PrintExpr(v)))
	}

	buf.Write([]byte(")"))
	return buf.String()
}
