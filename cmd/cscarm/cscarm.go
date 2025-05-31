package main

import (
	"encoding/binary"
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rytajczak/cscarm/internal/parser"
)

const MIN_ARGUEMENTS = 2

func main() {
	outputFile := flag.String("o", "a.out", "The file to output to")
	log.SetFlags(0)
	red := color.New(color.FgRed).SprintFunc()

	flag.Parse()

	if len(os.Args) < MIN_ARGUEMENTS {
		log.Fatalf("%s no input file", red("error:"))
	}

	par, err := parser.NewParser(os.Args[1])
	if err != nil {
		log.Fatalf("%s %s", red("error:"), err.Error())
	}

	bytes, errs := par.Parse()
	for _, err := range errs {
		log.Printf("%s\n\n", err.Error())
	}
	if errs != nil {
		os.Exit(1)
	}

	outFile, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("%s %s", red("error:"), err.Error())
	}

	if err = binary.Write(outFile, binary.LittleEndian, bytes); err != nil {
		log.Fatalf("%s %s", red("error:"), err.Error())
	}
}
