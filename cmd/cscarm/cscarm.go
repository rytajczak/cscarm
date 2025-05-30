package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rytajczak/cscarm/internal/parser"
)

const MIN_ARGUEMENTS = 2

func main() {
	outputFile := flag.String("o", "a.out", "The file to output to")

	flag.Parse()

	if len(os.Args) < MIN_ARGUEMENTS {
		fatalErrorf("no input file")
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fatalErrorf(err)
	}

	par := parser.NewParser(inputFile)

	bytes, err := par.Parse()
	if err != nil {
		fatalErrorf("failed to parse bytes: %w", err)
	}

	outFile, err := os.Create(*outputFile)
	if err != nil {
		fatalErrorf("failed to create output file: %w", err)
	}

	if err = binary.Write(outFile, binary.LittleEndian, bytes); err != nil {
		fatalErrorf("failed to write binary output: %w", err)
	}
}

func fatalErrorf(a ...any) {
	log.SetFlags(0)
	red := color.New(color.FgRed).SprintFunc()
	log.Fatalf(red("error: ")+"%v", fmt.Sprint(a...))
}
