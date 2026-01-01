package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type TXAHandler struct {
}

var txaHandler = &TXAHandler{}

func TXA() *TXAHandler {
	return txaHandler
}

func (handler *TXAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// X -> A
	cpu.RegisterAC = cpu.RegisterX
	return nil
}
