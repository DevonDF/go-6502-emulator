package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type INXHandler struct {
}

var inxHandler = &INXHandler{}

func INX() *INXHandler {
	return inxHandler
}

func (handler *INXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// X + 1 -> X
	cpu.RegisterX = cpu.RegisterX + 1

	cpu.SetNegativeFlag(cpu.RegisterX < 0)
	cpu.SetZeroFlag(cpu.RegisterX == 0)
	return nil
}
