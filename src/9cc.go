package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Errorf("the number of arguments is wrong.")
		return
	}

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	fmt.Printf("  mov rax, %s\n", os.Args[1])
	fmt.Println("  ret")
	return
}
