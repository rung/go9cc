package main

import (
	"fmt"
	"os"
)

// varibles map
var maps map[string]int = make(map[string]int)
var offset int = 1

func startGen(f *Func) {
	fmt.Println(".text")
	fmt.Printf(".global %s\n", f.name)
	fmt.Printf("%s:\n", f.name)

	// get an area of variables
	fmt.Println("  push rbp")
	fmt.Println("  mov rbp, rsp")
	fmt.Println("  sub rsp, 208")

	for _, c := range f.code {
		gen_stmt(c)
	}

	// ToDo: return文実装済みなので下記は不要かもしれない
	fmt.Println("  pop rax")
	// 最後の式の結果がRAXに残っているのでそれが返り値になる
	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")

}

func gen_stmt(n *Node) {
	if n.Ty == TK_CALL {
		for r := range n.Args {
			i := len(n.Args) - r - 1
			gen_stmt(n.Args[i])
			if i == 0 {
				fmt.Println("  pop rdi")
			} else if i == 1 {
				fmt.Println("  pop rsi")
			} else if i == 2 {
				fmt.Println("  pop rdx")
			} else if i == 3 {
				fmt.Println("  pop rcx")
			} else if i == 4 {
				fmt.Println("  pop r8")
			} else if i == 5 {
				fmt.Println("  pop r9")
			}
		}
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

	if n.Ty == TK_RETURN {
		gen_stmt(n.Ret)
		fmt.Println("  pop rax")
		fmt.Println("  mov rsp, rbp")
		fmt.Println("  pop rbp")
		fmt.Println("  ret")
		return
	}

	if n.Ty == "=" {
		genLval(n.Lhs)
		gen_stmt(n.Rhs)

		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

	gen_stmt(n.Lhs)
	gen_stmt(n.Rhs)

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
