package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CLDHandler struct {
}

var cldHandler = &CLDHandler{}

func CLD() *CLDHandler {
	return cldHandler
}

func (handler *CLDHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// 0 -> D
	cpu.SetDecimalFlag(false)
	return nil
}
