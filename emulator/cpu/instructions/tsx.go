package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type TSXHandler struct {
}

var tsxHandler = &TSXHandler{}

func TSX() *TSXHandler {
	return tsxHandler
}

func (handler *TSXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// SP -> X
	cpu.RegisterX = cpu.RegisterSP
	return nil
}
