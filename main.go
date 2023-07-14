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
	core.AddToken(data.END_OF_FILE, nil, global_ctx)

	for _, x := range i_ctx.Tokens {
		fmt.Println(x.Tp)
	}

	p := core.NewParser(global_ctx.Ctx.Tokens, int32(len(global_ctx.Ctx.Tokens)))
	expr := p.Parse()
	fmt.Println(expr)

	// b := core.Binary{
	// Left: &core.Unary{
	// Operator: data.Token{
	// Tp:      data.MINUS,
	// Lexeme:  "-",
	// Literal: nil,
	// Line:    1,
	// },
	// Right: &core.Literal{
	// Value: "123",
	// },
	// },
	// Operator: data.Token{
	// Tp:      data.STAR,
	// Lexeme:  "*",
	// Literal: nil,
	// Line:    1,
	// },
	// Right: &core.Grouping{
	// Expression: &core.Literal{
	// Value: 45.67,
	// },
	// },
	// }

	// // fmt.Println(l.Accept())

	// fmt.Println(core.PrintExpr(expr))

	// b.VisitBinaryExpr(b)

	// for _, t := range global_ctx.Ctx.Tokens {
	// fmt.Printf("%d -> %s -> %d -> %v -> literal_type(%T)\n", t.Tp, t.Lexeme, t.Line, t.Literal, t.Literal)
	// }
}
