package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BPLHandler struct {
}

var bplHandler = &BPLHandler{}

func BPL() *BPLHandler {
	return bplHandler
}

func (handler *BPLHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if N = 0
	if cpu.GetNegativeFlag() == 0 {
		// jump relative
		cpu.RegisterPC += uint16(instruction.Operands[0])
	}
	return nil
}
