package instructions

import (
	"bufio"
	"fmt"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

// InstructionHandler defines a handler that must implement an execute function.
type InstructionHandler interface {
	Execute(*cpu.CPU, *memory.Memory, *DecodedInstruction) error // execute the instruction on the given CPU & Memory.
}

// DecodedInstruction defines a decoded instruction within the 6502 instruction set.
type DecodedInstruction struct {
	Instruction *Instruction // the instruction that this relates to.
	Operands    []byte       // the operands for the given instruction.
}

// DecodeNextInstruction decodes the next instruction from a buffered byte reader and returns a DecodedInstruction.
func DecodeNextInstruction(reader *bufio.Reader) (DecodedInstruction, error) {
	opcode, err := reader.ReadByte()
	if err != nil {
		return DecodedInstruction{}, err
	}

	instruction := InstructionFromOpcode(opcode)

	operands := make([]byte, instruction.Size-1)
	_, err = reader.Read(operands)
	if err != nil {
		return DecodedInstruction{}, fmt.Errorf("failed to read operands for opcode %02X: %w", opcode, err)
	}

	decodedInstruction := DecodedInstruction{
		Instruction: instruction,
		Operands:    operands,
	}

	return decodedInstruction, nil
}
