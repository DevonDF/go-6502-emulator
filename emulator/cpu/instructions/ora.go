package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ORAHandler struct {
}

var oraHandler = &ORAHandler{}

func ORA() *ORAHandler {
	return oraHandler
}

func (handler *ORAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// A or M -> A
	value, err := instruction.GetOperandValue(cpu, memory)
	if err != nil {
		return err
	}

	cpu.RegisterAC = cpu.RegisterAC | int8(value)
	cpu.SetNegativeFlag(cpu.RegisterAC < 0)
	cpu.SetZeroFlag(cpu.RegisterAC == 0)
	return nil
}
