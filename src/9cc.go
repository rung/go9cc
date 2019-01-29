package main

import (
	"fmt"
	"os"
	"strconv"
)

type TokenType string

type Token struct {
	Ty    TokenType // tokentype
	Val   int       // value if it's integer
	Input string    // token string(for an error message)
}

type Node struct {
	Ty  TokenType
	Lhs *Node
	Rhs *Node
	Val int
}

const (
	TK_NUM   = "INT"
	TK_EOF   = "EOF"
	TK_PLUS  = "+"
	TK_MINUS = "-"
	TK_MUL   = "*"
	TK_DIV   = "/"
	TK_OP    = "("
	TK_CP    = ")"
)

// トークナイズした結果のトークン列はこのスライスに保存する
var tokens []Token
var pos int

func tokenize(l string) {
	tokens = []Token{}
	i := 0
	// Tokenize Loop
	// rp is ReadPointer
	for rp := 0; rp < len(l); {
		c := l[rp]

		if isSpace(l[rp]) {
			rp++
			continue
		}

		if c == '+' || c == '-' || c == '*' || c == '/' || c == '(' || c == ')' {
			t := Token{
				Ty:    TokenType(c),
				Input: l,
			}
			tokens = append(tokens, t)
			i++
			rp++
			continue
		}

		if isDigit(c) {
			num, n := readInt(l[rp:])
			t := Token{
				Ty:    TK_NUM,
				Input: l,
				Val:   num,
			}
			tokens = append(tokens, t)
			i++
			rp += n
			continue
		}

		fmt.Fprintf(os.Stderr, "Can't tokenize '%s'\n", l[rp:])
		os.Exit(1)
	}

	tokens = append(tokens, Token{Ty: TK_EOF})
}

func error(i int) {
	fmt.Fprintf(os.Stderr, "Unexpected token error: %s\n", tokens[i].Input)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "The number of arguments is wrong.\n")
		return
	}

	l := os.Args[1]
	tokenize(l)
	// n := add()
	// fmt.Printf("%T", n)

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	if tokens[0].Ty != TK_NUM {
		error(0)
	}

	fmt.Printf("  mov rax, %d\n", tokens[0].Val)

	i := 1
	for tokens[i].Ty != TK_EOF {
		if tokens[i].Ty == TK_PLUS {
			i++
			if tokens[i].Ty != TK_NUM {
				error(i)
			}
			fmt.Printf("  add rax, %d\n", tokens[i].Val)
			i++
			continue
		}

		if tokens[i].Ty == "-" {
			i++
			if tokens[i].Ty != TK_NUM {
				error(i)
			}
			fmt.Printf("  sub rax, %d\n", tokens[i].Val)
			i++
			continue
		}
		error(i)
	}
	fmt.Printf("  ret\n")
	return
}

// return converted int, len
func readInt(s string) (int, int) {
	p := 0
	for isDigit(s[p]) {
		p++
		if p >= len(s) {
			break
		}
	}
	i, err := strconv.Atoi(s[0:p])
	if err != nil {
		panic(err)
	}
	return i, len(s[0:p])
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func newNode(ty TokenType, lhs *Node, rhs *Node) *Node {
	n := &Node{
		Ty:  ty,
		Lhs: lhs,
		Rhs: rhs,
	}
	return n
}

func newNodeNum(val int) *Node {
	n := &Node{
		Ty:  TK_NUM,
		Val: val,
	}
	return n
}

func consume(ty TokenType) bool {
	if tokens[pos].Ty != ty {
		return false
	}
	pos++
	return true
}

func add() *Node {
	node := mul()

	for {
		if consume("+") {
			node = newNode("+", node, mul())
		} else if consume("-") {
			node = newNode("-", node, mul())
		} else {
			return node
		}
	}
}

func mul() *Node {
	node := term()

	for {
		if consume("*") {
			node = newNode("*", node, term())
		} else if consume("/") {
			node = newNode("/", node, term())
		} else {
			return node
		}
	}
}

func term() *Node {

	if consume("(") {
		node := add()
		if !consume(")") {
			fmt.Fprintf(os.Stderr, "There isn't a closing parenthesis: %s",
				tokens[pos].Input)
			os.Exit(1)
		}
		return node
	}

	if tokens[pos].Ty == TK_NUM {
		n := newNodeNum(tokens[pos].Val)
		pos++
		return n
	}

	fmt.Fprintf(os.Stderr, "This token isn't a number or opening parenhesis: %s",
		tokens[pos].Input)
	os.Exit((1))

	return nil
}
