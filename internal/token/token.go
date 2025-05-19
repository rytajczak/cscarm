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
	Column  int
}

func (t *Token) String() string {
	litStr := ""
	if t.Literal != nil {
		litStr = fmt.Sprintf(" (Lit: %v)", t.Literal)
	}
	return fmt.Sprintf("Type: %s, Lexeme: '%s'%s, Line: %d, Col: %d",
		t.Type.String(), t.Lexeme, litStr, t.Line, t.Column)
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
