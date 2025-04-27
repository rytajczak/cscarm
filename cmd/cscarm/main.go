package main

import (
	"cscasm/internal/assembler"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cscarm <file>")
		os.Exit(1)
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
		return
	}
	defer file.Close()

	bytes := assembler.AssembleFile(file)
	fmt.Println(bytes)
}
