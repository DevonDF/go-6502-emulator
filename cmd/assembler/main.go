package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/DevonDF/go-6502-emulator/assembler"
)

func main() {
	input := flag.String("input", "", "assembly file input")
	output := flag.String("output", "", "compiled bytecode output")
	loadAddressStr := flag.String("loadAddr", "0x8000", "hex address of where the ROM will be loaded")
	flag.Parse()

	if *input == "" || *output == "" {
		flag.Usage()
		return
	}

	loadAddress, err := strconv.ParseUint((*loadAddressStr)[2:], 16, 16)
	if err != nil {
		fmt.Printf("loadAddr failed to be parsed correctly: %v", err)
		flag.Usage()
		return
	}

	assembler := assembler.NewAssembler(uint16(loadAddress))

	err = assembler.Assemble(*input)
	if err != nil {
		fmt.Printf("error occured during assembly: %v", err)
		return
	}

	fmt.Println(assembler.ToHexDump())

	err = assembler.Write(*output)
	if err != nil {
		fmt.Printf("failed to write bytecode: %v", err)
		return
	}

	fmt.Printf("wrote assembled bytecode to %s", *output)
}
