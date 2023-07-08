package data

type TOKEN_TYPE int32

const (
	LEFT_PARAM TOKEN_TYPE = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUA
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER_INTEGER
	NUMBER_DOUBLE

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	END_OF_FILE
)

type Token struct {
	Tp      TOKEN_TYPE
	Lexeme  string
	Literal interface{}
}

func NewToken(tp TOKEN_TYPE, lexeme string, literal interface{}) *Token {
	return &Token{
		Tp:      tp,
		Lexeme:  lexeme,
		Literal: literal,
	}
}
