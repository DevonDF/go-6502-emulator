package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type TAYHandler struct {
}

var tayHandler = &TAYHandler{}

func TAY() *TAYHandler {
	return tayHandler
}

func (handler *TAYHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// A -> Y
	cpu.RegisterY = cpu.RegisterAC
	return nil
}
