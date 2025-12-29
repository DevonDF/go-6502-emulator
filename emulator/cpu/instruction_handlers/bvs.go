package instruction_handlers

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
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
