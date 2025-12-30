package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
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
		// cast the register & operand to SIGNED int32 such that negatives work here
		cpu.RegisterPC = uint16(int32(cpu.RegisterPC) + int32(int8(instruction.Operands[0])))
	}
	return nil
}
