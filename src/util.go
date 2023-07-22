package src

import "fmt"

func (l *Lexer) Advance() {
	l.Current++
}

func (l *Lexer) peek() string {
	return string(l.Input[l.Current])
}

func (l *Lexer) peek_next() string {
	return string(l.Input[l.Current+1])
}

func (l *Lexer) is_end() bool {
	return l.Current >= int32(len(l.Input))
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

func (l *Lexer) PrintTokens() {
	for _, x := range l.Tokens {
		fmt.Printf("%s %v\n", x.Type, x.Value)
	}
}
