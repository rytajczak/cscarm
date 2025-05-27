package lexer

import (
	"bufio"
	"cscasm/internal/token"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	reader      *bufio.Reader
	line        int
	col         int
	currentRune rune
	eof         bool
	nextToken   *token.Token
}

func NewLexer(reader io.Reader) *Lexer {
	l := &Lexer{
		reader: bufio.NewReader(reader),
		line:   1,
		col:    0,
	}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	if l.nextToken != nil {
		tok := l.nextToken
		l.nextToken = nil
		return tok
	}

	return l.generateToken()
}

func (l *Lexer) PeekToken() *token.Token {
	if l.nextToken == nil {
		l.nextToken = l.NextToken()
	}

	return l.nextToken
}

func (l *Lexer) generateToken() *token.Token {
	l.skipWhitespace()
	tok := &token.Token{Line: l.line, Col: l.col}
	r := l.currentRune

	if l.eof {
		tok.Type = token.EOF
		return tok
	}

	switch true {
	case r == '@' || r == ';':
		tok.Type = token.COMMENT
		tok.Literal = l.readComment()
	case unicode.IsLetter(r) && !unicode.IsDigit(l.peekRune()):
		tok.Type = token.IDENT
		tok.Literal = l.readIdentifier()
	case r == 'r' && unicode.IsDigit(l.peekRune()):
		tok.Type = token.REGISTER
		tok.Literal = l.readRegister()
	case r == '#':
		tok.Type = token.IMMEDIATE
		tok.Literal = l.readImmediate()
	case r == '[':
		tok.Type = token.LBRACK
		l.readRune()
	case r == '{':
		tok.Type = token.LBRACE
		l.readRune()
	case r == ',':
		tok.Type = token.COMMA
		l.readRune()
	case r == ']':
		tok.Type = token.RBRACK
		l.readRune()
	case r == '}':
		tok.Type = token.RBRACE
		l.readRune()
	case r == ':':
		tok.Type = token.COLON
		l.readRune()
	case r == '\n':
		tok.Type = token.NEWLINE
		l.readRune()
	default:
		tok.Type = token.ILLEGAL
		l.readRune()
	}
	return tok
}

func (l *Lexer) readRune() {
	if l.eof {
		return
	}

	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.eof = true
			l.currentRune = 0
			return
		}
		return
	}

	if r == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}

	l.currentRune = r
}

func (l *Lexer) peekRune() rune {
	if l.eof {
		return 0
	}
	r, err := l.reader.Peek(1)
	if err != nil || len(r) == 0 {
		return 0
	}
	return rune(r[0])
}

func (l *Lexer) skipWhitespace() {
	for !l.eof && unicode.IsSpace(l.currentRune) && l.currentRune != '\n' {
		l.readRune()
	}
}

func (l *Lexer) readComment() string {
	var comment string
	for l.currentRune != '\n' {
		comment += string(l.currentRune)
		l.readRune()
	}

	comment = strings.Trim(comment, "@")
	comment = strings.Trim(comment, ";")
	comment = strings.TrimLeft(comment, " ")

	return comment
}

func (l *Lexer) readIdentifier() string {
	var identifier string
	for unicode.IsLetter(l.currentRune) || l.currentRune == '_' {
		identifier += string(l.currentRune)
		l.readRune()
	}
	return identifier
}

func (l *Lexer) readRegister() string {
	var register string
	for unicode.IsLetter(l.currentRune) || unicode.IsNumber(l.currentRune) {
		register += string(l.currentRune)
		l.readRune()
	}
	return register
}

func (l *Lexer) readImmediate() string {
	var immediate string
	for !unicode.IsSpace(l.currentRune) {
		immediate += string(l.currentRune)
		l.readRune()
	}
	immediate = strings.Trim(immediate, "#")
	return immediate
}
