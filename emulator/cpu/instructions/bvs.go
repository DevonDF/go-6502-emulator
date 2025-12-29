package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BVSHandler struct {
}

var bvsHandler = &BVSHandler{}

func BVS() *BVSHandler {
	return bvsHandler
}

func (handler *BVSHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if V = 1
	if cpu.GetOverflowFlag() == 1 {
		// jump relative
		cpu.RegisterPC += uint16(instruction.Operands[0])
	}
	return nil
}
