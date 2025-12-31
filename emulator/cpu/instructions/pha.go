package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type PHAHandler struct {
}

var phaHandler = &PHAHandler{}

func PHA() *PHAHandler {
	return phaHandler
}

func (handler *PHAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// push A
	cpu.Stack.PushByte(cpu.RegisterAC, memory)
	return nil
}
