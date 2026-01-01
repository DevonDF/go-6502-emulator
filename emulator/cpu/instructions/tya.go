package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type TYAHandler struct {
}

var tyaHandler = &TYAHandler{}

func TYA() *TYAHandler {
	return tyaHandler
}

func (handler *TYAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// Y -> A
	cpu.RegisterAC = cpu.RegisterY
	return nil
}
