package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type TXSHandler struct {
}

var txsHandler = &TXSHandler{}

func TXS() *TXSHandler {
	return txsHandler
}

func (handler *TXSHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// X -> SP
	cpu.RegisterSP = cpu.RegisterX
	return nil
}
