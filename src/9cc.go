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
	n := add()

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	// generate assembly to read AST.
	gen(n)

	fmt.Println("  pop rax")
	fmt.Println("  ret")

	return
}
