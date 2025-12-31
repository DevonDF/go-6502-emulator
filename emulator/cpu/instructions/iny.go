package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type INYHandler struct {
}

var inyHandler = &INYHandler{}

func INY() *INYHandler {
	return inyHandler
}

func (handler *INYHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// Y + 1 -> Y
	cpu.RegisterY = cpu.RegisterY + 1

	cpu.SetNegativeFlag(cpu.RegisterY < 0)
	cpu.SetZeroFlag(cpu.RegisterY == 0)
	return nil
}
