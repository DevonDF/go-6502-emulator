package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type RTIHandler struct {
}

var rtiHandler = &RTIHandler{}

func RTI() *RTIHandler {
	return rtiHandler
}

func (handler *RTIHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// pull SR
	// pull PC
	// The status register is pulled with the break flag and bit 5 ignored. Then PC is pulled from the stack. TODO
	sr, err := cpu.Stack.PopByte(memory)
	if err != nil {
		return err
	}
	cpu.RegisterSR = uint8(sr)

	pc, err := cpu.Stack.PopDouble(memory)
	if err != nil {
		return err
	}

	// the fetch-decode-execute cycle will increment by the instruction size, we should counteract this
	cpu.RegisterPC = uint16(pc) - uint16(instruction.Instruction.Size)
	return nil
}
