package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type PLAHandler struct {
}

var plaHandler = &PLAHandler{}

func PLA() *PLAHandler {
	return plaHandler
}

func (handler *PLAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// pull A
	poppedByte, err := cpu.Stack.PopByte(memory)
	if err != nil {
		return err
	}
	cpu.RegisterAC = poppedByte
	return nil
}
