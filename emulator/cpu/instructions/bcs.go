package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BCSHandler struct {
}

var bcsHandler = &BCSHandler{}

func BCS() *BCSHandler {
	return bcsHandler
}

func (handler *BCSHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if C = 1
	if cpu.GetCarryFlag() == 1 {
		// jump relative
		cpu.RegisterPC += uint16(instruction.Operands[0])
	}
	return nil
}
