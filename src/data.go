package src

type TOKEN_TYPE string
type LEX_TYPE string

type Lexer struct {
	Input   string
	Start   int32
	Current int32
	Ch      byte

	TokenPosition int32
	Tokens        []*Token
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input:         input,
		Start:         0,
		Current:       0,
		TokenPosition: 0,
		Tokens:        make([]*Token, 0),
	}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	STRING = "STRING"
	INT    = "INT"

	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"

	// One or two character tokens.
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Keywords.
	FUNCTION = "fn"
	LET      = "let"
	IF       = "if"
	ELSE     = "else"
	TRUE     = "true"
	FALSE    = "false"
	RETURN   = "return"
)

var Keywords map[string]TOKEN_TYPE = map[string]TOKEN_TYPE{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

type Token struct {
	Type  TOKEN_TYPE
	Value string
}

func NewToken(tp TOKEN_TYPE, value string) *Token {
	return &Token{
		Type:  tp,
		Value: value,
	}
}
