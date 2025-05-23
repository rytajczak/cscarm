package assembler

import (
	"bufio"
	"cscasm/internal/lexer"
	"cscasm/internal/parser"
	"cscasm/internal/token"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func AssembleFile(filename string) ([]uint32, error) {
	file, err := os.Open(filename)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		return []uint32{}, fmt.Errorf("%s: failed to open file %q: %w", red("error"), filename, err)
	}
	defer file.Close()

	lex := lexer.NewLexer()
	tokenizedLines := [][]*token.Token{}
	symbolTable := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		tokens, _ := lex.TokenizeLine(line)

		if len(tokens) == 0 {
			continue
		}

		if tokens[0].Type == token.LABEL {
			name := tokens[0].Literal.(string)
			symbolTable[name] = len(tokenizedLines) + 1
			continue
		}

		tokenizedLines = append(tokenizedLines, tokens)
	}

	par := parser.NewParser(symbolTable)
	bytes := []uint32{}

	for i, toks := range tokenizedLines {
		ins, err := par.ParseInstructionFromTokens(toks)
		if err != nil {
			fmt.Printf("%d %0x\n", i+1, 0)
			continue
		}

		fmt.Printf("%d %0x\n", i+1, ins.Encoding())
		bytes = append(bytes, ins.Encoding())
	}

	return bytes, nil
}
