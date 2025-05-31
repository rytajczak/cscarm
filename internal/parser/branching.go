package parser

import (
	"fmt"

	"github.com/rytajczak/cscarm/internal/cond"
)

func (p *Parser) encodeBranchINS(mnemonic string) (uint32, error) {
	c := cond.AL
	var l uint32 = 0
	switch len(mnemonic) - 2 {
	case -1:
	case 0:
		l = 1
	case 1:
		cond, exists := cond.CondCodeMap[mnemonic[1:]]
		if exists {
			c = cond
		}
	case 2:
		cond, exists := cond.CondCodeMap[mnemonic[2:]]
		if exists {
			c = cond
		}
		l = 1
	default:
		return 0, fmt.Errorf("unknown mnemonic")
	}

	label, err := p.consumeIdent()
	if err != nil {
		return 0, err
	}
	target, exists := p.symbolTable[label]
	if !exists {
		return 0, fmt.Errorf("unknown label '%s'", label)
	}
	offset := uint32(target - p.currentLine - 2)

	e := c << 28
	e |= 0b101 << 25
	e |= l << 24
	e |= offset & 0x00FFFFFF

	return e, nil
}

func (p *Parser) encodeBranchExchangeINS() (uint32, error) {
	rm, err := p.consumeRegister()
	if err != nil {
		return 0, fmt.Errorf("failed to parse branch: %w", err)
	}

	e := cond.AL << 28
	e |= 0b0001_0010_1111_1111_1111_0001 << 4
	e |= rm

	return e, nil
}
