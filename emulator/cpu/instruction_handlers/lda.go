package instruction_handlers

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type LDAHandler struct {
}

var ldaHandler = &LDAHandler{}

func LDA() *LDAHandler {
	return ldaHandler
}

func (handler *LDAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	var value byte
	var err error

	switch instruction.Instruction.Opcode {
	case 0xA9: // immediate	LDA #oper
		value = byte(instruction.Operands[0])

	case 0xA5: // zeropage	LDA oper
		value, err = addressing.ReadZeropage(instruction.Operands, cpu, memory)

	case 0xB5: // zeropage,X	LDA oper,X
		value, err = addressing.ReadZeropageX(instruction.Operands, cpu, memory)

	case 0xAD: // absolute	LDA oper
		value, err = addressing.ReadAbsolute(instruction.Operands, cpu, memory)

	case 0xBD: // absolute,X	LDA oper,X
		value, err = addressing.ReadAbsoluteX(instruction.Operands, cpu, memory)

	case 0xB9: // absolute,Y	LDA oper,Y
		value, err = addressing.ReadAbsoluteY(instruction.Operands, cpu, memory)

	case 0xA1: // (indirect,X)	LDA (oper,X)
		value, err = addressing.ReadIndirectX(instruction.Operands, cpu, memory)

	case 0xA2: // (indirect),Y	LDA (oper),Y
		value, err = addressing.ReadIndirectY(instruction.Operands, cpu, memory)
	}

	if err != nil {
		return err
	}
	cpu.RegisterAC = int8(value)
	return nil
}
