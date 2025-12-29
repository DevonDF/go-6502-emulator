package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
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
		cpu.RegisterPC += uint16(instruction.Operands[0])
	}
	return nil
}
