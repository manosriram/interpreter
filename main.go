package main

import (
	"flag"
	"interpreter/src"
	"log"
)

type Flags struct {
	IsInteractive bool
}

func get_flags() *Flags {
	var interactive bool
	flag.BoolVar(&interactive, "i", false, "open interactive shell")
	flag.Parse()

	return &Flags{
		IsInteractive: interactive,
	}
}

func main() {
	flags := get_flags()
	if flags.IsInteractive {
		src.StartRepl()
		return
	}

	f, err := src.ReadFile("source")
	if err != nil {
		log.Fatal(err)
	}

	l := src.NewLexer(string(f.D))

	for i := 0; i < int(f.Size)-1; i++ {
		// fmt.Printf("-> %d ->", ctx.Start)
		l.Start = l.Current
		ok := l.Match()
		if !ok {
			log.Fatal("error compiling source file")
		}
	}
	l.AddToken(src.EOF)

	// l.PrintTokens()

	p := src.NewParser(l)
	program := p.ParseProgram()
	program.PrintStatements()
	// fmt.Println(program.String())
}
