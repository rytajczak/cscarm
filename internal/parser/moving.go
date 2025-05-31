package parser

import (
	"github.com/rytajczak/cscarm/internal/cond"
)

func (p *Parser) encodeMovwINS() (uint32, error) {
	rd, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	imm16, err := p.consumeImmediate()
	if err != nil {
		return 0, err
	}

	e := cond.AL << 28
	e |= 0b0011_0000 << 20
	e |= (imm16 & 0xF000) << 4
	e |= rd << 12
	e |= imm16 & 0x0FFF

	return e, nil
}

func (p *Parser) encodeMovtINS() (uint32, error) {
	rd, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	imm16, err := p.consumeImmediate()
	if err != nil {
		return 0, err
	}

	e := cond.AL << 28
	e |= 0b0011_0100 << 20
	e |= (imm16 & 0xF000) << 4
	e |= rd << 12
	e |= imm16 & 0x0FFF

	return e, nil
}
