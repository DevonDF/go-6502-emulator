package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type TAXHandler struct {
}

var taxHandler = &TAXHandler{}

func TAX() *TAXHandler {
	return taxHandler
}

func (handler *TAXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// A -> X
	cpu.RegisterX = cpu.RegisterAC
	return nil
}
