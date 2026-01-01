package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type RTSHandler struct {
}

var rtsHandler = &RTSHandler{}

func RTS() *RTSHandler {
	return rtsHandler
}

func (handler *RTSHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// pull PC
	// PC+1 -> PC

	pc, err := cpu.Stack.PopDouble(memory)
	if err != nil {
		return err
	}

	// the fetch-decode-execute cycle will increment by the instruction size, we should counteract this
	cpu.RegisterPC = (uint16(pc) + 1) - uint16(instruction.Instruction.Size)
	return nil
}
