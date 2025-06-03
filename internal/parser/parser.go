package parser

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rytajczak/cscarm/internal/lexer"
	"github.com/rytajczak/cscarm/internal/token"
)

type Parser struct {
	file         string
	lexer        *lexer.Lexer
	currentToken *token.Token
	symbolTable  map[string]int
	currentLine  int
}

func NewParser(file *os.File) (*Parser, error) {
	lex := lexer.NewLexer(file)
	p := &Parser{
		file:        file.Name(),
		lexer:       lex,
		symbolTable: make(map[string]int),
	}
	p.nextToken()

	return p, nil
}

func (p *Parser) nextToken() {
	p.currentToken = p.lexer.NextToken()
}

func (p *Parser) Parse() ([]uint32, []error) {
	for p.currentToken.Type != token.EOF {
		switch p.currentToken.Type {
		case token.LABEL:
			p.symbolTable[p.currentToken.Literal] = p.currentLine
			p.nextToken()
		case token.MNEMONIC:
			p.currentLine++
			p.nextToken()
		default:
			p.nextToken()
		}
	}
	p.lexer.Reset()
	p.nextToken()
	p.currentLine = 0

	var instructions []uint32
	var errors []error
	for p.currentToken.Type != token.EOF {
		if p.currentToken.Type != token.MNEMONIC {
			p.nextToken()
			continue
		}
		ins, err := p.encodeInstruction(p.currentToken.Literal)
		if err != nil {
			errors = append(errors, p.formatError(err))
		}
		instructions = append(instructions, ins)
		p.currentLine++
		p.nextToken()
	}
	return instructions, errors
}

func (p *Parser) encodeInstruction(mnemonic string) (uint32, error) {
	p.nextToken()
	m := strings.ToUpper(mnemonic)
	switch true {
	case m == "MOVW":
		return p.encodeMovwINS()
	case m == "MOVT":
		return p.encodeMovtINS()
	case m == "BX":
		return p.encodeBranchExchangeINS()
	case slices.Contains([]string{"B", "BL", "BPL", "BGE"}, m):
		return p.encodeBranchINS(m)
	case slices.Contains([]string{"ADD", "SUB", "SUBS", "ORR"}, m):
		return p.encodeDataProcessingINS(m)
	case slices.Contains([]string{"LDR", "STR"}, m):
		return p.encodeSingleDataTransferINS(m)
	case slices.Contains([]string{"STMEA", "LDMEA"}, m):
		return p.encodeBlockDataTransferINS(m)
	default:
		return 0, nil
	}
}

func parseRegister(reg string) uint32 {
	var r uint64
	switch reg {
	case "SP":
		r = 13
	case "LR":
		r = 14
	case "PC":
		r = 15
	default:
		r, _ = strconv.ParseUint(reg[1:], 0, 32)
	}
	return uint32(r)
}

func (p *Parser) consumeRegister() (uint32, error) {
	if p.currentToken.Type != token.REGISTER {
		return 0, fmt.Errorf("expected register, found %s", strings.ToLower(token.TokenMap[p.currentToken.Type]))
	}
	reg := parseRegister(strings.ToUpper(p.currentToken.Literal))
	p.nextToken()
	return uint32(reg), nil
}

func (p *Parser) consumeImmediate() (uint32, bool, error) {
	if p.currentToken.Type != token.IMMEDIATE {
		return 0, false, fmt.Errorf("expected immediate, found %s", strings.ToLower(token.TokenMap[p.currentToken.Type]))
	}

	var isNegative bool = false
	if p.currentToken.Literal[0] == '-' {
		isNegative = true
	}

	imm, _ := strconv.ParseInt(p.currentToken.Literal, 0, 32)
	p.nextToken()
	return uint32(imm), isNegative, nil
}

func (p *Parser) consumeIdent() (string, error) {
	if p.currentToken.Type != token.IDENT {
		return "", fmt.Errorf("expected identifier, found %s", strings.ToLower(token.TokenMap[p.currentToken.Type]))
	}
	ident := p.currentToken.Literal
	p.nextToken()
	return ident, nil
}

func (p *Parser) formatError(err error) error {
	r := color.New(color.FgRed).SprintFunc()
	b := color.New(color.FgBlue).SprintFunc()

	reason := fmt.Sprintf("%s %s", r("error:"), err.Error())
	location := fmt.Sprintf("%s %s:%d:%d", b(" -->"), p.file, p.currentToken.Line, p.currentToken.Col)

	return fmt.Errorf("%s\n%s", reason, location)
}
