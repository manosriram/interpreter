package src

import "fmt"

func advance(ctx *Context) {
	ctx.Current++
}

func peek(ctx *Context) string {
	return string(ctx.F.D[ctx.Current])
}

func peek_next(ctx *Context) string {
	return string(ctx.F.D[ctx.Current+1])
}

func is_end(ctx *Context) bool {
	return int64(ctx.Current) >= ctx.F.Size-1
}

func is_alpha(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || c == "_"
}

func is_digit(c string) bool {
	return c >= "0" && c <= "9"
}

func typeof(c interface{}) string {
	return fmt.Sprintf("%T", c)
}

func is_alpha_numeric(c string) bool {
	return is_alpha(c) || is_digit(c)
}

func PrintTokens(tokens []*Token) {
	for _, x := range tokens {
		fmt.Printf("%s %s\n", x.Lexeme, x.Type)
	}
}
