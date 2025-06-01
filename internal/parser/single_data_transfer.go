package parser

import (
	"fmt"

	"github.com/rytajczak/cscarm/internal/cond"
	"github.com/rytajczak/cscarm/internal/token"
)

func (p *Parser) encodeSingleDataTransferINS(mnemonic string) (uint32, error) {
	var ibit, pbit, ubit, bbit, wbit, lbit, rn, rd, offset uint32

	if mnemonic[0] == 'L' {
		lbit = 1
	}

	rd, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	if p.currentToken.Type != token.LBRACK {
		return 0, fmt.Errorf("invalid format")
	}
	p.nextToken()

	rn, err = p.consumeRegister()
	if err != nil {
		return 0, err
	}

	if p.currentToken.Type == token.RBRACK {
		p.nextToken()
		if p.currentToken.Type == token.IMMEDIATE {
			ubit = 1
			offset, _, err = p.consumeImmediate()
			if err != nil {
				return 0, fmt.Errorf("dont know how this happened")
			}
		}
	} else if p.currentToken.Type == token.IMMEDIATE {
		pbit = 1
		var isNeg bool = false
		offset, isNeg, err = p.consumeImmediate()
		if err != nil {
			return 0, fmt.Errorf("dont know how this happened")
		}
		if isNeg {
			offset = ^offset + 1
		}
		p.nextToken()
	}

	if p.currentToken.Type == token.EXCLAM {
		wbit = 1
	}

	e := cond.AL << 28
	e |= 0b01 << 26
	e |= ibit << 25
	e |= pbit << 24
	e |= ubit << 23
	e |= bbit << 22
	e |= wbit << 21
	e |= lbit << 20
	e |= rn << 16
	e |= rd << 12
	e |= offset & 0x00000FFF

	return e, err
}
