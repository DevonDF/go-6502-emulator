package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CPXHandler struct {
}

var cpxHandler = &CPXHandler{}

func CPX() *CPXHandler {
	return cpxHandler
}

func (handler *CPXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// X - M
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	// A - M
	// so here we add the negation of the value
	result := cpu.RegisterX - int8(value)

	cpu.SetNegativeFlag(result < 0)
	cpu.SetZeroFlag(result == 0)
	//cpu.SetCarryFlag() TODO
	return nil
}
