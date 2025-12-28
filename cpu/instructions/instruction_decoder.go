package instructions

import (
	"bufio"
	"fmt"
)

// Instruction defines an instruction within the 6502 instruction set.
type Instruction struct {
	opcode  byte                // the opcode for the given instruction.
	size    byte                // the number of bytes this instruction takes.
	cycles  byte                // the number of cpu cycles this operation takes.
	handler *InstructionHandler // the handler for this instruction.
}

// DecodedInstruction defines a decoded instruction within the 6502 instruction set.
type DecodedInstruction struct {
	instruction *Instruction // the instruction that this relates to.
	operands    []byte       // the operands for the given instruction.
}

// getInstructionSet returns the set of instructions as a map of opcode to Instruction.
func getInstructionSet() map[byte]Instruction {
	return map[byte]Instruction{

		0x00: { // BRK
			opcode:  0x00,
			size:    1,
			cycles:  7,
			handler: nil,
		},
		0x01: { // ORA x,ind
			opcode:  0x01,
			size:    2,
			cycles:  6,
			handler: nil,
		},
		0x05: { // ORA oper
			opcode:  0x05,
			size:    2,
			cycles:  3,
			handler: nil,
		},
		0x06: { // ASL oper
			opcode:  0x06,
			size:    2,
			cycles:  5,
			handler: nil,
		},
		0x08: { // PHP
			opcode:  0x08,
			size:    1,
			cycles:  3,
			handler: nil,
		},
		0x09: { // ORA #oper
			opcode:  0x09,
			size:    2,
			cycles:  2,
			handler: nil,
		},
		0x0A: { // ASL A
			opcode:  0x0A,
			size:    1,
			cycles:  2,
			handler: nil,
		},
		0x0D: { // ORA oper
			opcode:  0x0D,
			size:    3,
			cycles:  4,
			handler: nil,
		},
		0x0E: { // ASL oper
			opcode:  0x0E,
			size:    3,
			cycles:  6,
			handler: nil,
		},
	}
}

// DecodeNextInstruction decodes the next instruction from a buffered byte reader and returns a DecodedInstruction.
func DecodeNextInstruction(reader bufio.Reader) (DecodedInstruction, error) {
	opcode, err := reader.ReadByte()
	if err != nil {
		return DecodedInstruction{}, err
	}

	instructions := getInstructionSet()
	instruction, found := instructions[opcode]
	if !found {
		return DecodedInstruction{}, fmt.Errorf("no opcode found for %02X", opcode)
	}

	operands := make([]byte, instruction.size-1)
	_, err = reader.Read(operands)
	if err != nil {
		return DecodedInstruction{}, fmt.Errorf("failed to read operands for opcode %02X: %w", opcode, err)
	}

	decodedInstruction := DecodedInstruction{
		instruction: &instruction,
		operands:    operands,
	}

	return decodedInstruction, nil
}
