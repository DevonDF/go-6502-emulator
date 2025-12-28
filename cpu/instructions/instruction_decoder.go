package instructions

import (
	"bufio"
	"emulator/cpu"
	"emulator/memory"
	"fmt"
)

// InstructionHandler defines a handler that must implement an execute function.
type InstructionHandler interface {
	Execute(*cpu.CPU, *memory.Memory, *DecodedInstruction) error // execute the instruction on the given CPU & Memory.
}

// Instruction defines an instruction within the 6502 instruction set.
type Instruction struct {
	Opcode  byte               // the Opcode for the given instruction.
	Size    byte               // the number of bytes this instruction takes.
	Cycles  byte               // the number of cpu Cycles this operation takes.
	Handler InstructionHandler // the Handler for this instruction.
}

// DecodedInstruction defines a decoded instruction within the 6502 instruction set.
type DecodedInstruction struct {
	Instruction *Instruction // the instruction that this relates to.
	Operands    []byte       // the operands for the given instruction.
}

// getInstructionSet returns the set of instructions as a map of Opcode to Instruction.
func getInstructionSet() map[byte]Instruction {
	return map[byte]Instruction{
		// ADC - Add Memory to Accumulator with Carry
		0x69: { // ADC #oper
			Opcode:  0x69,
			Size:    2,
			Cycles:  2,
			Handler: ADC(),
		},
		0x65: { // ADC oper
			Opcode:  0x65,
			Size:    2,
			Cycles:  3,
			Handler: ADC(),
		},
		0x75: { // ADC oper,X
			Opcode:  0x75,
			Size:    2,
			Cycles:  4,
			Handler: ADC(),
		},
		0x6D: { // ADC oper
			Opcode:  0x6D,
			Size:    3,
			Cycles:  4,
			Handler: ADC(),
		},
		0x7D: { // ADC oper,X
			Opcode:  0x7D,
			Size:    3,
			Cycles:  4,
			Handler: ADC(),
		},
		0x79: { // ADC oper,Y
			Opcode:  0x79,
			Size:    3,
			Cycles:  4,
			Handler: ADC(),
		},
		0x61: { // ADC (oper,X)
			Opcode:  0x61,
			Size:    2,
			Cycles:  6,
			Handler: ADC(),
		},
		0x71: { // ADC (oper),Y
			Opcode:  0x71,
			Size:    2,
			Cycles:  5,
			Handler: ADC(),
		},

		// AND - AND Memory with Accumulator
		0x29: { // AND #oper
			Opcode:  0x29,
			Size:    2,
			Cycles:  2,
			Handler: AND(),
		},
		0x25: { // AND oper
			Opcode:  0x25,
			Size:    2,
			Cycles:  3,
			Handler: AND(),
		},
		0x35: { // AND oper,X
			Opcode:  0x35,
			Size:    2,
			Cycles:  4,
			Handler: AND(),
		},
		0x2D: { // AND oper
			Opcode:  0x2D,
			Size:    3,
			Cycles:  4,
			Handler: AND(),
		},
		0x3D: { // AND oper,X
			Opcode:  0x7D,
			Size:    3,
			Cycles:  4,
			Handler: AND(),
		},
		0x39: { // AND oper,Y
			Opcode:  0x39,
			Size:    3,
			Cycles:  4,
			Handler: AND(),
		},
		0x21: { // AND (oper,X)
			Opcode:  0x21,
			Size:    2,
			Cycles:  6,
			Handler: AND(),
		},
		0x31: { // AND (oper),Y
			Opcode:  0x31,
			Size:    2,
			Cycles:  5,
			Handler: AND(),
		},

		// ASL - Shift Left One Bit (Memory or Accumulator)
		0x0A: { // ASL A
			Opcode:  0x0A,
			Size:    1,
			Cycles:  2,
			Handler: ASL(),
		},
		0x06: { // ASL oper
			Opcode:  0x06,
			Size:    2,
			Cycles:  5,
			Handler: ASL(),
		},
		0x16: { // ASL oper,X
			Opcode:  0x16,
			Size:    2,
			Cycles:  6,
			Handler: ASL(),
		},
		0x0E: { // ASL oper
			Opcode:  0x0E,
			Size:    3,
			Cycles:  6,
			Handler: ASL(),
		},
		0x1E: { // ASL oper,X
			Opcode:  0x1E,
			Size:    3,
			Cycles:  7,
			Handler: ASL(),
		},

		// BRK - Break & Interrupt
		0x00: { // BRK
			Opcode:  0x00,
			Size:    1,
			Cycles:  7,
			Handler: BRK(),
		},
	}
}

// DecodeNextInstruction decodes the next instruction from a buffered byte reader and returns a DecodedInstruction.
func DecodeNextInstruction(reader *bufio.Reader) (DecodedInstruction, error) {
	Opcode, err := reader.ReadByte()
	if err != nil {
		return DecodedInstruction{}, err
	}

	instructions := getInstructionSet()
	instruction, found := instructions[Opcode]
	if !found {
		return DecodedInstruction{}, fmt.Errorf("no opcode found for %02X", Opcode)
	}

	operands := make([]byte, instruction.Size-1)
	_, err = reader.Read(operands)
	if err != nil {
		return DecodedInstruction{}, fmt.Errorf("failed to read operands for Opcode %02X: %w", Opcode, err)
	}

	decodedInstruction := DecodedInstruction{
		Instruction: &instruction,
		Operands:    operands,
	}

	return decodedInstruction, nil
}
