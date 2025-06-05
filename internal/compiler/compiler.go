package compiler

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/rytajczak/cscarm/internal/lexer"
	"github.com/rytajczak/cscarm/internal/token"
)

type Compiler struct {
	out          *os.File
	lexer        *lexer.Lexer
	currentToken *token.Token
}

func NewCompiler(file *os.File) *Compiler {
	lex := lexer.NewLexer(file)
	c := &Compiler{
		lexer: lex,
	}
	c.nextToken()
	return c
}

func (c *Compiler) nextToken() {
	c.currentToken = c.lexer.NextToken()
}

func (c *Compiler) append(ins string, a ...any) {
	fmt.Fprintf(c.out, ins+"\n", a...)
}

func (c *Compiler) Compile() error {
	file, err := os.Create("a.s")
	if err != nil {
		fmt.Printf("failed to create temp file")
	}
	c.out = file
	defer c.out.Close()

	blinkCount, err := strconv.Atoi(c.currentToken.Literal)
	if err != nil {
		return fmt.Errorf("thats not a number, silly")
	}

	c.initStack()
	c.placeValueInRegister(0x3f200000, "r4")
	c.setPinMode(21, OUTPUT)
	c.placeValueInRegister(blinkCount, "r6")
	c.push("r6")
	c.createLabel("loop")
	c.branch("blink", true, "")
	c.append("\tsubs r6, r6, #1")
	c.branch("loop", false, "ne")
	c.delay(4000)
	c.peek("r6")
	c.branch("loop", false, "")

	c.blinkSubroutine()
	c.delaySubroutine()

	return nil
}

func (c *Compiler) placeValueInRegister(value int, reg string) {
	low := value & 0x0000ffff
	high := value & 0xffff0000 >> 16
	c.append("\tmovw %s, #0x%x", reg, low)
	c.append("\tmovt %s, #0x%x", reg, high)
}

func (c *Compiler) createLabel(label string) {
	fmt.Fprintf(c.out, "\n%s:\n", label)
}

func (c *Compiler) branch(label string, link bool, cond string) {
	if cond == "AL" {
		cond = ""
	}
	if link {
		c.append("\tbl%s %s", cond, label)
	} else {
		c.append("\tb%s %s", cond, label)
	}
}

func (c *Compiler) setPinMode(pin int, mode PinModeType) error {
	var gpfsel = 0x0
	switch true {
	case pin >= 0 && pin < 10:
	case pin >= 10 && pin < 20:
		gpfsel = 0x4
	case pin >= 20 && pin < 30:
		gpfsel = 0x8
	default:
		return fmt.Errorf("invalid pin number")
	}

	var pinMode = 0b000
	switch mode {
	case INPUT:
		pinMode = 0b000
	case OUTPUT:
		pinMode = 0b001
	default:
		return fmt.Errorf("invalid pin mode")
	}

	fsel := (pin % 10) * 3
	fsel = pinMode << fsel

	c.append("\tadd r2, r4, #%x", gpfsel)
	c.append("\tldr r3, [r2]")
	c.append("\torr r3, r3, #0x%x", fsel)
	c.append("\tstr r3, [r2]")
	return nil
}

func (c *Compiler) initStack() {
	c.placeValueInRegister(0, "sp")
}

func (c *Compiler) push(reg string) {
	c.append("\tstr %s, [sp], #4", reg)
}

func (c *Compiler) pop(reg string) {
	c.append("\tldr %s, [sp, #-4]!", reg)
}

func (c *Compiler) peek(reg string) {
	c.append("\tldr %s, [sp, #-4]", reg)
}

func (c *Compiler) delay(ms int) {
	time := 1_400_000_000 * (float64(ms) / 1000)
	loopCount := int(math.Ceil(time / float64(1400)))
	c.placeValueInRegister(loopCount, "r5")
	c.push("r5")
	c.branch("delay", true, "AL")
}

func (c *Compiler) pinWrite(pin int, write PinWriteType) {
	set := 0x28
	if write == HIGH {
		set = 0x1c
	}
	c.append("\tadd r3, r4, #0x%x", set)
	c.placeValueInRegister(1<<pin, "r2")
	c.append("\tstr r2, [r3]")
}

func (c *Compiler) blinkSubroutine() {
	c.createLabel("blink")
	c.push("lr")
	c.pinWrite(21, HIGH)
	c.delay(1000)
	c.pinWrite(21, LOW)
	c.delay(1000)
	c.pop("lr")
	c.append("\tbx lr")
}

func (c *Compiler) delaySubroutine() {
	c.createLabel("delay")
	c.pop("r5")
	c.createLabel("delay_loop")
	c.append("\tsubs r5, r5, #1")
	c.branch("delay_loop", false, "ge")
	c.append("\tbx lr")
}
