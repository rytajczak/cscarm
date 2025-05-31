package parser

import (
	"fmt"

	"github.com/rytajczak/cscarm/internal/cond"
	"github.com/rytajczak/cscarm/internal/token"
)

func (p *Parser) encodeSingleDataTransferINS(mnemonic string) (uint32, error) {
	var l uint32 = 0
	if mnemonic[0] == 'L' {
		l = 1
	}

	rd, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	if p.currentToken.Type != token.LBRACK {
		return 0, fmt.Errorf("invalid format")
	}
	p.nextToken()

	rn, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	e := cond.AL << 28
	e |= 0b01 << 26
	e |= l << 20
	e |= rn << 16
	e |= rd << 12

	return e, err
}
