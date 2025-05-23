package parser

import (
	"cscasm/internal/cond"
	"cscasm/internal/token"
)

type BblIns struct {
	Cond   uint32
	L      uint32
	Offset uint32
}

func (i *BblIns) Encoding() uint32 {
	var e uint32 = 0

	e |= i.Cond << 28
	e |= 0b101 << 25
	e |= i.L << 24
	e |= i.Offset & 0x00FFFFFF

	return e
}

func (p *Parser) newBblIns(mnemonic string, toks []*token.Token) (*BblIns, error) {
	c := cond.AL
	if len(mnemonic) >= 3 {
		c = cond.CondCodeMap[mnemonic[len(mnemonic)-2:]]
	}

	var l uint32 = 0
	if len(mnemonic) > 1 && mnemonic[1] == 'L' {
		l = 1
	}

	var offset uint32 = 0
	if toks[1].Type == token.IMMEDIATE {
		offset = uint32(toks[1].Literal.(int32))
	}
	if toks[1].Type == token.IDENTIFIER {
		location := p.ST[toks[1].Lexeme]
		offset = uint32(location - p.currentIns - 2)
	}

	return &BblIns{
		Cond:   c,
		L:      l,
		Offset: offset,
	}, nil
}
