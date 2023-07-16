package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Repl struct {
	Line string
	Ctx  *Context
}

func NewRepl(line string, Ctx *Context) *Repl {
	return &Repl{
		Line: line,
		Ctx:  Ctx,
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

		repl := NewRepl(line, &Context{Type: "REPL", F: &File{D: []byte(line), Size: int64(len(line))}})

		for i := 0; i < len(line); i++ {
			repl.Ctx.Start = repl.Ctx.Current
			ok := Match(repl.Ctx)
			if !ok {
				log.Fatal(fmt.Sprintf("error at %d\n", repl.Ctx.Line))
			}
		}
		AddToken(EOF, nil, repl.Ctx)

		PrintTokens(repl.Ctx.Tokens)
	}
}
