package main

import (
	"fmt"
	"os"
)

type Node struct {
	Ty   TokenType
	Lhs  *Node
	Rhs  *Node
	Val  int    // // used when Ty is TK_INT
	Name string // used when Ty is TK_IDENT
}

var code []*Node

func program() {
	code = []*Node{}

	for tokens[pos].Ty != TK_EOF {
		code = append(code, stmt())
	}
}
func stmt() *Node {
	n := assign()
	if !consume(";") {
		fmt.Fprintf(os.Stderr, "This token is not ';': %s", tokens[pos].Input)
	}
	return n
}

func assign() *Node {
	node := equality()

	for {
		if consume("=") {
			node = newNode("=", node, equality())
		} else {
			return node
		}
	}
}

func equality() *Node {
	node := add()

	for {
		if consume("==") {
			node = newNode("==", node, add())
		} else if consume("!=") {
			node = newNode("!=", node, add())
		} else {
			return node
		}
	}
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

	if tokens[pos].Ty == TK_IDENT {
		n := newNodeIdent(tokens[pos].Input)
		pos++
		return n
	}

	fmt.Fprintf(os.Stderr, "This token isn't a number or opening parenhesis: %s",
		tokens[pos].Input)
	os.Exit((1))

	return nil
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

func newNodeIdent(name string) *Node {
	n := &Node{
		Ty:   TK_IDENT,
		Name: name,
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
