package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "The number of arguments is wrong.\n")
		return
	}

	l := os.Args[1]

	// tokenize and parse
	tokenize(l)
	program()

	fmt.Println(".intel_syntax noprefix")

	for _, v := range funcs {
		startGen(v)

	}
	return
}
