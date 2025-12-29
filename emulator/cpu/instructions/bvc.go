package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BVCHandler struct {
}

var bvcHandler = &BVCHandler{}

func BVC() *BVCHandler {
	return bvcHandler
}

func (handler *BVCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if V = 0
	if cpu.GetOverflowFlag() == 0 {
		// jump relative
		cpu.RegisterPC += uint16(instruction.Operands[0])
	}
	return nil
}
