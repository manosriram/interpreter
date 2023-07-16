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

	ctx := src.NewContext(f)

	for i := 0; i < int(f.Size)-1; i++ {
		// fmt.Printf("-> %d ->", ctx.Start)
		ctx.Start = ctx.Current
		ok := src.Match(ctx)
		if !ok {
			log.Fatal("error compiling source file")
		}
	}
	src.AddToken(src.EOF, nil, ctx)

	src.PrintTokens(ctx.Tokens)
}
