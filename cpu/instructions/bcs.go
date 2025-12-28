package instructions

import (
	"emulator/cpu"
	"emulator/memory"
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
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
