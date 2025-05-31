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

	MNEMONIC
	REGISTER
	IMMEDIATE
	LABEL
	IDENT

	LBRACK // [
	LBRACE // {
	RBRACK // ]
	RBRACE // }
	EXCLAM // !
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
}

var TokenMap = [...]string{
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
	RBRACK: "]",
	RBRACE: "}",
	EXCLAM: "!",
}

func (tok Token) String() string {
	if tok.Literal == "" {
		return fmt.Sprintf("type(%s) line: %d, col: %d", TokenMap[tok.Type], tok.Line, tok.Col)
	}
	return fmt.Sprintf("type(%s) literal(%s) line: %d, col: %d", TokenMap[tok.Type], tok.Literal, tok.Line, tok.Col)
}
