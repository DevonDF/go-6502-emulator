package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type EORHandler struct {
}

var eorHandler = &EORHandler{}

func EOR() *EORHandler {
	return eorHandler
}

func (handler *EORHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// A xor M -> A
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.RegisterAC = cpu.RegisterAC ^ int8(value)
	cpu.SetNegativeFlag(cpu.RegisterAC < 0)
	cpu.SetZeroFlag(cpu.RegisterAC == 0)
	return nil
}
