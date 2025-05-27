package main

import (
	"cscasm/internal/assembler"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cscarm <file>")
		os.Exit(1)
		return
	}
	filename := os.Args[1]

	bytes, err := assembler.AssembleFile(filename)
	if err != nil {
		fmt.Printf("%s\n\n", err)
		os.Exit(1)
	}

	file, err := os.Create("a.out")
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("%s: %s\n\n", red("error"), err.Error())
		os.Exit(1)
	}

	binary.Write(file, binary.LittleEndian, bytes)
}
