package assembler

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DevonDF/go-6502-emulator/instructions"
)

// assembleLine takes a line from assembley code and assembles it into bytecode.
func assembleLine(line string) ([]byte, error) {
	// split out instruction and operands
	assemblyString, operands, found := strings.Cut(line, " ")
	if !found {
		return nil, fmt.Errorf("invalid assembly line encountered: %s", line)
	}

	// attempt to discover the addressing mode
	var addressingMode instructions.AddressingMode
	var operandBytes []byte

	if operands == "A" { // accumulator addressing
		addressingMode = instructions.AddrAccumulator
		operandBytes = []byte{}
	} else if operands[0] == '#' { // immediate addressing
		addressingMode = instructions.AddrImmediate
		if operands[1] != '&' {
			return nil, fmt.Errorf("invalid immediate addressing used: %s", line)
		}
		v, err := strconv.ParseUint(operands[2:], 16, 8)
		if err != nil {
			return nil, err
		}
		operandBytes = []byte{byte(v)}
	} else if operands[0] == '&' { // zeropage/absolute
		commaFound := strings.Contains(operands, ",")
		if !commaFound {
			if len(operands) == 3 { // zeropage addressing
				addressingMode = instructions.AddrZeropage
				v, err := strconv.ParseUint(operands[1:], 16, 8)
				if err != nil {
					return nil, err
				}
				operandBytes = []byte{byte(v)}
			} else { // absolute addressing
				addressingMode = instructions.AddrAbsolute
				v, err := strconv.ParseUint(operands[1:], 16, 8)
				if err != nil {
					return nil, err
				}
				operandBytes = []byte{byte((v >> 8) & 0xFF), byte(v & 0xFF)}
			}
		} else { // X/Y registers used, disgusting
			return nil, fmt.Errorf("unimplemented use of addressing mode: %s", line)
		}
	} else if operands[0] == '(' { // indirect addressing
		return nil, fmt.Errorf("unimplemented use of addressing mode: %s", line)
	} else {
		return nil, fmt.Errorf("unrecognised addressing mode: %s", line)
	}

	instruction := instructions.InstructionFromAssembly(assemblyString, addressingMode)
	if instruction == nil {
		return nil, fmt.Errorf("unknown instruction: %s", line)
	}

	bytecode := []byte{instruction.Opcode}
	bytecode = append(bytecode, operandBytes...)
	return bytecode, nil
}

func Assemble(inputFilePath string, outputFilePath string) error {

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	bytecode := []byte{}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		assemblyLine := scanner.Text()
		lineBytecode, err := assembleLine(assemblyLine)
		if err != nil {
			return err
		}
		bytecode = append(bytecode, lineBytecode...)
	}

	fmt.Print(hex.Dump(bytecode))

	err = os.WriteFile(outputFilePath, bytecode, 0644)
	if err != nil {
		return err
	}

	return nil
}
