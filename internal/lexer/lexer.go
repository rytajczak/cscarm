package lexer

import (
	"cscasm/internal/token"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	lineCount int
}

var charsToIgnore = []byte{
	',', '(', ')', '[', ']', '.',
}

var registers = []string{
	"R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8",
	"R9", "R10", "R11", "R12", "R13", "R14", "R15",
	"SP", "LR", "FC",
}

func NewLexer() *Lexer {
	return &Lexer{
		lineCount: 0,
	}
}

func (l *Lexer) TokenizeLine(line string) ([]*token.Token, error) {
	l.lineCount++
	charPos := 0
	tokens := []*token.Token{}

	line = strings.Split(line, "@")[0]
	lineLen := len(line)

	if lineLen == 0 {
		return []*token.Token{}, nil
	}

	for charPos < lineLen {
		if unicode.IsSpace(rune(line[charPos])) {
			charPos++
			continue
		}

		var lexemeB strings.Builder
		startPos := charPos
		for charPos < lineLen && !unicode.IsSpace(rune(line[charPos])) {
			if !slices.Contains(charsToIgnore, line[charPos]) {
				lexemeB.WriteByte(line[charPos])
			}
			charPos++
		}
		lexeme := lexemeB.String()

		tok := &token.Token{
			Type:    token.IDENTIFIER,
			Lexeme:  lexeme,
			Literal: nil,
			Line:    l.lineCount,
			Col:     startPos + 1,
		}

		switch true {
		case strings.HasSuffix(lexeme, ":"):
			tok.Type = token.LABEL
			tok.Literal = strings.TrimSuffix(lexeme, ":")
		case slices.Contains(registers, strings.ToUpper(lexeme)):
			tok.Type = token.REGISTER
			tok.Literal = parseRegister(lexeme)
		case strings.HasPrefix(lexeme, "#"):
			tok.Type = token.IMMEDIATE
			tok.Literal = parseImmediate(lexeme)
		}

		tokens = append(tokens, tok)
	}

	return tokens, nil
}

func parseImmediate(value string) int32 {
	value = strings.ToLower(value)
	value = strings.TrimPrefix(value, "#")
	num, _ := strconv.ParseInt(value, 0, 32)

	return int32(num)
}

func parseRegister(reg string) uint32 {
	switch true {
	case reg == "SP":
		return 13
	case reg == "LR":
		return 14
	case reg == "PC":
		return 15
	default:
		reg = reg[1:]
		out, _ := strconv.ParseInt(reg, 0, 32)
		return uint32(out)
	}
}
