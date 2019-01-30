package main

import (
	"fmt"
	"os"
)

type Node struct {
	Ty  TokenType
	Lhs *Node
	Rhs *Node
	Val int
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
