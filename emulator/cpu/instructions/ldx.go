package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type LDXHandler struct {
}

var ldxHandler = &LDXHandler{}

func LDX() *LDXHandler {
	return ldxHandler
}

func (handler *LDXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// M -> X
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.RegisterX = int8(value)
	cpu.SetNegativeFlag(cpu.RegisterX < 0)
	cpu.SetZeroFlag(cpu.RegisterX == 0)
	return nil
}
