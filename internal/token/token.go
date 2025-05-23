package token

import "fmt"

type TokenType int

const (
	IDENTIFIER TokenType = iota
	REGISTER
	IMMEDIATE
	LABEL
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
	Col     int
}

var tokenTypeStrings = map[TokenType]string{
	IDENTIFIER: "IDENTIFIER",
	REGISTER:   "REGISTER",
	IMMEDIATE:  "IMMEDIATE",
	LABEL:      "LABEL",
}

func (tt TokenType) String() string {
	if s, ok := tokenTypeStrings[tt]; ok {
		return s
	}
	return fmt.Sprintf("Unknown token type(%d)", tt)
}
