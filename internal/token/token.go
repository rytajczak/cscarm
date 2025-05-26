package token

import (
	"fmt"
)

type TokenType int

const (
	ILLEGAL TokenType = iota
	NEWLINE
	EOF
	COMMENT

	literal_beg
	IDENT
	REGISTER
	IMMEDIATE
	literal_end

	delimiters_beg
	LBRACK // [
	LBRACE // {
	COMMA  // ,

	RBRACK // ]
	RBRACE // }
	COLON  // :
	delimiters_end
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
}

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	NEWLINE: "NEWLINE",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:     "IDENT",
	REGISTER:  "REGISTER",
	IMMEDIATE: "IMMEDIATE",

	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",

	RBRACK: "]",
	RBRACE: "}",
	COLON:  ":",
}

func (tok Token) String() string {
	return fmt.Sprintf("type: %s, literal: %s, line: %d, col: %d", tokens[tok.Type], tok.Literal, tok.Line, tok.Col)
}
