package src

import (
	"fmt"
	"log"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precendences = map[TOKEN_TYPE]int{
	EQUAL:      EQUALS,
	BANG_EQUAL: EQUALS,
	LESS:       LESSGREATER,
	GREATER:    LESSGREATER,
	PLUS:       SUM,
	MINUS:      SUM,
	SLASH:      PRODUCT,
	STAR:       PRODUCT,
}

type (
	prefixParseFunction func() Expression
	infixParseFunction  func(Expression) Expression
)

type Parser struct {
	L *Lexer

	CurrentToken *Token
	PeekToken    *Token

	prefixParseFunctions map[TOKEN_TYPE]prefixParseFunction
	infixParseFunctions  map[TOKEN_TYPE]infixParseFunction

	errors []string
}

func (p *Parser) PeekPrecedence() int {
	return precendences[p.PeekToken.Type]
}

func (p *Parser) CurrentPrecedence() int {
	return precendences[p.CurrentToken.Type]
}

func (p *Parser) RegisterPrefixParseFunction(t TOKEN_TYPE, parse_function prefixParseFunction) {
	p.prefixParseFunctions[t] = parse_function
}

func (p *Parser) RegisterInfixParseFunction(t TOKEN_TYPE, parse_function infixParseFunction) {
	p.infixParseFunctions[t] = parse_function
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		L:            l,
		CurrentToken: nil,
		PeekToken:    nil,
	}
	p.CurrentToken = p.L.PeekToken()
	p.PeekToken = p.L.PeekNextToken()

	p.prefixParseFunctions = make(map[TOKEN_TYPE]prefixParseFunction)
	p.infixParseFunctions = make(map[TOKEN_TYPE]infixParseFunction)

	p.RegisterPrefixParseFunction(IDENT, p.parseIdentifier)
	p.RegisterPrefixParseFunction(INT, p.parseIntegerLiteral)
	p.RegisterPrefixParseFunction(BANG, p.parsePrefixExpression)
	p.RegisterPrefixParseFunction(MINUS, p.parsePrefixExpression)

	p.RegisterInfixParseFunction(PLUS, p.parseInfixExpression)
	p.RegisterInfixParseFunction(EQUAL, p.parseInfixExpression)

	return p
}

func (p *Parser) NextToken() {
	if len(p.errors) > 0 {
		log.Fatal(p.errors[0])
	}

	p.CurrentToken = p.PeekToken
	p.PeekToken = p.L.PeekNextToken()
}

func (p *Parser) expectPeek(token_type TOKEN_TYPE) bool {
	if p.CurrentToken.Type != token_type {
		p.errors = append(p.errors, fmt.Sprintf("expected token '%v' but got '%v'\n", token_type, p.CurrentToken.Value))
		return false
	}
	return true
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expr := &InfixExpression{
		Token:    p.CurrentToken,
		Operator: p.CurrentToken.Value,
		Left:     left,
	}

	precedence := p.CurrentPrecedence()
	p.NextToken()
	expr.Right = p.parseExpression(precedence)

	return expr
}

func (p *Parser) parsePrefixExpression() Expression {
	expr := &PrefixExpression{
		Token:    p.CurrentToken,
		Operator: p.CurrentToken.Value,
	}
	p.NextToken()
	expr.Right = p.parseExpression(PREFIX)

	return expr
}

func (p *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: p.CurrentToken}

	value, err := strconv.ParseInt(p.CurrentToken.Value, 0, 1)
	if err != nil {
		msg := fmt.Sprintf("could not parse %s as integer", p.CurrentToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value

	return lit
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Value}
}

func (p *Parser) parseLetStatement() Statement {
	stmt := &LetStatement{Token: p.CurrentToken}
	p.NextToken()

	name := &Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Value}
	stmt.Name = name
	p.NextToken()

	p.expectPeek(EQUAL)

	p.NextToken()
	for p.CurrentToken != nil && p.CurrentToken.Type != SEMICOLON {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() Statement {
	stmt := &ReturnStatement{Token: p.CurrentToken}
	p.NextToken()

	// TODO: parse expression

	for p.CurrentToken != nil && p.CurrentToken.Type != SEMICOLON {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precendence int) Expression {
	prefix := p.prefixParseFunctions[p.CurrentToken.Type]
	if prefix == nil {
		return nil
	}

	left_expr := prefix()

	if p.PeekToken.Type != SEMICOLON && precendence < p.PeekPrecedence() {
		infix := p.infixParseFunctions[p.CurrentToken.Type]
		if infix == nil {
			return left_expr
		}
		p.NextToken()
		left_expr = infix(left_expr)
	}

	return left_expr
}

func (p *Parser) parseExpressionStatement() Statement {
	stmt := &ExpressionStatement{Token: p.CurrentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.PeekToken.Type == SEMICOLON {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseStatement() Statement {
	switch p.CurrentToken.Type {
	case LET:
		return p.parseLetStatement()
	case RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *Program {
	pr := &Program{}
	pr.Statements = []Statement{}

	for p.CurrentToken != nil && p.CurrentToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			pr.Statements = append(pr.Statements, stmt)
		}

		p.NextToken()
	}

	return pr
}
