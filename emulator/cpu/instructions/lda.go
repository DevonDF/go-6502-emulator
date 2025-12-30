package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type LDAHandler struct {
}

var ldaHandler = &LDAHandler{}

func LDA() *LDAHandler {
	return ldaHandler
}

func (handler *LDAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.RegisterAC = int8(value)
	cpu.SetNegativeFlag(cpu.RegisterAC < 0)
	cpu.SetZeroFlag(cpu.RegisterAC == 0)
	return nil
}
