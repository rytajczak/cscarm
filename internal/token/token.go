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
	MNEMONIC
	REGISTER
	IMMEDIATE
	LABEL
	IDENT
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
	EOF:     "EOF",
	NEWLINE: "\\n",
	COMMENT: "COMMENT",

	MNEMONIC:  "MNEMONIC",
	REGISTER:  "REGISTER",
	IMMEDIATE: "IMMEDIATE",
	LABEL:     "LABEL",
	IDENT:     "IDENT",

	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",

	RBRACK: "]",
	RBRACE: "}",
	COLON:  ":",
}

func (tok Token) String() string {
	if tok.Literal == "" {
		return fmt.Sprintf("type(%s) line: %d, col: %d", tokens[tok.Type], tok.Line, tok.Col)
	}
	return fmt.Sprintf("type(%s) literal(%s) line: %d, col: %d", tokens[tok.Type], tok.Literal, tok.Line, tok.Col)
}
