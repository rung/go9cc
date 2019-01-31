package main

import (
	"fmt"
	"os"
)

// varibles map
var maps map[string]int = make(map[string]int)
var offset int = 1

func gen(n *Node) {
	if n.Ty == TK_CALL {
		fmt.Printf("  call %s\n", n.Name)
		//fmt.Println("  push rax")
		fmt.Println("  push 0")
		return
	}

	if n.Ty == TK_NUM {
		fmt.Printf("  push %d\n", n.Val)
		return
	}

	if n.Ty == TK_IDENT {
		genLval(n)
		fmt.Println("  pop rax")
		fmt.Println("  mov rax, [rax]")
		fmt.Println("  push rax")
		return
	}

	if n.Ty == "=" {
		genLval(n.Lhs)
		gen(n.Rhs)

		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

	gen(n.Lhs)
	gen(n.Rhs)

	fmt.Println("  pop rdi")
	fmt.Println("  pop rax")

	switch n.Ty {
	case "+":
		fmt.Println("  add rax, rdi")
	case "-":
		fmt.Println("  sub rax, rdi")
	case "*":
		fmt.Println("  mul rdi")
	case "/":
		fmt.Println("  mov rdx, 0")
		fmt.Println("  div rdi")
	case "==":
		fmt.Println("  cmp rdi, rax")
		fmt.Println("  sete al")
		fmt.Println("  movzb rax, al")
	case "!=":
		fmt.Println("  cmp rdi, rax")
		fmt.Println("  setne al")
		fmt.Println("  movzb rax, al")
	}

	fmt.Println("  push rax")

}

func genLval(node *Node) {
	if node.Ty != TK_IDENT {
		fmt.Fprintln(os.Stderr, "代入の左辺値が変数ではありません")
	}

	o := getMap(node.Name) * 8
	fmt.Println("  mov rax, rbp")
	fmt.Printf("  sub rax, %d\n", o)
	fmt.Println("  push rax")
}

func getMap(v string) int {
	i, ok := maps[v]
	if ok == false {
		maps[v] = offset
		offset++
		return maps[v]
	}
	return i
}
