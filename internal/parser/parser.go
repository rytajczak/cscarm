package parser

import (
	"cscasm/internal/token"
	"fmt"
	"slices"
	"strings"
)

type Parser struct {
	ST         map[string]int
	currentIns int
}

type Instruction interface {
	Encoding() uint32
}

// Data processing opcodes
var dp = []string{"SUB", "ADD", "ORR", "SUBS"}

// Single data transfer opcodes
var sdt = []string{"LDR", "STR"}

// Branch and branch with link opcodes
var bbl = []string{"B", "BL", "BPL"}

func (p *Parser) ParseInstructionFromTokens(toks []*token.Token) (Instruction, error) {
	p.currentIns++

	mnemonic := strings.ToUpper(toks[0].Lexeme)

	switch true {
	// (I still dont know what this is called)
	case mnemonic == "MOVW":
		return p.newMovwIns(toks)
	case mnemonic == "MOVT":
		return p.newMovtIns(toks)
	// Data processing
	case slices.Contains(dp, mnemonic):
		return p.newDpIns(mnemonic, toks)
	// Single data transfer
	case slices.Contains(sdt, mnemonic):
		return p.newSdtIns(mnemonic, toks)
	// Branching
	case slices.Contains(bbl, mnemonic):
		return p.newBblIns(mnemonic, toks)
	case mnemonic == "BX":
		return p.newBxIns(toks)
	default:
		return nil, fmt.Errorf("failed to identify mnemonic")
	}
}

func NewParser(st map[string]int) *Parser {
	return &Parser{
		ST:         st,
		currentIns: 0,
	}
}

// type ParserError struct {
// 	Line int
// 	Col  int
// 	Err  error
// }

// func (p *ParserError) Error() string {
// 	return p.Err.Error()
// }

// func newParserError(tok *token.Token, format string, args ...any) *ParserError {
// 	line, col := 0, 0
// 	if tok != nil {
// 		line = tok.Line
// 		col = tok.Col
// 	}
// 	return &ParserError{
// 		Line: line,
// 		Col:  col,
// 		Err:  fmt.Errorf(format, args...),
// 	}
// }
