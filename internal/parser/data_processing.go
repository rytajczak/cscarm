package parser

import (
	"fmt"

	"github.com/rytajczak/cscarm/internal/cond"
	"github.com/rytajczak/cscarm/internal/token"
)

var dpOpCodes = map[string]uint32{
	"SUB": 0b0010,
	"ADD": 0b0100,
	"ORR": 0b1100,
}

func (p *Parser) encodeDataProcessingINS(mnemonic string) (uint32, error) {
	var ibit, opcode, sbit, rn, rd, op2 uint32

	opcode, exists := dpOpCodes[mnemonic[:3]]
	if !exists {
		return 0, fmt.Errorf("failed to identiy opcode")
	}

	if len(mnemonic) == 4 && mnemonic[3] == 'S' {
		sbit = 1
	}

	rd, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	rn, err = p.consumeRegister()
	if err != nil {
		return 0, err
	}

	if p.currentToken.Type == token.IMMEDIATE {
		ibit = 1
	}

	op2, _, err = p.consumeImmediate()
	if err != nil {
		return 0, err
	}

	e := cond.AL << 28
	e |= ibit << 25
	e |= opcode << 21
	e |= sbit << 20
	e |= rn << 16
	e |= rd << 12
	e |= op2

	return e, nil
}
