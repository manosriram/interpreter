package core

import (
	"bytes"
	"fmt"
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
		v := fmt.Sprint(t.Value)
		return v
	case *Grouping:
		return Parenthesize("group", t.Expression)
	default:
		return "nil"
	}
}

func Parenthesize(name string, exprs ...Expr) string {
	buf := bytes.Buffer{}

	buf.WriteString("(")
	buf.WriteString(name)

	for _, v := range exprs {
		buf.WriteString(" ")
		buf.WriteString(v.Accept().(string))
	}

	buf.WriteString(")")
	return buf.String()
}
