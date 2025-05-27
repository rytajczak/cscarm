package assembler

import (
	"cscasm/internal/parser"
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

	par := parser.NewParser(file)

	return par.Parse()
}
