package main

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rytajczak/cscarm/internal/compiler"
	"github.com/rytajczak/cscarm/internal/parser"
)

const MIN_ARGUEMENTS = 2

func main() {
	log.SetFlags(0)

	if len(os.Args) < MIN_ARGUEMENTS {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf("%s no input file", red("error:"))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf("%s failed to open file", red("error:"))
	}
	defer file.Close()

	comp := compiler.NewCompiler(file)
	if err := comp.Compile(); err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf("%s failed to compile file", red("error:"))
	}

	file.Close()
	file, err = os.Open("a.s")
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf("%s failed to open file", red("error:"))
	}

	par := parser.NewParser(file)

	bytes, errs := par.Parse()
	for _, err := range errs {
		log.Printf("%s\n\n", err.Error())
	}
	if errs != nil {
		os.Exit(1)
	}

	out, err := os.Create("a.out")
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf("%s %s", red("error:"), err.Error())
	}

	if err = binary.Write(out, binary.LittleEndian, bytes); err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Fatalf("%s %s", red("error:"), err.Error())
	}
}
