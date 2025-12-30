package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CMPHandler struct {
}

var cmpHandler = &CMPHandler{}

func CMP() *CMPHandler {
	return cmpHandler
}

func (handler *CMPHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	// A - M
	// so here we add the negation of the value
	result := cpu.RegisterAC - int8(value)

	cpu.SetNegativeFlag(result < 0)
	cpu.SetZeroFlag(result == 0)
	//cpu.SetCarryFlag() TODO
	return nil
}
