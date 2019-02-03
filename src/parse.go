package main

import (
	"fmt"
	"os"
)

type Node struct {
	Ty   TokenType
	Lhs  *Node
	Rhs  *Node
	Val  int     // // used when Ty is TK_INT
	Name string  // used when Ty is TK_IDENT
	Args []*Node // used when Ty is TK_CALL
	Ret  *Node   // used when Ty is TK_RETURN
}

type Func struct {
	name string
	code []*Node
}

var code []*Node
var funcs map[string]*Func = make(map[string]*Func)

func program() {

	for tokens[pos].Ty != TK_EOF {
		toplevel()
	}
}

// function定義
func toplevel() {
	name := ident()
	funcs[name.Name] = &Func{name: name.Name}

	if !consume("(") {
		fmt.Fprintf(os.Stderr, "This token is not '(': %s", tokens[pos].Input)
		os.Exit(1)
	}
	if !consume(")") {
		fmt.Fprintf(os.Stderr, "This token is not '(': %s", tokens[pos].Input)
		os.Exit(1)
	}

	if consume("{") {
		funcs[name.Name].code = []*Node{}
		for !consume("}") {
			funcs[name.Name].code = append(funcs[name.Name].code, stmt())
		}
	}
}

func stmt() *Node {
	if consume(TK_RETURN) {
		n := assign()
		n = newNodeReturn(n)

		if !consume(";") {
			fmt.Fprintf(os.Stderr, "This token is not ';': %s", tokens[pos].Input)
		}

		return n
	}

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
	node := funcCall()

	for {
		if consume("==") {
			node = newNode("==", node, funcCall())
		} else if consume("!=") {
			node = newNode("!=", node, funcCall())
		} else {
			return node
		}
	}
}

func funcCall() *Node {
	node := add()

	if consume("(") {
		node = newNodeCall(node.Name)
		if consume(")") {
			return node
		}

		node.Args = append(node.Args, add())

		for consume(",") {
			node.Args = append(node.Args, add())
		}

		if !consume(")") {
			fmt.Fprintf(os.Stderr, "There isn't a closing parenthesis: %s\n",
				tokens[pos].Input)
			os.Exit(1)
		}
	}
	return node
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

func ident() *Node {
	if tokens[pos].Ty == TK_IDENT {
		n := newNodeIdent(tokens[pos].Input)
		pos++
		return n
	}

	fmt.Fprintf(os.Stderr, "This token isn't a ident: %s",
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

func newNodeReturn(node *Node) *Node {
	n := &Node{
		Ty:  TK_RETURN,
		Ret: node,
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

func newNodeCall(name string) *Node {
	n := &Node{
		Ty:   TK_CALL,
		Name: name,
		Args: []*Node{},
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
