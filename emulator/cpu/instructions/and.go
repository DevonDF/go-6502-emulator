package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ANDHandler struct {
}

var andHandler = &ANDHandler{}

func AND() *ANDHandler {
	return andHandler
}

func (handler *ANDHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.Accumulator.And(uint8(value))
	return nil
}
