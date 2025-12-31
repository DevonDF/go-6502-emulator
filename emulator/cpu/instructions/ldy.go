package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type LDYHandler struct {
}

var ldyHandler = &LDYHandler{}

func LDY() *LDYHandler {
	return ldyHandler
}

func (handler *LDYHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// M -> Y
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.RegisterY = int8(value)
	cpu.SetNegativeFlag(cpu.RegisterY < 0)
	cpu.SetZeroFlag(cpu.RegisterY == 0)
	return nil
}
