package parser

import (
	"cscasm/internal/cond"
	"cscasm/internal/token"
)

var dpOpCodes = map[string]uint32{
	"SUB": 0b0010,
	"ADD": 0b0100,
	"ORR": 0b1100,
}

type DpIns struct {
	Cond   uint32
	I      uint32
	OpCode uint32
	S      uint32
	Rn     uint32
	Rd     uint32
	Op2    uint32
}

func (i *DpIns) Encoding() uint32 {
	var e uint32

	e |= i.Cond << 28
	e |= i.I << 25
	e |= i.OpCode << 21
	e |= i.S << 20
	e |= i.Rn << 16
	e |= i.Rd << 12
	e |= i.Op2

	return e
}

func (p *Parser) newDpIns(mnemonic string, toks []*token.Token) (*DpIns, error) {
	opCode := dpOpCodes[mnemonic[:3]]

	var s uint32 = 0
	if len(mnemonic) == 4 && mnemonic[3] == 'S' {
		s = 1
	}

	rd := toks[1].Literal.(uint32)
	rn := toks[2].Literal.(uint32)
	op2 := toks[3].Literal.(int32)

	var i uint32 = 0
	if toks[3].Type == token.IMMEDIATE {
		i = 1
	}

	return &DpIns{
		Cond:   cond.AL,
		I:      i,
		OpCode: opCode,
		S:      s,
		Rn:     rn,
		Rd:     rd,
		Op2:    uint32(op2),
	}, nil
}
