package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type BNEHandler struct {
}

var bneHandler = &BNEHandler{}

func BNE() *BNEHandler {
	return bneHandler
}

func (handler *BNEHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if Z = 0
	if cpu.GetZeroFlag() == 0 {
		// jump relative
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
