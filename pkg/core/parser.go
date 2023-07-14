package core

import (
	"interpreter/pkg/data"
	"log"
)

type Parser struct {
	Tokens  []*data.Token
	Current int32
	Size    int32
}

func NewParser(tokens []*data.Token, size int32) *Parser {
	return &Parser{
		Tokens:  tokens,
		Current: 0,
		Size:    size,
	}
}

func (p *Parser) is_end() bool {
	return p.peek().Tp == data.END_OF_FILE
}

func (p *Parser) advance() *data.Token {
	current := p.Current
	p.Current++

	return p.Tokens[current]
}

func (p *Parser) peek() *data.Token {
	if len(p.Tokens) == 0 {
		return nil
	}
	return p.Tokens[p.Current]
}

func (p *Parser) previous() *data.Token {
	if len(p.Tokens) == 0 {
		return nil
	}
	return p.Tokens[p.Current-1]
}

func (p *Parser) check(token_type data.TOKEN_TYPE) bool {
	if p.is_end() {
		return false
	}
	return p.peek().Tp == token_type
}

func (p *Parser) match(types ...data.TOKEN_TYPE) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(token_type data.TOKEN_TYPE, message string) {
	if p.check(token_type) {
		p.advance()
	}

	log.Fatal(p.peek(), message)
}

func (p *Parser) primary() Expr {
	if p.match(data.FALSE) {
		return Literal{Value: false}
	}
	if p.match(data.TRUE) {
		return Literal{Value: true}
	}
	if p.match(data.NIL) {
		return Literal{Value: nil}
	}
	if p.match(data.NUMBER, data.STRING) {
		return Literal{Value: p.previous().Literal}
	}

	if p.match(data.LEFT_PAREN) {
		expr := p.expression()
		p.consume(data.RIGHT_PAREN, "Expect ')' after expression")
		return Grouping{
			Expression: expr,
		}
	}

	log.Fatal(p.peek(), "Expect expression")
	return nil
}

func (p *Parser) unary() Expr {
	if p.match(data.BANG, data.MINUS) {
		operator := p.previous()
		right := p.unary()
		return Unary{
			Operator: *operator,
			Right:    right,
		}
	}
	return p.primary()
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(data.SLASH, data.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = Binary{
			Operator: *operator,
			Right:    right,
			Left:     expr,
		}
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(data.MINUS, data.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = Binary{
			Operator: *operator,
			Right:    right,
			Left:     expr,
		}
	}
	return expr
}

func (p *Parser) comparision() Expr {
	expr := p.term()

	for p.match(data.GREATER, data.GREATER_EQUAL, data.LESS, data.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = Binary{
			Operator: *operator,
			Right:    right,
			Left:     expr,
		}
	}
	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparision()
	for p.match(data.BANG_EQUAL, data.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparision()
		expr = Binary{
			Left:     expr,
			Operator: *operator,
			Right:    right,
		}
	}
	return expr
}

// func (p *Parser) Primary() Expr {
// if p.match(data.FALSE) {
// return Literal{Value: false}
// }
// if p.match(data.TRUE) {
// return Literal{Value: true}
// }
// if p.match(data.NIL) {
// return Literal{Value: nil}
// }
// if p.match(data.NUMBER, data.STRING) {
// return Literal{Value: p.previous().Literal}
// }

// if p.match(data.LEFT_PAREN, data.RIGHT_PAREN) {

// }

// return nil
// }

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) Parse() Expr {
	return p.expression()
}
