package instructions

import (
	"emulator/cpu"
	"emulator/memory"
	"errors"
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
