package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CLCHandler struct {
}

var clcHandler = &CLCHandler{}

func CLC() *CLCHandler {
	return clcHandler
}

func (handler *CLCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// 0 -> C
	cpu.SetCarryFlag(false)
	return nil
}
