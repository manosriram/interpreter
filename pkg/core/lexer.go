package core

import (
	"fmt"
	"interpreter/pkg/data"
	"strconv"
)

func AddToken(tp data.TOKEN_TYPE, literal interface{}, g *data.GlobalCtx) bool {
	d := string(g.F.D)
	i_ctx := g.Ctx
	lt := literal

	var lexeme string
	if tp == data.STRING {
		lexeme = d[i_ctx.Start+1 : i_ctx.Current]
	} else {
		lexeme = d[i_ctx.Start:i_ctx.Current]
	}

	t := data.NewToken(tp, lexeme, lt, g.Ctx.Line)

	i_ctx.Tokens = append(i_ctx.Tokens, t)
	return true
}

func advance(i_ctx *data.I_Context) {
	i_ctx.Current++
}

func peek(g_ctx *data.GlobalCtx) string {
	return string(g_ctx.F.D[g_ctx.Ctx.Current])
}

func peek_next(g_ctx *data.GlobalCtx) string {
	return string(g_ctx.F.D[g_ctx.Ctx.Current+1])
}

func is_end(g_ctx *data.GlobalCtx) bool {
	return int64(g_ctx.Ctx.Current) >= g_ctx.F.Size-1
}

func is_alpha(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || c == "_"
}

func is_digit(c string) bool {
	return c >= "1" && c <= "9"
}

func typeof(c interface{}) string {
	return fmt.Sprintf("%T", c)
}

func is_alpha_numeric(c string) bool {
	return is_alpha(c) || is_digit(c)
}

func identifier(c string, g *data.GlobalCtx) {
	for !is_end(g) && is_alpha_numeric(peek(g)) {
		advance(g.Ctx)
	}

	lexeme := string(g.F.D[g.Ctx.Start:g.Ctx.Current])
	if data.ReservedTokenMap[lexeme] >= 22 && data.ReservedTokenMap[lexeme] <= 37 {
		AddToken(data.ReservedTokenMap[lexeme], nil, g)
	} else {
		AddToken(data.IDENTIFIER, nil, g)
	}
}

func Match(g *data.GlobalCtx) bool {
	if is_end(g) {
		return true
	}
	i_ctx := g.Ctx

	char_c := string(g.F.D[i_ctx.Current])
	advance(i_ctx)

	switch char_c {
	case "\n":
		i_ctx.Line++
		return true
	case " ", "\r", "\t":
		return true
	case "(":
		AddToken(data.LEFT_PAREN, nil, g)
		return true
	case ")":
		AddToken(data.RIGHT_PAREN, nil, g)
		return true
	case "{":
		AddToken(data.LEFT_BRACE, nil, g)
		return true
	case "}":
		AddToken(data.RIGHT_BRACE, nil, g)
		return true
	case ",":
		AddToken(data.COMMA, nil, g)
		return true
	case ".":
		AddToken(data.DOT, nil, g)
		return true
	case "-":
		AddToken(data.MINUS, nil, g)
		return true
	case "+":
		AddToken(data.PLUS, nil, g)
		return true
	case ";":
		AddToken(data.SEMICOLON, nil, g)
		return true
	case "*":
		AddToken(data.STAR, nil, g)
		return true
	case "=":
		if peek(g) == "=" {
			AddToken(data.EQUAL_EQUAL, nil, g)
			advance(i_ctx)
			return true
		} else {
			AddToken(data.EQUAL, nil, g)
			return true
		}
	case ">":
		if peek(g) == "=" {
			AddToken(data.GREATER_EQUAL, nil, g)
			advance(i_ctx)
			return true
		} else {
			AddToken(data.GREATER, nil, g)
			return true
		}
	case "<":
		if peek(g) == "=" {
			AddToken(data.LESS_EQUAL, nil, g)
			advance(i_ctx)
			return true
		} else {
			AddToken(data.LESS, nil, g)
			return true
		}
	case "/":
		if peek(g) == "/" {
			for !is_end(g) && peek(g) != "\n" {
				advance(i_ctx)
			}
			advance(i_ctx)
			g.Ctx.Line++
			return true
		} else {
			AddToken(data.SLASH, nil, g)
			return true
		}
	case "\"":
		for !is_end(g) && peek(g) != "\"" {
			advance(i_ctx)
		}
		if is_end(g) {
			return false
		}

		// lit := g.F.D[i_ctx.Start:i_ctx.Current]
		AddToken(data.STRING, g.F.D[i_ctx.Start:i_ctx.Current], g)
		advance(i_ctx)
		return true

	default:
		if is_digit(char_c) {
			for !is_end(g) && is_digit(peek(g)) {
				advance(i_ctx)
			}
			var is_double bool = false
			if peek(g) == "." {
				is_double = true
				advance(i_ctx)
			}
			for !is_end(g) && is_digit(peek(g)) {
				advance(i_ctx)
			}

			if is_double {
				double_value, _ := strconv.ParseFloat(string(g.F.D[i_ctx.Start:i_ctx.Current]), 0)
				AddToken(data.NUMBER, double_value, g)
			} else {
				int_value, _ := strconv.ParseInt(string(g.F.D[i_ctx.Start:i_ctx.Current]), 0, data.INT32)
				AddToken(data.NUMBER, int_value, g)
			}

			return true
		} else if is_alpha(char_c) {
			identifier(char_c, g)
			return true
		}

		return false
	}
}
