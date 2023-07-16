package src

import (
	"strconv"
)

func AddToken(token_type TOKEN_TYPE, literal interface{}, ctx *Context) bool {
	d := string(ctx.F.D)
	lt := literal

	var lexeme string
	if token_type == STRING {
		lexeme = d[ctx.Start+1 : ctx.Current]
	} else {
		lexeme = d[ctx.Start:ctx.Current]
	}

	t := NewToken(token_type, lexeme, lt, ctx.Line)

	ctx.Tokens = append(ctx.Tokens, t)
	return true
}

func identifier(c string, ctx *Context) {
	for !is_end(ctx) && is_alpha_numeric(peek(ctx)) {
		advance(ctx)
	}

	lexeme := string(ctx.F.D[ctx.Start:ctx.Current])
	v, ok := Keywords[lexeme]

	if ok {
		AddToken(v, nil, ctx)
	} else {
		AddToken(IDENT, nil, ctx)
	}
}

func Match(ctx *Context) bool {
	if is_end(ctx) {
		return true
	}

	char_c := string(ctx.F.D[ctx.Current])
	advance(ctx)

	switch char_c {
	case "\n":
		ctx.Line++
		return true
	case " ", "\r", "\t":
		return true
	case "(":
		AddToken(LEFT_PAREN, nil, ctx)
		return true
	case ")":
		AddToken(RIGHT_PAREN, nil, ctx)
		return true
	case "{":
		AddToken(LEFT_BRACE, nil, ctx)
		return true
	case "}":
		AddToken(RIGHT_BRACE, nil, ctx)
		return true
	case ",":
		AddToken(COMMA, nil, ctx)
		return true
	case ".":
		AddToken(DOT, nil, ctx)
		return true
	case "-":
		AddToken(MINUS, nil, ctx)
		return true
	case "+":
		AddToken(PLUS, nil, ctx)
		return true
	case ";":
		AddToken(SEMICOLON, nil, ctx)
		return true
	case "*":
		AddToken(STAR, nil, ctx)
		return true
	case "=":
		if peek(ctx) == "=" {
			AddToken(EQUAL_EQUAL, nil, ctx)
			advance(ctx)
			return true
		} else {
			AddToken(EQUAL, nil, ctx)
			return true
		}
	case ">":
		if peek(ctx) == "=" {
			AddToken(GREATER_EQUAL, nil, ctx)
			advance(ctx)
			return true
		} else {
			AddToken(GREATER, nil, ctx)
			return true
		}
	case "<":
		if peek(ctx) == "=" {
			AddToken(LESS_EQUAL, nil, ctx)
			advance(ctx)
			return true
		} else {
			AddToken(LESS, nil, ctx)
			return true
		}
	case "/":
		if peek(ctx) == "/" {
			for !is_end(ctx) && peek(ctx) != "\n" {
				advance(ctx)
			}
			advance(ctx)
			ctx.Line++
			return true
		} else {
			AddToken(SLASH, nil, ctx)
			return true
		}
	case "\"":
		for !is_end(ctx) && peek(ctx) != "\"" {
			advance(ctx)
		}
		if is_end(ctx) {
			return false
		}

		// lit := g.F.D[i_ctx.Start:i_ctx.Current]
		AddToken(STRING, ctx.F.D[ctx.Start:ctx.Current], ctx)
		advance(ctx)
		return true

	default:
		if is_digit(char_c) {
			for !is_end(ctx) && is_digit(peek(ctx)) {
				advance(ctx)
			}
			int_value, _ := strconv.ParseInt(string(ctx.F.D[ctx.Start:ctx.Current]), 0, 32)
			AddToken(INT, int_value, ctx)
			return true
		} else if is_alpha(char_c) {
			identifier(char_c, ctx)
			return true
		}

		return false
	}
}
