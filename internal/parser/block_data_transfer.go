package parser

import (
	"fmt"

	"github.com/rytajczak/cscarm/internal/cond"
	"github.com/rytajczak/cscarm/internal/token"
)

func (p *Parser) encodeBlockDataTransferINS(mnemonic string) (uint32, error) {
	var pbit, ubit, wbit, lbit, rn, regList uint32

	if mnemonic[0] == 'L' {
		lbit = 1
	}

	switch mnemonic {
	case "LDMED", "LDMIB", "STMFA", "STMIB":
		pbit = 1
		ubit = 1
	case "LDMFD", "LDMIA", "STMEA", "STMIA":
		ubit = 1
	case "LDMEA", "LDMDB", "STMFD", "STMDB":
		pbit = 1
	case "LDMFA", "LDMDA", "STMED", "STMDA":
	default:
		return 0, fmt.Errorf("unknown addressing mode")
	}

	rn, err := p.consumeRegister()
	if err != nil {
		return 0, err
	}

	if p.currentToken.Type == token.EXCLAM {
		wbit = 1
		p.nextToken()
	}

	if p.currentToken.Type != token.LBRACE {
		return 0, fmt.Errorf("invalid format")
	}
	p.nextToken()

	regList, err = p.consumeRegList()
	if err != nil {
		return 0, err
	}

	e := cond.AL << 28
	e |= 0b100 << 25
	e |= pbit << 24
	e |= ubit << 23
	e |= wbit << 21
	e |= lbit << 20
	e |= rn << 16
	e |= regList

	return e, nil
}

func (p *Parser) consumeRegList() (uint32, error) {
	var regList uint32
	for p.currentToken.Type != token.RBRACE {
		if p.currentToken.Type != token.REGISTER {
			p.nextToken()
			continue
		}
		reg, err := p.consumeRegister()
		if err != nil {
			return 0, fmt.Errorf("honestly, idk how this one happened")
		}
		if p.currentToken.Type == token.MINUS {
			p.nextToken()
			endReg, err := p.consumeRegister()
			if err != nil {
				return 0, fmt.Errorf("expected end register for range")
			}
			for i := reg; i <= endReg; i++ {
				regList |= 1 << i
			}
		} else {
			regList |= 1 << reg
		}
	}
	p.nextToken()
	return regList, nil
}
