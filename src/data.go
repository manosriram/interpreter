package src

type TOKEN_TYPE string
type LEX_TYPE string

type Context struct {
	Start   int32
	Current int32
	Line    int32
	Tokens  []*Token
	Type    LEX_TYPE

	F *File
}

func NewContext(F *File) *Context {
	return &Context{
		Start:   0,
		Current: 0,
		Line:    1,
		Tokens:  make([]*Token, 0),
		F:       F,
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
	Type    TOKEN_TYPE
	Lexeme  string
	Literal interface{}
	Line    int32
}

func NewToken(tp TOKEN_TYPE, lexeme string, literal interface{}, line int32) *Token {
	return &Token{
		Type:    tp,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
