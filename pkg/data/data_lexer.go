package data

type TOKEN_TYPE int32

const (
	LEFT_PAREN TOKEN_TYPE = iota
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
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

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
	Line    int32
}

func NewToken(tp TOKEN_TYPE, lexeme string, literal interface{}, line int32) *Token {
	return &Token{
		Tp:      tp,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
