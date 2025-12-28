package instructions

import (
	"emulator/cpu"
	"emulator/memory"
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
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
