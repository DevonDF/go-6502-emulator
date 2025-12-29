package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BEQHandler struct {
}

var beqHandler = &BEQHandler{}

func BEQ() *BEQHandler {
	return beqHandler
}

func (handler *BEQHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if Z = 1
	if cpu.GetZeroFlag() == 1 {
		// jump relative
		cpu.RegisterPC += uint16(instruction.Operands[0])
	}
	return nil
}
