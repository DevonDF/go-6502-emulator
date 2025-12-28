package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type BCCHandler struct {
}

var bccHandler = &BCCHandler{}

func BCC() *BCCHandler {
	return bccHandler
}

func (handler *BCCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if C = 0
	if cpu.GetCarryFlag() == 0 {
		// jump relative
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
