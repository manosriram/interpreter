package main

import (
	"fmt"
	"interpreter/pkg/data"
	"interpreter/pkg/file"
	"log"
)

func main() {
	f, err := file.ReadFile("source")
	if err != nil {
		log.Fatal(err)
	}
	i_ctx := data.New_IContext(0, 0, 1)
	log.Println(f, i_ctx)
	// fmt.Println(string(f.D))

	for i := 0; i < int(f.Size); i++ {
		t := string(f.D[i])
		fmt.Println(t)
	}
}
