package instructions

import (
	"errors"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BRKHandler struct {
}

var brkHandler = &BRKHandler{}

func BRK() *BRKHandler {
	return brkHandler
}

func (handler *BRKHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	return errors.New("BRK executed")
}
