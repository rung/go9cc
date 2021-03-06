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

const (
	TK_NUM    = "INT"
	TK_EOF    = "EOF"
	TK_PLUS   = "+"
	TK_MINUS  = "-"
	TK_MUL    = "*"
	TK_DIV    = "/"
	TK_OP     = "("
	TK_CP     = ")"
	TK_IDENT  = "IDENT"
	TK_ASSIGN = "="
	TK_EQ     = "=="
	TK_NOTEQ  = "!="
	TK_CALL   = "CALL"
	TK_FUNC   = "FUNC"
	TK_COMMA  = ","
	TK_LBRACE = "{"
	TK_RBRACE = "}"
	TK_RETURN = "RETURN"
)

// トークナイズした結果のトークン列はこのスライスに保存する
var tokens []Token
var pos int

func tokenize(l string) {
	tokens = []Token{}
	// Tokenize Loop
	// rp is ReadPointer
	for rp := 0; rp < len(l); {
		c := l[rp]

		if isSpace(l[rp]) {
			rp++
			continue
		}

		if c == '+' || c == '-' || c == '*' || c == '/' || c == '(' || c == ')' || c == ';' || c == ',' || c == '{' || c == '}' {
			t := Token{
				Ty:    TokenType(c),
				Input: l,
			}
			tokens = append(tokens, t)
			rp++
			continue
		}

		if c == '=' {
			// "==" op
			if l[rp+1] == '=' {
				t := Token{
					Ty:    TokenType("=="),
					Input: l[rp:],
				}
				tokens = append(tokens, t)
				rp += 2
				continue
			} else {
				// "=" op
				t := Token{
					Ty:    TokenType("="),
					Input: l[rp:],
				}
				tokens = append(tokens, t)
				rp++
				continue
			}
		}

		// "!="
		if c == '!' {
			if l[rp+1] == '=' {
				t := Token{
					Ty:    TokenType("!="),
					Input: l[rp:],
				}
				tokens = append(tokens, t)
				rp += 2
				continue
			} else {
				fmt.Fprintf(os.Stderr, "Can't tokenize '%s'\n", l[rp:])
				os.Exit(1)
			}
		}

		if isDigit(c) {
			num, n := readInt(l[rp:])
			t := Token{
				Ty:    TK_NUM,
				Input: l,
				Val:   num,
			}
			tokens = append(tokens, t)
			rp += n
			continue
		}

		if 'a' <= c && c <= 'z' {
			str, n := readString(l[rp:])
			t := Token{}
			if str == "return" {
				t = Token{
					Ty:    TK_RETURN,
					Input: str,
				}
			} else {
				t = Token{
					Ty:    TK_IDENT,
					Input: str,
				}
			}

			tokens = append(tokens, t)
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

func readString(s string) (string, int) {
	p := 0
	for 'a' <= s[p] && s[p] <= 'z' {
		p++
		if p >= len(s) {
			break
		}
	}
	str := s[0:p]
	return str, len(str)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
