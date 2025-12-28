package instructions

import (
	"emulator/cpu"
	"emulator/memory"
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
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
