package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ADCHandler struct {
}

var adcHandler = &ADCHandler{}

func ADC() *ADCHandler {
	return adcHandler
}

func (handler *ADCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {

	var valueToAdd byte
	var err error

	switch instruction.Instruction.Opcode {
	case 0x69: // immediate	ADC #oper
		valueToAdd = byte(instruction.Operands[0])

	case 0x65: // zeropage	ADC oper
		valueToAdd, err = addressing.ReadZeropage(instruction.Operands, cpu, memory)

	case 0x75: // zeropage,X	ADC oper,X
		valueToAdd, err = addressing.ReadZeropageX(instruction.Operands, cpu, memory)

	case 0x6D: // absolute	ADC oper
		valueToAdd, err = addressing.ReadAbsolute(instruction.Operands, cpu, memory)

	case 0x7D: // absolute,X	ADC oper,X
		valueToAdd, err = addressing.ReadAbsoluteX(instruction.Operands, cpu, memory)

	case 0x79: // absolute,Y	ADC oper,Y
		valueToAdd, err = addressing.ReadAbsoluteY(instruction.Operands, cpu, memory)

	case 0x61: // (indirect,X)	ADC (oper,X)
		valueToAdd, err = addressing.ReadIndirectX(instruction.Operands, cpu, memory)

	case 0x71: // (indirect),Y	ADC (oper),Y
		valueToAdd, err = addressing.ReadIndirectY(instruction.Operands, cpu, memory)
	}

	if err != nil {
		return err
	}

	cpu.Accumulator.Add(int8(valueToAdd))
	return nil
}
