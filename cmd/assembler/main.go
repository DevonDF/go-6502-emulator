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

	assembler := assembler.NewAssembler()
	err := assembler.Assemble(*input)
	if err != nil {
		fmt.Printf("error occured during assembly: %v", err)
		return
	}

	fmt.Printf(assembler.ToHexDump())

	err = assembler.Write(*output)
	if err != nil {
		fmt.Printf("failed to write bytecode: %v", err)
		return
	}

	fmt.Printf("wrote assembled bytecode to %s", *output)
}
