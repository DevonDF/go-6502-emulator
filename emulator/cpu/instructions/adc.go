package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ADCHandler struct {
}

var adcHandler = &ADCHandler{}

func ADC() *ADCHandler {
	return adcHandler
}

func (handler *ADCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.Accumulator.Add(int8(value))
	return nil
}
