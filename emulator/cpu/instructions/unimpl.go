package instructions

import (
	"errors"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
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
