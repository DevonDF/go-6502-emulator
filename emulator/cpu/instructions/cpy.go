package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CPYHandler struct {
}

var cpyHandler = &CPYHandler{}

func CPY() *CPYHandler {
	return cpyHandler
}

func (handler *CPYHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// Y - M
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	result := cpu.RegisterY - int8(value)

	cpu.SetNegativeFlag(result < 0)
	cpu.SetZeroFlag(result == 0)
	//cpu.SetCarryFlag() TODO
	return nil
}
