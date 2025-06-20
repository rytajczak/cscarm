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
	mnemOnLine  bool
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
	l.skipWhitespace()
	r := l.currentRune
	tok := &token.Token{Line: l.line, Col: l.col}

	switch true {
	case l.eof:
		tok.Type = token.EOF
		return tok
	case unicode.IsLetter(r) && !unicode.IsDigit(l.peekRune()):
		tok.Literal = l.readText()
		switch true {
		case slices.Contains([]string{"SP", "LR", "PC"}, strings.ToUpper(tok.Literal)):
			tok.Type = token.REGISTER
		case l.currentRune == ':':
			tok.Type = token.LABEL
			l.readRune()
		case !l.mnemOnLine:
			tok.Type = token.MNEMONIC
			l.mnemOnLine = true
		default:
			tok.Type = token.IDENT
		}
	case r == 'r' && unicode.IsDigit(l.peekRune()):
		tok.Type = token.REGISTER
		tok.Literal = l.readRegister()
	case r == '#' || unicode.IsDigit(l.currentRune):
		tok.Type = token.IMMEDIATE
		tok.Literal = l.readImmediate()
	case r == '[':
		tok.Type = token.LBRACK
		l.readRune()
	case r == '{':
		tok.Type = token.LBRACE
		l.readRune()
	case r == ']':
		tok.Type = token.RBRACK
		l.readRune()
	case r == '}':
		tok.Type = token.RBRACE
		l.readRune()
	case r == '!':
		tok.Type = token.EXCLAM
		l.readRune()
	case r == '-':
		tok.Type = token.MINUS
		l.readRune()
	case r == '@' || r == ';':
		tok.Type = token.COMMENT
		tok.Literal = l.readComment()
	case r == '\n':
		l.line++
		l.col = 0
		tok.Type = token.NEWLINE
		l.mnemOnLine = false
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
	for !l.eof && (unicode.IsSpace(l.currentRune) || l.currentRune == ',') && l.currentRune != '\n' {
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
	for !unicode.IsSpace(l.currentRune) && l.currentRune != ']' {
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
	l.mnemOnLine = false
	l.readRune()
}
