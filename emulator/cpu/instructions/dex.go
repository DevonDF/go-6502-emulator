package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type DEXHandler struct {
}

var dexHandler = &DEXHandler{}

func DEX() *DEXHandler {
	return dexHandler
}

func (handler *DEXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// X <- X - 1
	cpu.RegisterX = cpu.RegisterX - 1

	cpu.SetNegativeFlag(cpu.RegisterX < 0)
	cpu.SetZeroFlag(cpu.RegisterX == 0)
	return nil
}
