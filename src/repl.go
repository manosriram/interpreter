package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Repl struct {
	Line  string
	Lexer *Lexer
}

func NewRepl(line string, L *Lexer) *Repl {
	return &Repl{
		Line:  line,
		Lexer: L,
	}
}

func StartRepl() {
	for {
		fmt.Printf(">> ")
		var line string

		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		repl := NewRepl(line, &Lexer{Input: line})

		for i := 0; i < len(line); i++ {
			repl.Lexer.Start = repl.Lexer.Current
			ok := repl.Lexer.Match()
			if !ok {
				log.Fatal("error parsing line")
				// log.Fatal(fmt.Sprintf("error at %d\n", repl.Ctx.Line))
			}
		}
		repl.Lexer.AddToken(EOF)

		repl.Lexer.PrintTokens()
	}
}
