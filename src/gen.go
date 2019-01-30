package main

import (
	"fmt"
	"os"
)

func gen(n *Node) {
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
	}

	fmt.Println("  push rax")

}

func genLval(node *Node) {
	if node.Ty != TK_IDENT {
		fmt.Fprintln(os.Stderr, "代入の左辺値が変数ではありません")
	}

	// stack - charをしてアドレスをpush
	offset := ('z' - byte(node.Name[0]) + 1) * 8
	fmt.Println("  mov rax, rbp")
	fmt.Printf("  sub rax, %d\n", offset)
	fmt.Println("  push rax")
}
