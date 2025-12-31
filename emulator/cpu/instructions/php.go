package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type PHPHandler struct {
}

var phpHandler = &PHPHandler{}

func PHP() *PHPHandler {
	return phpHandler
}

func (handler *PHPHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// push SR
	// The status register will be pushed with the break flag and bit 5 set to 1. TODO
	cpu.Stack.PushByte(int8(cpu.RegisterSR), memory)
	return nil
}
