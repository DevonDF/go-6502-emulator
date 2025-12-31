package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type DEYHandler struct {
}

var deyHandler = &DEYHandler{}

func DEY() *DEYHandler {
	return deyHandler
}

func (handler *DEYHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// Y <- Y - 1
	cpu.RegisterY = cpu.RegisterY - 1

	cpu.SetNegativeFlag(cpu.RegisterY < 0)
	cpu.SetZeroFlag(cpu.RegisterY == 0)
	return nil
}
