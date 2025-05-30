package lexer

import (
	"bufio"
	"io"
	"slices"
	"strings"
	"unicode"

	"github.com/rytajczak/cscarm/internal/token"
)

type Lexer struct {
	readSeeker  io.ReadSeeker
	reader      *bufio.Reader
	line        int
	col         int
	currentRune rune
	prevToken   *token.Token
	nextToken   *token.Token
	eof         bool
}

func NewLexer(reader io.Reader) *Lexer {
	readSeeker := reader.(io.ReadSeeker)
	l := &Lexer{
		readSeeker: readSeeker,
		reader:     bufio.NewReader(readSeeker),
		line:       1,
		col:        0,
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
	r := l.currentRune
	tok := &token.Token{Line: l.line, Col: l.col}

	switch true {
	case l.eof:
		tok.Type = token.EOF
		return tok
	case r == '@' || r == ';':
		tok.Type = token.COMMENT
		tok.Literal = l.readComment()
	case unicode.IsLetter(r) && !unicode.IsDigit(l.peekRune()):
		tok.Literal = l.readText()
		switch true {
		case slices.Contains([]string{"SP", "LR", "PC"}, strings.ToUpper(tok.Literal)):
			tok.Type = token.REGISTER
		case l.PeekToken().Type == token.COLON:
			tok.Type = token.LABEL
		case l.prevToken != nil && l.prevToken.Type == token.MNEMONIC:
			tok.Type = token.IDENT
		default:
			tok.Type = token.MNEMONIC
		}
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
		l.line++
		l.col = 0
		tok.Type = token.NEWLINE
		l.readRune()
	default:
		tok.Type = token.ILLEGAL
		l.readRune()
	}

	l.prevToken = tok
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

	l.col++
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

func (l *Lexer) readText() string {
	var text string
	for unicode.IsLetter(l.currentRune) || l.currentRune == '_' {
		text += string(l.currentRune)
		l.readRune()
	}
	return text
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

func (l *Lexer) Reset() {
	_, err := l.readSeeker.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	l.reader = bufio.NewReader(l.readSeeker)
	l.line = 1
	l.col = 0
	l.eof = false
	l.readRune()
	l.prevToken = nil
}
