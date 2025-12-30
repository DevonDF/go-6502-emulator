package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CLIHandler struct {
}

var cliHandler = &CLIHandler{}

func CLI() *CLIHandler {
	return cliHandler
}

func (handler *CLIHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// 0 -> I
	cpu.SetInterruptFlag(false)
	return nil
}
