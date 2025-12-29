package main

import (
	"flag"
	"fmt"

	"github.com/DevonDF/go-6502-emulator/assembler"
)

func main() {
	input := flag.String("input", "", "assembly file input")
	output := flag.String("output", "", "compiled bytecode output")
	flag.Parse()

	if *input == "" || *output == "" {
		flag.Usage()
		return
	}

	err := assembler.Assemble(*input, *output)
	if err != nil {
		fmt.Printf("error occured during assembly: %w", err)
	} else {
		fmt.Printf("wrote assembled bytecode to %s", *output)
	}
}
