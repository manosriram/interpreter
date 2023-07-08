package main

import (
	"fmt"
	"interpreter/pkg/core"
	"interpreter/pkg/data"
	"interpreter/pkg/file"
	"log"
)

func main() {
	f, err := file.ReadFile("source")
	if err != nil {
		log.Fatal(err)
	}
	// f.D = append(f.D, byte("a"))

	i_ctx := data.New_IContext(0, 0, 1)
	// fmt.Println(string(f.D))

	global_ctx := &data.GlobalCtx{
		F:   f,
		Ctx: i_ctx,
	}

	for i := 0; i < int(f.Size)-1; i++ {
		i_ctx.Start = i_ctx.Current
		// fmt.Printf("-> %d ->", i_ctx.Start)
		ok := core.Match(global_ctx)
		if !ok {
			log.Fatal("error compiling source file")
		}
	}

	for _, t := range global_ctx.Ctx.Tokens {
		fmt.Printf("%d -> %s -> %d -> %v -> literal_type(%T)\n", t.Tp, t.Lexeme, t.Line, t.Literal, t.Literal)
	}
}
