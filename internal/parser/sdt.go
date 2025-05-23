package parser

import (
	"cscasm/internal/cond"
	"cscasm/internal/token"
)

type SdtIns struct {
	Cond   uint32
	I      uint32
	P      uint32
	U      uint32
	B      uint32
	W      uint32
	L      uint32
	Rn     uint32
	Rd     uint32
	Offset uint32
}

func (i *SdtIns) Encoding() uint32 {
	var e uint32 = 0

	e |= i.Cond << 28
	e |= 0b01 << 26
	e |= i.I << 25
	e |= i.P << 24
	e |= i.U << 23
	e |= i.B << 22
	e |= i.W << 21
	e |= i.L << 20
	e |= i.Rn << 16
	e |= i.Rd << 12
	e |= i.Offset

	return e
}

func (p *Parser) newSdtIns(mnemonic string, toks []*token.Token) (*SdtIns, error) {
	var l uint32 = 0
	if mnemonic[0] == 'L' {
		l = 1
	}

	rd := toks[1].Literal.(uint32)
	rn := toks[2].Literal.(uint32)

	return &SdtIns{
		Cond: cond.AL,
		I:    0,
		P:    0,
		U:    0,
		W:    0,
		L:    l,
		Rn:   rn,
		Rd:   rd,
	}, nil
}
