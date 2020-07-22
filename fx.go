package main

import (
	"bufio"
	"fmt"
	"fxlex"
	"fxparse"
	"os"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Fprintf(os.Stderr, "%s requires one argument\n", os.Args[0])
		os.Exit(1)
	}

	filename := args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, "opening file:", err)
		return
	}

	reader := bufio.NewReader(file)
	lexer, err := fxlex.NewLexer(reader, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lexer instantiation failed\n")
		os.Exit(1)
	}

	parser, err := fxparse.NewParser(lexer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parser instantiation failed\n")
		os.Exit(1)
	}

	parser.Parse()

	file.Close()
}
