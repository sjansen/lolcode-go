package main

import (
	"fmt"
	"os"

	"github.com/sjansen/lolcode-go/internal/ast"
)

func main() {
	filename := os.Args[1]

	parser := &ast.Parser{Trace: true}
	program, err := parser.ParseFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	env := &ast.Environment{
		Output: os.Stderr,
	}

	program.Execute(env)
}
