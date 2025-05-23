package parser

import (
	"cscasm/internal/cond"
	"cscasm/internal/token"
)

type MovwIns struct {
	Cond  uint32
	Rd    uint32
	Imm16 uint32
}

func (i *MovwIns) Encoding() uint32 {
	var e uint32

	e |= i.Cond << 28
	e |= 0b00110000 << 20
	e |= (i.Imm16 & 0xF000) << 4
	e |= i.Rd << 12
	e |= i.Imm16 & 0x0FFF

	return e
}

func (p *Parser) newMovwIns(toks []*token.Token) (*MovwIns, error) {
	rd := toks[1].Literal.(uint32)
	imm16 := toks[2].Literal.(int32)

	return &MovwIns{
		Cond:  cond.AL,
		Rd:    rd,
		Imm16: uint32(imm16),
	}, nil
}

type MovtIns struct {
	Cond  uint32
	Rd    uint32
	Imm16 uint32
}

func (i *MovtIns) Encoding() uint32 {
	var e uint32

	e |= i.Cond << 28
	e |= 0b00110100 << 20
	e |= (i.Imm16 & 0xF000) << 4
	e |= i.Rd << 12
	e |= i.Imm16 & 0x0FFF

	return e
}

func (p *Parser) newMovtIns(toks []*token.Token) (*MovtIns, error) {
	rd := toks[1].Literal.(uint32)
	imm16 := toks[2].Literal.(int32)

	return &MovtIns{
		Cond:  cond.AL,
		Rd:    rd,
		Imm16: uint32(imm16),
	}, nil
}
