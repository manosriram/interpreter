package src

func (l *Lexer) PeekNextToken() *Token {
	if l.Tokens[l.TokenPosition].Type == EOF {
		return nil
	}
	next := l.Tokens[l.TokenPosition+1]
	l.TokenPosition++
	return next
}

func (l *Lexer) PeekToken() *Token {
	return l.Tokens[l.TokenPosition]
}

func (l *Lexer) AddToken(token_type TOKEN_TYPE) bool {
	var lexeme string
	if token_type == STRING {
		lexeme = l.Input[l.Start+1 : l.Current]
	} else {
		lexeme = l.Input[l.Start:l.Current]
	}

	t := NewToken(token_type, lexeme)

	l.Tokens = append(l.Tokens, t)
	return true
}

func (l *Lexer) identifier(c string) {
	for !l.is_end() && is_alpha_numeric(l.peek()) {
		l.Advance()
	}

	lexeme := string(l.Input[l.Start:l.Current])
	v, ok := Keywords[lexeme]

	if ok {
		l.AddToken(v)
	} else {
		l.AddToken(IDENT)
	}
}

func (l *Lexer) Match() bool {
	if l.is_end() {
		return true
	}

	l.Ch = byte(l.Input[l.Current])
	char_c := string(l.Ch)
	l.Advance()

	switch char_c {
	case "\n":
		// ctx.Line++
		return true
	case " ", "\r", "\t":
		return true
	case "(":
		l.AddToken(LEFT_PAREN)
		return true
	case ")":
		l.AddToken(RIGHT_PAREN)
		return true
	case "{":
		l.AddToken(LEFT_BRACE)
		return true
	case "}":
		l.AddToken(RIGHT_BRACE)
		return true
	case ",":
		l.AddToken(COMMA)
		return true
	case ".":
		l.AddToken(DOT)
		return true
	case "-":
		l.AddToken(MINUS)
		return true
	case "+":
		l.AddToken(PLUS)
		return true
	case ";":
		l.AddToken(SEMICOLON)
		return true
	case "*":
		l.AddToken(STAR)
		return true
	case "=":
		if l.peek() == "=" {
			l.AddToken(EQUAL_EQUAL)
			l.Advance()
			return true
		} else {
			l.AddToken(EQUAL)
			return true
		}
	case ">":
		if l.peek() == "=" {
			l.AddToken(GREATER_EQUAL)
			l.Advance()
			return true
		} else {
			l.AddToken(GREATER)
			return true
		}
	case "<":
		if l.peek() == "=" {
			l.AddToken(LESS_EQUAL)
			l.Advance()
			return true
		} else {
			l.AddToken(LESS)
			return true
		}
	case "/":
		if l.peek() == "/" {
			for !l.is_end() && l.peek() != "\n" {
				l.Advance()
			}
			l.Advance()
			// ctx.Line++
			return true
		} else {
			l.AddToken(SLASH)
			return true
		}
	case "\"":
		for !l.is_end() && l.peek() != "\"" {
			l.Advance()
		}
		if l.is_end() {
			return false
		}

		// lit := g.F.D[i_ctx.Start:i_ctx.Current]
		l.AddToken(STRING)
		l.Advance()
		return true

	default:
		// if is_digit(char_c) {
		// for !l.is_end() && is_digit(l.peek()) {
		// l.Advance()
		// }
		// int_value, _ := strconv.ParseInt(string(ctx.F.D[ctx.Start:ctx.Current]), 0, 32)
		// l.AddToken(INT)
		// return true
		// } else if is_alpha(char_c) {
		l.identifier(char_c)
		return true
		// }

		return false
	}
}
