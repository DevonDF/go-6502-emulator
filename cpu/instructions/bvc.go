package instructions

import (
	"emulator/cpu"
	"emulator/memory"
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
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
