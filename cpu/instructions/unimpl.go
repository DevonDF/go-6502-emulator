package instructions

import (
	"emulator/cpu"
	"emulator/memory"
	"errors"
)

type UnimplHandler struct {
}

var unimpHandler = &UnimplHandler{}

func Unimpl() *UnimplHandler {
	return unimpHandler
}

func (handler *UnimplHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	return errors.New("unimplemented instruction")
}
