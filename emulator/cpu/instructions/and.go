package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ANDHandler struct {
}

var andHandler = &ANDHandler{}

func AND() *ANDHandler {
	return andHandler
}

func (handler *ANDHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {

	var err error
	var value byte

	switch instruction.Instruction.Opcode {
	case 0x29: // immediate	AND #oper
		value = byte(instruction.Operands[0])

	case 0x25: // zeropage	AND oper
		value, err = addressing.ReadZeropage(instruction.Operands, cpu, memory)

	case 0x35: // zeropage,X	AND oper,X
		value, err = addressing.ReadZeropageX(instruction.Operands, cpu, memory)

	case 0x2D: // absolute	AND oper
		value, err = addressing.ReadAbsolute(instruction.Operands, cpu, memory)

	case 0x3D: // absolute,X	AND oper,X
		value, err = addressing.ReadAbsoluteX(instruction.Operands, cpu, memory)

	case 0x39: // absolute,Y	AND oper,Y
		value, err = addressing.ReadAbsoluteY(instruction.Operands, cpu, memory)

	case 0x21: // (indirect,X)	AND (oper,X)
		value, err = addressing.ReadIndirectX(instruction.Operands, cpu, memory)

	case 0x31: // (indirect),Y	AND (oper),Y
		value, err = addressing.ReadAbsoluteY(instruction.Operands, cpu, memory)
	}

	if err != nil {
		return err
	}

	cpu.Accumulator.And(uint8(value))
	return nil
}
