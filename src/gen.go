package main

import "fmt"

func gen(n *Node) {
	if n.Ty == TK_NUM {
		fmt.Printf("  push %d\n", n.Val)
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
