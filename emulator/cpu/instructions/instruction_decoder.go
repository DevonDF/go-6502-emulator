package instructions

import (
	"bufio"
	"fmt"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
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

// GetOperandValue gets the value of the operand using the correct addressing mode for the instruction.
func (instruction *DecodedInstruction) GetOperandValue(cpu_ *cpu.CPU, memory_ *memory.Memory) (byte, error) {
	var value byte
	var err error

	switch instruction.Instruction.AddressingMode {
	case AddrAccumulator:
		return 0x0, fmt.Errorf("attempt to call GetOperandValue on accumulator operand")

	case AddrImplied:
		return 0x0, fmt.Errorf("attempt to call GetOperandValue on implied operand")

	case AddrImmediate:
		value = byte(instruction.Operands[0])

	case AddrZeropage:
		value, err = addressing.ReadZeropage(instruction.Operands, cpu_, memory_)

	case AddrZeropageX:
		value, err = addressing.ReadZeropageX(instruction.Operands, cpu_, memory_)

	case AddrAbsolute:
		value, err = addressing.ReadAbsolute(instruction.Operands, cpu_, memory_)

	case AddrAbsoluteX:
		value, err = addressing.ReadAbsoluteX(instruction.Operands, cpu_, memory_)

	case AddrAbsoluteY:
		value, err = addressing.ReadAbsoluteY(instruction.Operands, cpu_, memory_)

	case AddrIndirectX:
		value, err = addressing.ReadIndirectX(instruction.Operands, cpu_, memory_)

	case AddrIndirectY:
		value, err = addressing.ReadIndirectY(instruction.Operands, cpu_, memory_)

	default:
		return 0x0, fmt.Errorf("invalid addressing mode for instruction: %v", instruction)
	}

	return value, err
}

// GetOperandReferencedAddress gets the raw memory address that the operand is referencing.
func (instruction *DecodedInstruction) GetOperandReferencedAddress(cpu_ *cpu.CPU, memory_ *memory.Memory) (uint16, error) {
	var addr uint16
	var err error

	switch instruction.Instruction.AddressingMode {
	case AddrAccumulator:
		return 0x0, fmt.Errorf("attempt to call GetOperandReferencedAddress on accumulator operand")

	case AddrImplied:
		return 0x0, fmt.Errorf("attempt to call GetOperandReferencedAddress on implied operand")

	case AddrZeropage:
		addr, err = addressing.GetZeropageAddress(instruction.Operands, cpu_, memory_)

	case AddrZeropageX:
		addr, err = addressing.GetZeropageXAddress(instruction.Operands, cpu_, memory_)

	case AddrAbsolute:
		addr, err = addressing.GetAbsoluteAddress(instruction.Operands, cpu_, memory_)

	case AddrAbsoluteX:
		addr, err = addressing.GetAbsoluteXAddress(instruction.Operands, cpu_, memory_)

	case AddrAbsoluteY:
		addr, err = addressing.GetAbsoluteYAddress(instruction.Operands, cpu_, memory_)

	case AddrIndirectX:
		addr, err = addressing.GetIndirectXAddress(instruction.Operands, cpu_, memory_)

	case AddrIndirectY:
		addr, err = addressing.GetIndirectYAddress(instruction.Operands, cpu_, memory_)

	default:
		return 0x0, fmt.Errorf("invalid addressing mode for instruction: %v", instruction)
	}

	return addr, err
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
