package src

import (
	"fmt"
)

/*
	Program = [ Statement1, Statement2, ... ]
	Statement = [ Expression1, Expression2, ... ]
*/

type Node interface {
	TokenLiteral() string
	// String() string
}

type Expression interface {
	Node
	ExpressionNode()
}

type Statement interface {
	Node
	StatementNode()
}

type Program struct {
	Statements []Statement
}

// func (p *Program) String() string {
// var out bytes.Buffer
// for _, s := range p.Statements {
// out.WriteString(s.String())
// }
// return out.String()
// }

func (p *Program) PrintStatements() {
	for _, statement := range p.Statements {
		typed_statement := statement.(interface{})
		switch typed_statement.(type) {
		case *LetStatement:
			stmt := typed_statement.(*LetStatement)
			fmt.Printf("statement_type = %v, identifier_name = %v, identifier_value = %v\n", stmt.Token, stmt.Value, stmt.Name.Value)
			break
		case *ReturnStatement:
			stmt := typed_statement.(*ReturnStatement)
			fmt.Printf("statement_type = %v, identifier_name = %v\n", stmt.Token, stmt.ReturnValue)
			break
		default:
			stmt := typed_statement.(*ExpressionStatement)
			fmt.Printf("statement_type = %v, expression = %v\n", stmt.Token, stmt.Expression)
			break
		}
	}
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token *Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Value
}

type ReturnStatement struct {
	Token       *Token
	ReturnValue Expression
}

func (r *ReturnStatement) StatementNode() {}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Value
}

type ExpressionStatement struct {
	Token      *Token
	Expression Expression
}

func (r *ExpressionStatement) StatementNode() {}

func (r *ExpressionStatement) TokenLiteral() string {
	return r.Token.Value
}

// func (ls *LetStatement) String() string {
// var out bytes.Buffer
// out.WriteString(ls.TokenLiteral() + " ")
// out.WriteString(ls.Name.String())
// out.WriteString(" = ")
// if ls.Value != nil {
// out.WriteString(ls.Value.String())
// }
// out.WriteString(";")
// return out.String()
// }
// func (rs *ReturnStatement) String() string {
// var out bytes.Buffer
// out.WriteString(rs.TokenLiteral() + " ")
// if rs.ReturnValue != nil {
// out.WriteString(rs.ReturnValue.String())
// }
// out.WriteString(";")
// return out.String()
// }

// func (es *ExpressionStatement) String() string {
// if es.Expression != nil {
// return es.Expression.String()
// }
// return ""
// }

type Identifier struct {
	Token *Token
	Value string
}

func (i *Identifier) ExpressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Value
}

func (i *Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token *Token
	Value int64
}

func (i *IntegerLiteral) ExpressionNode() {}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Value
}

type PrefixExpression struct {
	Token    *Token
	Operator string
	Right    Expression
}

func (i *PrefixExpression) ExpressionNode() {}

func (i *PrefixExpression) TokenLiteral() string {
	return i.Token.Value
}

type InfixExpression struct {
	Token    *Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) ExpressionNode() {}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Value
}
