package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type ADCHandler struct {
}

var adcHandler = &ADCHandler{}

func ADC() *ADCHandler {
	return adcHandler
}

func (handler *ADCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	switch instruction.Instruction.Opcode {

	case 0x69:
		operand := int8(instruction.Operands[0])
		cpu.Accumulator.Add(operand)

	}
	return nil
}
