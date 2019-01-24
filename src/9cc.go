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

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	// Lexering Loop
	// rp is ReadPointer
	for rp := 0; rp < len(l); rp++ {
		c := l[rp]

		if rp == 0 {
			s, n := readInt(l[rp:])
			rp += (n-1)
			fmt.Printf("  mov rax, %s\n", s)
			continue
		}

		if c == '+' {
			rp++
			s, n := readInt(l[rp:])
			rp += (n-1)
			fmt.Printf("  add rax, %s\n", s)
			continue
		}

		if c == '-' {
			rp++
			s, n := readInt(l[rp:])
			rp += (n-1)
			fmt.Printf("  sub rax, %s\n", s)
			continue
		}

		fmt.Fprintf(os.Stderr, "Unexpected character error: '%c'\n", c)
		os.Exit(1)
	}

	fmt.Println("  ret")
	return
}

// return converted int, len
func readInt(s string) (string, int) {
	p := 0
	for isDigit(s[p]){
		p++
		if p >= len(s){
			break
		}
	}
	return s[0:p], len(s[0:p])
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}