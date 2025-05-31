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
	opCode, exists := dpOpCodes[mnemonic[:3]]
	if !exists {
		return 0, fmt.Errorf("failed to identiy opcode")
	}

	var s uint32 = 0
	if len(mnemonic) == 4 && mnemonic[3] == 'S' {
		s = 1
	}

	rd, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	rn, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	var i uint32 = 0
	if p.currentToken.Type == token.IMMEDIATE {
		i = 1
	}

	op2, err := p.consumeImmediate()
	if err != nil {
		return 0, err
	}

	e := cond.AL << 28
	e |= i << 25
	e |= opCode << 21
	e |= s << 20
	e |= rn << 16
	e |= rd << 12
	e |= op2

	return e, nil
}
