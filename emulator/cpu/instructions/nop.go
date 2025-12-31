package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type NOPHandler struct {
}

var nopHandler = &NOPHandler{}

func NOP() *NOPHandler {
	return nopHandler
}

func (handler *NOPHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// NOP
	return nil
}
